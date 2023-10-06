package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Options struct {
	Option string `json:"option" bson:"option"`
	Status bool   `json:"status" bson:"status"`
}

type QuestionSchema struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	QuestionId   string             `json:"questionId" bson:"questionId"`
	QuestionType string             `json:"questionType" bson:"questionType"`
	Question     string             `json:"question" bson:"question"`
	Options      []*Options         `json:"options" bson:"options"`
	Status       bool               `json:"status" bson:"status"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}

const QUESTION_SCHEMA_NAME = "questions"
