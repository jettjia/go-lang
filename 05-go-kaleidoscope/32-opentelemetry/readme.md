# 安装SigNoz

https://signoz.io/docs/install/docker/

```shell
git clone -b main https://github.com/SigNoz/signoz.git && cd signoz/deploy/

docker-compose -f docker/clickhouse-setup/docker-compose.yaml up -d

docker ps


```

```
CONTAINER ID   IMAGE                                          COMMAND                  CREATED          STATUS                    PORTS                                                                            NAMES
1ad413fc12aa   signoz/frontend:0.8.0                          "nginx -g 'daemon of…"   20 minutes ago   Up 20 minutes             80/tcp, 0.0.0.0:3301->3301/tcp, :::3301->3301/tcp                                frontend
419f7b440412   signoz/alertmanager:0.23.0-0.1                 "/bin/alertmanager -…"   20 minutes ago   Up 20 minutes             9093/tcp                                                                         clickhouse-setup_alertmanager_1
95f5fab00c3c   signoz/otelcontribcol:0.43.0-0.1               "/otelcontribcol --c…"   21 minutes ago   Up 21 minutes             0.0.0.0:4317-4318->4317-4318/tcp, :::4317-4318->4317-4318/tcp, 55679-55680/tcp   clickhouse-setup_otel-collector_1
c1640c215d10   signoz/otelcontribcol:0.43.0-0.1               "/otelcontribcol --c…"   21 minutes ago   Up 21 minutes             4317/tcp, 55679-55680/tcp                                                        clickhouse-setup_otel-collector-metrics_1
9db88c61f7fd   signoz/query-service:0.8.0                     "./query-service -co…"   21 minutes ago   Up 21 minutes (healthy)   8080/tcp                                                                         query-service
509ab96c5393   clickhouse/clickhouse-server:22.4-alpine       "/entrypoint.sh"         22 minutes ago   Up 21 minutes (healthy)   8123/tcp, 9000/tcp, 9009/tcp                                                     clickhouse-setup_clickhouse_1
eb7a2e23c0c0   grubykarol/locust:1.2.3-python3.9-alpine3.12   "/docker-entrypoint.…"   22 minutes ago   Up 21 minutes             5557-5558/tcp, 8089/tcp                                                          load-hotrod
f234b5cb4512   jaegertracing/example-hotrod:1.30              "/go/bin/hotrod-linu…"   22 minutes ago   Up 21 minutes             8080-8083/tcp                                                                    hotrod
```



访问 UI

http://10.4.7.71:3301



# gin使用

参考：

https://signoz.io/blog/opentelemetry-gin/



## step1

在`main.go`中声明以下变量,我们将使用这些变量来配置 OpenTelemetry:

```go
var (
    serviceName  = os.Getenv("SERVICE_NAME")
    collectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
    insecure     = os.Getenv("INSECURE_MODE")
)
```



## step2

初始化 OpenTelemetry。在`main.go`文件中添加以下代码片

```go
import (
  .....

    "github.com/gin-gonic/gin"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"

    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() func(context.Context) error {

    secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
    if len(insecure) > 0 {
        secureOption = otlptracegrpc.WithInsecure()
    }

    exporter, err := otlptrace.New(
        context.Background(),
        otlptracegrpc.NewClient(
            secureOption,
            otlptracegrpc.WithEndpoint(collectorURL),
        ),
    )

    if err != nil {
        log.Fatal(err)
    }
    resources, err := resource.New(
        context.Background(),
        resource.WithAttributes(
            attribute.String("service.name", serviceName),
            attribute.String("library.language", "go"),
        ),
    )
    if err != nil {
        log.Printf("Could not set resources: ", err)
    }

    otel.SetTracerProvider(
        sdktrace.NewTracerProvider(
            sdktrace.WithSampler(sdktrace.AlwaysSample()),
            sdktrace.WithBatcher(exporter),
            sdktrace.WithResource(resources),
        ),
    )
    return exporter.Shutdown
}
```



## step3

在 main.go 中初始化跟踪器

```go
func main() {
    cleanup := initTracer()
    defer cleanup(context.Background())

    ......
}
```



## step4

在`main.go`中添加以下行来配置 Gin 以使用中间件。

```go
import (
    ....
  "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
    ......
    r := gin.Default()
    r.Use(otelgin.Middleware(serviceName))
    ......
}
```

## step5

设置环境变量并运行您的 Go Gin 应用程序

现在您已经使用 OpenTelemetry 检测了 Go Gin 应用程序,您需要设置一些环境变量以将数据发送到 SigNoz 后端:

`SERVICE_NAME`:goGinApp(你可以随意命名)

`OTEL_EXPORTER_OTLP_ENDPOINT` 10.4.7.71:4317  ( signoz/signoz-otel-collector收集器的地址)



因此,最终的运行命令如下所示:

```shell
SERVICE_NAME=goGinApp INSECURE_MODE=true OTEL_EXPORTER_OTLP_ENDPOINT=localhost:4317 go run main.go
```



## step6

创造测试数据

点击书店应用程序的`/books`端点http://localhost:8090/books。多次刷新它以产生负载,然后等待 1-2 分钟让数据出现在 SigNoz 仪表板上。



## step7

控制台验证

http://10.4.7.71:3301/application



