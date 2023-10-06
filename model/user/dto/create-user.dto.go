package dto

type CreateUserDto struct {
	FullName     string `json:"fullName"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	MobileNumber string `json:"mobileNumber"`
	Role         string `json:"role"`
}
