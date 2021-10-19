package instana

import (
	instana "github.com/instana/go-sensor"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"
	"os"
)

type Instana struct {
	Logger *logrus.Logger
}

func NewInstana(logger *logrus.Logger) Instana {
	return Instana {
		Logger : logger,
	}
}

func (it *Instana) InitMetrics(context echo.Context) {
	instana.InitSensor(&instana.Options{
		Service: "cv-manager-api",
		LogLevel: instana.Debug,
	})

	tracer := instana.NewTracerWithOptions(instana.DefaultOptions())

	opentracing.InitGlobalTracer(tracer)

	opts := []opentracing.StartSpanOption{
		ext.SpanKindRPCServer,
		opentracing.Tags{
			"http.host": context.Request().Host,
			"http.method": context.Request().Method,
			"http.protocol": context.Request().URL.Scheme,
			"http.path": context.Request().URL.Path,
		},
	}

	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(context.Request().Header),
	)

	if err!=nil {
		opts = append(opts, ext.RPCServerOption(wireContext))
	}

	span := opentracing.GlobalTracer().StartSpan("g.http", opts...)

	defer span.Finish()
	if _,ok := os.LookupEnv("INSTANA_DEBUG"); ok {
		it.Logger.Level = logrus.DebugLevel
	}

	instana.SetLogger(it.Logger)
}


