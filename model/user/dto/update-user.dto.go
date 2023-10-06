package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id           primitive.ObjectID `json:"_id,omitempty"`
	FullName     string             `json:"fullName"`
	Email        string             `json:"email"`
	Password     string             `json:"password,omitempty"`
	MobileNumber string             `json:"mobileNumber" `
	Role         string             `json:"role"`
}
