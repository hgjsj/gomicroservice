package transport

import (
	"context"
	"errors"
	"fmt"
	"go-microservice/service"
	"net/url"
	"strings"
	"time"

	//"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/log"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
	//"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"

	//"github.com/sony/gobreaker"
	//"golang.org/x/time/rate"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type proxymw struct {
	next      service.StringService // Serve most requests via this service...
	uppercase endpoint.Endpoint     // ...except Uppercase, which gets served by this endpoint
}

func (mw proxymw) Uppercase(s string) (string, error) {
	var err error
	var response interface{}
	for num := 0; num < 3; num++ {
		response, err = mw.uppercase(context.Background(), UppercaseRequest{S: s})
		if err == nil {
			break
		}
	}

	resp := response.(UppercaseResponse)
	if resp.Err != "" {
		return resp.V, errors.New(resp.Err)
	}
	return resp.V, nil
}

func (mw proxymw) Count(s string) int {
	return mw.next.Count(s)
}

func ProxyingMiddleware(instance string, logger log.Logger) service.ServiceMiddleware {
	// If instances is empty, don't proxy.
	if instance == "" {
		return func(next service.StringService) service.StringService { return next }
	}

	var (
		instanceList = split(instance)
		endpointer   sd.FixedEndpointer
	)

	var (
		//qps         = 100                    // beyond which we will return an error
		maxAttempts = 3                      // per request, before giving up
		maxTime     = 250 * time.Millisecond // wallclock time, before giving up
	)

	logger.Log("proxy_to", fmt.Sprint(instanceList))
	for _, instance := range instanceList {
		var e endpoint.Endpoint
		e = makeUppercaseProxy(instance)
		//e = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(e)
		//e = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), qps))(e)
		endpointer = append(endpointer, e)
	}

	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(maxAttempts, maxTime, balancer)

	return func(next service.StringService) service.StringService {
		return proxymw{next, retry}
	}
}

func makeUppercaseProxy(instance string) endpoint.Endpoint {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://:" + instance
	}
	instance = instance + "/uppercase"
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}

	return httptransport.NewClient(
		"GET",
		u,
		EncodeRequest,
		DecodeUppercaseResponse,
	).Endpoint()
}

func split(s string) []string {
	a := strings.Split(s, ",")
	for i := range a {
		a[i] = strings.TrimSpace(a[i])
	}
	return a
}
