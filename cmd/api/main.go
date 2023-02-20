// Code generated by hertz generator.

package main

import (
	"mydouyin/cmd/api/biz/cache"
	"mydouyin/cmd/api/biz/mw"
	"mydouyin/cmd/api/biz/rpc"
	videohandel "mydouyin/cmd/api/biz/videoHandel"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/network/standard"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
)

func Init() {
	mw.InitJWT()
	rpc.Init()
	cache.Init()
	videohandel.Init()
	//hlog init
	hlog.SetLogger(hertzlogrus.NewLogger())
	hlog.SetLevel(hlog.LevelInfo)
}

func main() {
	Init()
	tracer, cfg := tracing.NewServerTracer()
	h := server.New(
		server.WithHostPorts(":8080"),
		server.WithStreamBody(true),
		server.WithTransport(standard.NewTransporter),
		server.WithHandleMethodNotAllowed(true),
		tracer,
	)
	//use pprof mw
	pprof.Register(h)
	//user otel mw
	h.Use(tracing.ServerMiddleware(cfg))
	register(h)
	h.Spin()
}
