package repository

import (
	"context"
	"cronos/internal/domain/wfm/entity"
	"cronos/internal/infrastructure/database"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type FlowRepository struct{}

func NewFlowStateRepository() *FlowRepository {
	return &FlowRepository{}
}

func (r *FlowRepository) UpdateOrInsert(flow entity.FlowState) error {
	ifExistFlow := r.FindOneFlowStateId(flow.FlowStateId)
	if ifExistFlow != nil {
		err := r.Update(flow)
		if err != nil {
			return err
		}
		return nil
	}
	err := r.Save(flow)
	if err != nil {
		return err
	}
	return nil
}

func (r *FlowRepository) FindOneFlowStateId(flowStateId int) *entity.FlowState {
	collection := database.GetMongoDB().Collection("flowstates")
	ctx := context.Background()
	var flow entity.FlowState
	err := collection.FindOne(ctx, bson.M{"flowstateid": flowStateId}).Decode(&flow)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return nil
		}
		return nil
	}
	if &flow != nil {
		return &flow
	}
	return nil
}

func (r *FlowRepository) Update(flow entity.FlowState) error {
	log.Println("UPDATE: ", flow.FlowStateId)
	collection := database.GetMongoDB().Collection("flowstates")
	ctx := context.Background()
	update := bson.M{
		"$set": flow,
	}
	filter := bson.M{"flowstateid": flow.FlowStateId}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *FlowRepository) Save(flow entity.FlowState) error {
	log.Println("INSERT:", flow.FlowStateId)
	collection := database.GetMongoDB().Collection("flowstates")
	ctx := context.Background()
	_, err := collection.InsertOne(ctx, flow)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *FlowRepository) FindAll() ([]*entity.FlowState, error) {
	collection := database.GetMongoDB().Collection("flowstates")
	ctx := context.Background()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var flows []*entity.FlowState
	for cursor.Next(ctx) {
		var flow entity.FlowState
		if err := cursor.Decode(&flow); err != nil {
			log.Fatal(err)
			return nil, err
		}
		flows = append(flows, &flow)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return flows, nil
}
