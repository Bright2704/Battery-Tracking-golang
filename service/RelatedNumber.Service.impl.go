package service

import (
	"api/Battery-Tracking-go/entity"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RelatedNumberServiceImpl struct {
	relatedNumbercollection *mongo.Collection
	ctx                     context.Context
}

// GetRelatedNumber implements RelatedNumberService
func (*RelatedNumberServiceImpl) GetRelatedNumber(*string) (*entity.RelatedNumber, error) {
	panic("unimplemented")
}

func NewRelatedNumberService(relatedNumbercollection *mongo.Collection, ctx context.Context) RelatedNumberService {
	return &RelatedNumberServiceImpl{
		relatedNumbercollection: relatedNumbercollection,
		ctx:                     ctx,
	}
}

func (u *RelatedNumberServiceImpl) CreateRelatedNumber(relatedNumber *entity.RelatedNumber) error {
	_, err := u.relatedNumbercollection.InsertOne(u.ctx, relatedNumber)
	return err
}

func (u *RelatedNumberServiceImpl) GetRelatedNUmber(relatedProcess *string) (*entity.RelatedNumber, error) {
	var user *entity.RelatedNumber
	query := bson.D{bson.E{Key: "relatedProcess", Value: relatedProcess}}
	err := u.relatedNumbercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *RelatedNumberServiceImpl) GetAll() ([]*entity.RelatedNumber, error) {
	var users []*entity.RelatedNumber
	cursor, err := u.relatedNumbercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user entity.RelatedNumber
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("document not found")
	}
	return users, nil
}

func (u *RelatedNumberServiceImpl) UpdateRelatedNumber(user *entity.RelatedNumber) error {
	filter := bson.D{primitive.E{Key: "relatedprocess", Value: user.RelatedProcess}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "relatedProcess", Value: user.RelatedProcess}}}}
	result, _ := u.relatedNumbercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *RelatedNumberServiceImpl) DeleteRelatedNumber(relatedprocess *string) error {
	filter := bson.D{primitive.E{Key: "relatedprocess", Value: relatedprocess}}
	result, _ := u.relatedNumbercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
