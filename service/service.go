package service

import (
	"errors"

	"strings"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// type GrpcStringService interface {
// 	Uppercase(string) (string, error)
// 	Count(string) int
// }

type stringService struct{}

var ErrEmpty = errors.New("Empty string")

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

type ServiceMiddleware func(StringService) StringService

// func (s stringService) Uppercase(str string) (string, error) {
// 	cstr, err := s.HttpUppercase(str)
// 	if err != nil {
// 		return "", err
// 	}
// 	return cstr, nil
// }

// func (s stringService) Count(str string) int {
// 	count := s.HttpCount(str)
// 	return count
// }

func NewStringService() StringService {
	return &stringService{}
}

// func NewGrpcStringService() GrpcStringService {
// 	return &stringService{}
// }
