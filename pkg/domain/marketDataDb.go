package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dibyendu/trading_platform/lib/constants"
	"github.com/dibyendu/trading_platform/lib/errs"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type MarketDataRepositoryDb struct {
	client      *mongo.Client
	redisClient *redis.Client
	database    string
	collection  map[string]string
}

func NewMarketDataRepositoryDb(dbClient *mongo.Client, redisClient *redis.Client, database string, collection map[string]string) MarketDataRepositoryDb {
	return MarketDataRepositoryDb{
		client:      dbClient,
		redisClient: redisClient,
		database:    database,
		collection:  collection,
	}
}

const BinanceAPIURL = "https://api.binance.com/api/v3/ticker/24hr"
var httpClient *http.Client

func (s MarketDataRepositoryDb) GetMarketData (ctx context.Context, symbol string) (*MarketData, *errs.AppError) {
    url := fmt.Sprintf("%s?symbol=%s", BinanceAPIURL, symbol)
    
    resp, err := httpClient.Get(url)
    if err != nil {
        return nil, errs.NewUnexpectedError(constants.UNEXPECTED_ERROR)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errs.NewUnexpectedError(constants.UNEXPECTED_ERROR)
    }

    var marketData MarketData
    if err := json.NewDecoder(resp.Body).Decode(&marketData); err != nil {
        return nil, errs.NewUnexpectedError(constants.UNEXPECTED_ERROR)
    }

    return &marketData, nil
}