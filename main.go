package main

import (
	"context"
	"github.com/farwydi/cleanwhale/config"
	"github.com/farwydi/cleanwhale/log"
	"github.com/farwydi/cleanwhale/metrics"
	"github.com/farwydi/cleanwhale/transport/whalehttp"
	"github.com/farwydi/cleanwhale/wave"
	"github.com/farwydi/thunderstorm/domain"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	var cfg domain.Config
	err := config.LoadConfigs(&cfg, "config.yml")
	if err != nil {
		log.Fatal(err)
	}

	logger, err := log.NewLogger(cfg.Project)
	if err != nil {
		log.Fatal(err)
	}

	mlog := logger.Named("main")

	metrics.RegisterMetrics("", mlog)

	app, cleanup, err := setup(cfg)
	if err != nil {
		logger.Fatal("fail setup", zap.Error(err))
	}

	defer cleanup()

	w := wave.NewWave(context.Background(), mlog)

	w.AddSever(whalehttp.NewHTTPServer(
		cfg.Transport.Webhook, logger.Named("webhook"), app.webhook))

	w.Run()
}

type application struct {
	webhook http.Handler
}

func newApplication(webhook http.Handler) application {
	return application{
		webhook: webhook,
	}
}
