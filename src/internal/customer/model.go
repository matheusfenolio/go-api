package customer

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
}

type CustomerResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type CustomerResquest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required"`
}
