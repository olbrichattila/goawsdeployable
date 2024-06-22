// Package sharedconfig is an interface to be used across both config files. The config have to use getters
// to ensure they follow the same structure
package sharedconfig

const (
	TypeLambda = "lambda"
	TypeHttp   = "http"
)

type SharedConfiger interface {
	GetConfigType() string
}
