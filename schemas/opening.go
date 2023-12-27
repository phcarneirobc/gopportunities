package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct{
	gorm.Model
	Name string
	CPF string
	Location string
	Priority bool
	Number string
	//Salary int64
}

type OpeningResponse struct{
	ID uint `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	CPF string `json:"cpf"`
	Name string `json:"name"`
	Location string `json:"location"`
	Priority bool `json:"priority"`
	Number string `json:"number"`
	//Salary int64 `json:"salary"`
}
	