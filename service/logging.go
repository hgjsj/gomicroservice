package service

import (
	"time"

	"github.com/go-kit/log"
)

type LoggingMiddleware struct {
	logger   log.Logger
	next     StringService
	protocal string
}

func (mw LoggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"time", begin.UTC().Format(time.RFC3339),
			"protocol", mw.protocal,
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func (mw LoggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"time", begin.UTC().Format(time.RFC3339),
			"protocol", mw.protocal,
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.next.Count(s)
	return
}

func NewLogging(logger log.Logger, next StringService, protocol string) StringService {
	return LoggingMiddleware{logger, next, protocol}
}
