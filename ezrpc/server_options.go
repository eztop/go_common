package ezrpc

import (
	"github.com/eztop/go_common/ezrpc/interceptor"
)

// ServerOptions ServerOptions
type ServerOptions struct {
	codec        Codec
	interceptors []interceptor.ServerInterceptor
	router       Router
}

// ServerOption ServerOption
type ServerOption func(*ServerOptions)

var (
	defaultServerOptions = ServerOptions{
		codec:        JSONCodec(),
		interceptors: []interceptor.ServerInterceptor{},
		router:       NewRouter(),
	}
)

// ServerWithRouter ServerWithRouter
func ServerWithRouter(router Router) ServerOption {
	return func(o *ServerOptions) {
		o.router = router
	}
}

// ServerWithCodec ServerWithCodec
func ServerWithCodec(codec Codec) ServerOption {
	return func(o *ServerOptions) {
		o.codec = codec
	}
}

// ServerWithInterceptors ServerWithInterceptors
func ServerWithInterceptors(interceptors ...interceptor.ServerInterceptor) ServerOption {
	return func(o *ServerOptions) {
		o.interceptors = interceptors
	}
}
