package domain

import (
	"context"

	"github.com/dibyendu/trading_platform/lib/constants"
	"github.com/dibyendu/trading_platform/lib/errs"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TradingHistoryRepositoryDb struct {
	client      *mongo.Client
	redisClient *redis.Client
	database    string
	collection  map[string]string
}

func NewTradingHistoryRepositoryDb(dbClient *mongo.Client, redisClient *redis.Client, database string, collection map[string]string) TradingHistoryRepositoryDb {
	return TradingHistoryRepositoryDb{
		client:      dbClient,
		redisClient: redisClient,
		database:    database,
		collection:  collection,
	}
}

func(n TradingHistoryRepositoryDb)GetTradeHistoryByUserID(ctx context.Context, userId string) ([]*TradeHistory, *errs.AppError){
		var trades []*TradeHistory
		filter := bson.M{"user_id": userId}
		cursor, err := n.client.Database(n.database).Collection(n.collection["tradeHistory"]).Find(ctx, filter)
		if err != nil {
			return nil, errs.NewUnexpectedError(constants.UNEXPECTED_ERROR)
		}
		defer cursor.Close(context.TODO())
	
		for cursor.Next(context.TODO()) {
			var trade TradeHistory
			if err := cursor.Decode(&trade); err != nil {
				return nil, errs.NewUnexpectedError(constants.UNEXPECTED_ERROR)
			}
			trades = append(trades, &trade)
		}
		return trades, nil

}
