module go-otel

go 1.16

require (
	github.com/gin-gonic/gin v1.9.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.40.0
	go.opentelemetry.io/otel v1.14.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.14.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.14.0
	go.opentelemetry.io/otel/sdk v1.14.0
	google.golang.org/grpc v1.54.0
)
