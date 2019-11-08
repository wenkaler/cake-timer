package main

import (
	"flag"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	kitlog "github.com/wenkaler/kit/log"
	"github.com/wenkaler/kit/log/level"
	"github.com/wenkaler/xfreehack/storage"
)

type configurate struct {
	ServiceName string `envconfig:"service_name" default:"CakeTimer"`
	PathDB      string `envconfig:"path_db" default:"/db/cake_timer.db"`
	Telegram    struct {
		Token       string `envconfig:"telegram_token"`
		UpdateTimer int    `envconfig:"telegram_update_timer" default:"60"`
	}
}

var serviceVersion = "dev"

func main() {
	version := flag.Bool("version", false, "print service version and exit")
	flag.Parse()

	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)
	log.SetOutput(kitlog.NewStdlibAdapter(logger))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)

	if *version {
		level.Info(logger).Log("version", serviceVersion)
		os.Exit(0)
	}
	var cfg configurate
	err := envconfig.Process("", &cfg)
	if err != nil {
		level.Error(logger).Log("msg", "failed to load configuration", "err", err)
		os.Exit(1)
	}

	s, err := storage.New(cfg.PathDB, logger)
	if err != nil {
		level.Error(logger).Log("msg", "failed create storage", "err", err)
		os.Exit(1)
	}

}
