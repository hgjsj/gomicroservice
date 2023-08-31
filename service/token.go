package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spiffe/go-spiffe/v2/svid/jwtsvid"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

var secret = []byte("cloudservice")

func NewToken() (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "CloudService",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}


func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("No match algrithem: %s", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return false, err
	}

	if token.Valid {
		return true, nil
	} else {
		return false, token.Claims.Valid()
	}
}

type  SpiffeJwtSource struct {
	jwts	*workloadapi.JWTSource
}

var spiffeJwtSource = new(SpiffeJwtSource)

func NewSpiffeJWTSource(ctx context.Context,  socketPath string) (*SpiffeJwtSource, error){
	clientOptions := workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath))
	jwtSource, err := workloadapi.NewJWTSource(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("unable to create JWTSource: %s", err.Error())
	} 
	spiffeJwtSource.jwts = jwtSource
	return spiffeJwtSource, nil
}

func (j SpiffeJwtSource)NewSpiffeJWT(ctx context.Context,  spiffeID string) (string, error){
	if token, err := j.jwts.FetchJWTSVID(ctx, jwtsvid.Params{Audience: spiffeID,}); err == nil {
		return token.Marshal(), nil
	} else {
		return "", fmt.Errorf("unable to fetch JWT: %s", err.Error())
	}	
}

func (j SpiffeJwtSource)ValidateSpiffeJWT(ctx context.Context,  token string, audiences []string) (*jwtsvid.SVID, error){
	if svid, err := jwtsvid.ParseAndValidate(token, j.jwts, audiences); err == nil {
		return svid, nil
	} else {
		return nil, fmt.Errorf("unable to fetch JWT: %s", err.Error())
	}	
}

func (j SpiffeJwtSource) Close() {
	j.jwts.Close()
}
