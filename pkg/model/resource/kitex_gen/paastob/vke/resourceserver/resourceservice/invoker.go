// Code generated by Kitex v0.1.0. DO NOT EDIT.

package resourceservice

import (
	"github.com/cloudwego/kitex/server"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/paastob/vke/resourceserver"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler resourceserver.ResourceService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}