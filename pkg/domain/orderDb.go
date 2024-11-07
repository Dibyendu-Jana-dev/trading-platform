package domain

import (
	"context"
	"time"

	"github.com/dibyendu/trading_platform/lib/constants"
	"github.com/dibyendu/trading_platform/lib/errs"
	"github.com/dibyendu/trading_platform/lib/logger"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepositoryDb struct {
	client     *mongo.Client
	redisClient *redis.Client
	database   string
	collection map[string]string
}

func NewOrderRepositoryDb(dbClient *mongo.Client, redisClient *redis.Client, database string, collection map[string]string) OrderRepositoryDb {
	return OrderRepositoryDb{
		client:      dbClient,
		redisClient: redisClient,
		database:    database,
		collection:  collection,
	}
}

func (n OrderRepositoryDb) PlaceOrder(ctx context.Context, req Order) (*Order, *errs.AppError) {
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	// Insert the order document into MongoDB
	result, err := n.client.Database(n.database).Collection(n.collection["order"]).InsertOne(ctx, req)
			if err != nil {
				logger.Error("error inserting user log: " + err.Error())
				return nil, errs.NewUnexpectedError(constants.UNEXPECTED_ERROR)
			}

	// Set the ID field on the order object
	req.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return &req, nil
}

func (n OrderRepositoryDb) DeleteOrder(ctx context.Context, orderId string)(*DeleteOrderResponse, *errs.AppError){
	filter := bson.M{"_id": orderId}
	result, err := n.client.Database(n.database).Collection(n.collection["order"]).DeleteOne(ctx, filter)
    if err != nil {
        logger.Error("error inserting user log: " + err.Error())
		return nil, errs.NewUnexpectedError(constants.UNEXPECTED_ERROR)
    }
    if result.DeletedCount == 0 {
        return &DeleteOrderResponse{}, errs.NewNotFoundError("order not found")
    }
    return &DeleteOrderResponse{
		Message:constants.ORDER_DELETED, 
		OrderID: orderId,
	}, nil
}