package app

import (
	"context"
	"fmt"
	golog "log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type LogWriter struct {
	*zap.Logger
}

func (l LogWriter) Write(data []byte) (n int, err error) {
	l.Error("Api Server Error", zap.ByteString("msg", data))
	return len(data), nil
}

func Start(ctx context.Context, config Config) {

	r := newRouter(config)

	portString := fmt.Sprintf(":%d", config.Port)
	srv := &http.Server{
		Addr:         portString,
		Handler:      r,
		ReadTimeout:  time.Duration(10) * time.Second,
		WriteTimeout: time.Duration(10) * time.Second,
		IdleTimeout:  time.Duration(120) * time.Second,
		ErrorLog:     golog.New(LogWriter{Logger: zap.L()}, "", 0),
	}

	if err := srv.ListenAndServe(); err != nil {
		golog.Panic("Failed to start server", fmt.Sprintf(err.Error()))
	}
}
