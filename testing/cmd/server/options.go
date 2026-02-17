package main

import (
	"flag"
	"os"

	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/log"
)

type Cfg struct {
	Debug    bool
	GRPCAddr string
	HttpAddr string
}

func HandleCfg() *Cfg {
	cfg := Cfg{}

	flag.BoolVar(&cfg.Debug, "debug", true, "Enable Debug into the server")
	flag.StringVar(&cfg.GRPCAddr, "grpc-addr", ":9091", "GRPC address to listen on")
	flag.StringVar(&cfg.HttpAddr, "http-addr", ":8081", "HTTP address to listen on")

	return &cfg
}

func SetupLogger(cfg *Cfg) log.Logger {
	var logger log.Logger

	{
		logger = log.NewLogfmtLogger(os.Stderr)

		if cfg.Debug {
			logger = level.NewFilter(logger, level.AllowDebug())
		} else {
			logger = level.NewFilter(logger, level.AllowInfo())
		}

		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	return logger
}
