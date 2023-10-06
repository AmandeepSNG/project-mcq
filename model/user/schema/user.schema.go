package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName     string             `json:"fullName" bson:"fullName"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	MobileNumber string             `json:"mobileNumber" bson:"mobileNumber"`
	Role         string             `json:"role" bson:"role"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
}
