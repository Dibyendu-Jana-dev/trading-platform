package domain

import (
	"context"

	"github.com/dibyendu/trading_platform/lib/constants"
	"github.com/dibyendu/trading_platform/lib/errs"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PositionRepositoryDb struct {
	client      *mongo.Client
	redisClient *redis.Client
	database    string
	collection  map[string]string
}

func NewPositionRepositoryDb(dbClient *mongo.Client, redisClient *redis.Client, database string, collection map[string]string) PositionRepositoryDb {
	return PositionRepositoryDb{
		client:      dbClient,
		redisClient: redisClient,
		database:    database,
		collection:  collection,
	}
}


func (n PositionRepositoryDb) GetUserPositions(ctx context.Context, userID string) ([]*Position, *errs.AppError) {

	var positions []*Position
	filter := bson.M{"user_id": userID}
	cursor, err := n.client.Database(n.database).Collection(n.collection["position"]).Find(ctx, filter)
	if err != nil {
		return nil, errs.NewUnexpectedError(constants.UNEXPECTED_ERROR)
	}

	if err = cursor.All(ctx, &positions); err != nil {
		return nil, errs.NewUnexpectedError(constants.UNEXPECTED_ERROR)
	}

	return positions, nil
}