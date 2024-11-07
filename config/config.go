package config

import (
	"os"

	"github.com/dibyendu/trading_platform/lib/constants"
	"github.com/dibyendu/trading_platform/pkg/client/db"
	"github.com/dibyendu/trading_platform/pkg/client/redis"
)

type AppConfig struct {
	DB    *db.Config
	Redis *redis.Config
}

var (
	appConfig AppConfig
)

func Init() *AppConfig {
	userCollection := make(map[string]string)
	orderCollection := make(map[string]string)
	positionCollection := make(map[string]string)
	tradingHistoryCollection := make(map[string]string)

	userCollection["user"] = constants.USER_COLLECTION
	orderCollection["order"] = constants.ORDER_COLLECTION
	positionCollection["position"] = constants.POSITION_COLLECTION
	tradingHistoryCollection["tradeHistory"] = constants.TRADE_HISTORY

	appConfig = AppConfig{
		DB: &db.Config{
			Host:           os.Getenv("DB_HOST"),
			Port:           os.Getenv("DB_PORT"),
			MaxPool:        os.Getenv("MAX_POOL"),
			Database:       os.Getenv("DB_NAME"),
			UserCollection: userCollection,
			OrderCollection: orderCollection,
			PositionCollection: positionCollection,
			TradingHistoryCollection: tradingHistoryCollection,
		},
		Redis: &redis.Config{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASS"),
		},
	}
	return &appConfig
}
