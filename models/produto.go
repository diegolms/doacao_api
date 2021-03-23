package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Produto struct {
	ID            primitive.ObjectID    `bson:"_id"`
	Nome          string         		`bson:"nome"`
	Validade 	  time.Time 	   		`bson:"validade"`
}