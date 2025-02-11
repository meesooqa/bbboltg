package main

import (
	"context"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/meesooqa/bbboltg/app/config"
	"github.com/meesooqa/bbboltg/app/telegram"
)

type options struct {
	Conf string `short:"f" long:"conf" env:"BBBOLTG_CONF" default:"etc/bbboltg.yml" description:"config file (yml)"`
}

func main() {
	log.Println("bbboltg")
	godotenv.Load()
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		os.Exit(1)
	}

	var err error
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("[ERROR] can't obtain logger, %v", err)
	}
	defer logger.Sync()

	conf, err := config.Load(opts.Conf)
	if err != nil {
		logger.Error("can't load config", zap.Any("config", opts.Conf), zap.Error(err))
	}

	p := telegram.NewProcessor(conf)
	if err := p.Do(context.Background(), logger); err != nil {
		logger.Error("processor failed", zap.Error(err))
	}
}
