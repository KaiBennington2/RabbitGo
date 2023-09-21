package users

import (
	"context"
	"errors"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/internal/domain/entity"
	outPorts "github.com/KaiBennington2/RabbitGo/internal/domain/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var _ outPorts.IUserRepository = (*Mongo)(nil)

type Mongo struct {
	db         ports.IExecutorNoSqlDB
	collection string
}

func NewUserMongoRepo(db ports.IExecutorNoSqlDB) *Mongo {
	return &Mongo{
		db:         db,
		collection: CollectionName,
	}
}

func (rpo *Mongo) Save(ctx context.Context, e *entity.User) (*string, error) {
	var err error
	var codeLastInsert *string
	collection := rpo.db.Collection(rpo.collection)

	result, err := collection.InsertOne(ctx, e)
	if err != nil {
		return codeLastInsert, err
	}

	objID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return codeLastInsert, err
	}

	codeLastInsert, err = rpo.getLastID(ctx, objID)
	if err != nil {
		return codeLastInsert, err
	}
	return codeLastInsert, nil
}

func (rpo *Mongo) Delete(ctx context.Context, code *string) (bool, error) {
	collection := rpo.db.Collection(rpo.collection)

	filter := bson.M{"code": code}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}

	if result.DeletedCount <= 0 {
		return false, err
	}
	return result.DeletedCount > 0, nil
}

func (rpo *Mongo) LogicalDelete(ctx context.Context, code *string) (bool, error) {
	collection := rpo.db.Collection(rpo.collection)

	filter := bson.M{"code": code}
	update := bson.M{"$set": bson.M{"date_deleted": time.Now(), "deletedAt": true}}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	if result.ModifiedCount == 0 {
		return false, err
	}
	return result.ModifiedCount > 0, nil
}

func (rpo *Mongo) ExistsById(ctx context.Context, code *string, withoutLogicDeleted ...bool) (bool, error) {
	collection := rpo.db.Collection(rpo.collection)

	filter := bson.M{"code": code}
	if len(withoutLogicDeleted) > 0 && withoutLogicDeleted[0] {
		filter["deletedAt"] = false
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	if count <= 0 {
		return false, err
	}
	return count > 0, nil
}

func (rpo *Mongo) FindByObjectId(ctx context.Context, id primitive.ObjectID) (*entity.User, error) {
	collection := rpo.db.Collection(rpo.collection)

	filter := bson.M{"_id": id}
	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("the user with id %v was not found", id)
		}
		return nil, result.Err()
	}

	var user entity.User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (rpo *Mongo) getLastID(ctx context.Context, objID primitive.ObjectID) (*string, error) {
	var err error
	var codeLastInsert *string

	newEnt, err := rpo.FindByObjectId(ctx, objID)
	if err != nil {
		return codeLastInsert, err
	}

	codeLastInsert = &newEnt.Code
	return codeLastInsert, nil
}
