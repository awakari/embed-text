package main

import (
	"fmt"
	apiGrpc "github.com/awakari/embed-text/api/grpc"
	"github.com/awakari/embed-text/config"
	"github.com/awakari/embed-text/service"
	"github.com/knights-analytics/hugot"
	"github.com/knights-analytics/hugot/options"
	"log/slog"
	"os"
)

func main() {

	slog.Info("starting...")
	cfg, err := config.NewConfigFromEnv()
	if err != nil {
		panic(err)
	}
	opts := slog.HandlerOptions{
		Level: slog.Level(cfg.Log.Level),
	}
	log := slog.New(slog.NewTextHandler(os.Stdout, &opts))

	session, err := hugot.NewORTSession(options.WithExecutionMode(true))
	if err != nil {
		panic(err)
	}
	defer session.Destroy()

	feCfg := hugot.FeatureExtractionConfig{
		ModelPath:    cfg.Model.Path,
		Name:         cfg.Model.Name,
		OnnxFilename: cfg.Model.FileName,
	}

	pipeline, err := hugot.NewPipeline(session, feCfg)
	if err != nil {
		panic(err)
	}

	svc := service.New(pipeline)
	svc = service.NewLogging(svc, log)

	log.Info(fmt.Sprintf("starting to listen the API @ port #%d...", cfg.Api.Port))
	err = apiGrpc.Serve(svc, cfg.Api.Port)
	if err != nil {
		panic(err)
	}
}
