package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() (err error) {
	// 优雅地处理中断信号（Ctrl+C）。
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// 设置并初始化 OpenTelemetry SDK。
	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		return
	}
	// 确保在程序结束之前调用 shutdown 方法清理资源。
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// 启动 HTTP 服务器。
	srv := &http.Server{
		Addr:         ":8080",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// 等待中断信号
	select {
	case err = <-srvErr:
		// 启动 HTTP 服务器时出错。
		return
	case <-ctx.Done():
		// 等待第一个 CTRL+C 信号。
		// 尽快停止接受信号通知。
		stop()
	}

	err = srv.Shutdown(context.Background())
	return
	// http.HandleFunc("/rolldice", rolldice)

	// srv := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: nil, // Use http.DefaultServeMux
	// }

	// // Start server
	// go func() {
	// 	log.Println("Starting server on :8080")
	// 	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 		log.Fatalf("ListenAndServe(): %v", err)
	// 	}
	// }()

	// // Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit
	// log.Println("Shutting down server...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server forced to shutdown:", err)
	// }

	// log.Println("Server exiting")
}

func newHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	// handleFunc 是对 mux.HandleFunc 的封装，
	// 它将 handler 注册到指定路径 pattern（如 "/rolldice/"）上，
	// 并在 OpenTelemetry 中记录该路径作为 http.route 插桩标签，用于丰富 HTTP 插桩信息。
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// 为 HTTP 插桩配置 "http.route" 标签。
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		// 这里是真正的处理函数。
		mux.Handle(pattern, handler)
	}

	// 注册 Handler。
	handleFunc("/rolldice/", rolldice)
	handleFunc("/rolldice/{player}", rolldice)

	// 为整个服务器添加 HTTP 插桩处理器。
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}
