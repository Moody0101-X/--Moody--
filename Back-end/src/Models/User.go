package models

import (
	"time"
)

type User_ struct {
	Id			int 	`json:"id"`
	Email		string 	`json:"email"`
	Name		string 	`json:"name"`
	Balance 	int 	`json:"balance"`
	PhoneNumber string 	`json:"phone_number"`
}

type User struct {
	Id			int 	`json:"id"`
	Email		string 	`json:"email"`
	Pwd     	string 	`json:"pwd"`	
	Name		string 	`json:"name"`
	Balance 	int 	`json:"balance"`
	PhoneNumber string 	`json:"phone_number"`
	JWT   string  `json:"jwt"`
}

// Seller and buyer are two struct that are used in handling sales operations between users......
type Seller struct {
	Identity User_ `json:"user"`
}

type Buyer struct {
	Identity     User_ 		  `json:"user"`
	JWT          string 	  `json:"jwt"`
	TimeStamp    time.Time    `json:"ts"`
	Target		 Product      `json:"product"`
}