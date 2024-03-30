package repository

import (
	"context"
	"cronos/internal/domain/wfm/entity"
	"cronos/internal/infrastructure/database"
	entity2 "cronos/pkg/entity"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type WFMRepositoryImpl struct{}

func NewWFMRepository() *WFMRepositoryImpl {
	return &WFMRepositoryImpl{}
}

func (r *WFMRepositoryImpl) Save(ticket entity.TicketInput) error {
	collection := database.GetMongoDB().Collection("tickets")
	ctx := context.Background()

	_, err := collection.InsertOne(ctx, ticket)
	if err != nil {
		fmt.Println("Err:", err)
		return err
	}

	return nil
}

func (r *WFMRepositoryImpl) FindOneIdTikect(idtikect int) (*entity.TicketInput, error) {
	collection := database.GetMongoDB().Collection("tickets")
	ctx := context.Background()
	var ticket entity.TicketInput
	err := collection.FindOne(ctx, bson.M{"idtikect": idtikect}).Decode(&ticket)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return nil, nil
		}
		return nil, err
	}
	return &ticket, nil
}

func (r *WFMRepositoryImpl) Update(ticket entity.TicketInput, id entity2.ID) error {
	collection := database.GetMongoDB().Collection("tickets")
	ctx := context.Background()

	update := bson.M{
		"$set": ticket,
	}

	filter := bson.M{"id": id}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println("Erro ao atualizar o ticket:", err)
		return err
	}

	return nil
}

func (r *WFMRepositoryImpl) UpsertTicket(ticket entity.TicketInput) error {
	tikectExist, err := r.FindOneIdTikect(ticket.IDTikect)
	if err != nil {
		log.Println(err)
	}

	if tikectExist == nil {
		err := r.Save(ticket)
		if err != nil {
			return err
		}
	}

	if tikectExist != nil {
		err := r.Update(ticket, tikectExist.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
