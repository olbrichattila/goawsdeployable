// Package httpconfig will be built if you are building for http
// Please make sure you are using the same structure for both
package httpconfig

import "sharedconfig"

// New is a function to create a new shared config.
// This is a config, with shared valeus, and lambda config specific values
// The interface ensures you will use the same strucure for both enfironment
// First implement the getter into the shared config interface above. Add getters only!
func New() sharedconfig.SharedConfiger {
	return &config{}
}

type config struct {
	// todo your config here
}

func (c *config) GetConfigType() string {
	return sharedconfig.TypeHttp
}
