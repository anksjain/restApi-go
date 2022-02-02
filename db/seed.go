package db

import (
	"restsample/db/models"
)

type UserDb struct {
	Users []models.User
}

// inserting custom data
func Seed() *UserDb {
	userDb := &UserDb{}
	var users []models.User
	users = append(users, models.User{
		Id:      1,
		FName:   "Ankit",
		City:    "LA",
		Phone:   1234567890,
		Height:  5.7,
		Married: false,
	})
	users = append(users, models.User{
		Id:      2,
		FName:   "Mehak",
		City:    "Tokya",
		Phone:   3656748374,
		Height:  5.5,
		Married: false,
	})
	users = append(users, models.User{
		Id:      3,
		FName:   "Joey",
		City:    "Pune",
		Phone:   6736467374,
		Height:  5.1,
		Married: true,
	})
	users = append(users, models.User{
		Id:      4,
		FName:   "Chandler",
		City:    "Mumbai",
		Phone:   346376848,
		Height:  7.1,
		Married: true,
	})
	users = append(users, models.User{
		Id:      5,
		FName:   "John",
		City:    "California",
		Phone:   7343767388,
		Height:  2.7,
		Married: false,
	})
	users = append(users, models.User{
		Id:      6,
		FName:   "Vikram",
		City:    "Utsonomiya",
		Phone:   7487484783,
		Height:  7.1,
		Married: true,
	})
	users = append(users, models.User{
		Id:      7,
		FName:   "Shubham",
		City:    "Oska",
		Phone:   8439893984,
		Height:  6.1,
		Married: false,
	})
	userDb.Users = users
	return userDb

}
