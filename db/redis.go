package db

import (
	"errors"
	"log"
	"os"

	"context"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var (
	Ctx = context.TODO()	
	client *redis.Client
)

// storing our user information in redis data . 
func init() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file: ", err)
	}

	addr := os.Getenv("REDIS-ADDR")
	client = redis.NewClient(&redis.Options{
		Addr : addr,
		Password : "", 
		DB : 0 ,
	})

	if err := client.Ping(Ctx).Err(); err!= nil {
		log.Fatal("Don't have ping to Database : ", err)
	}


}

func IsValid(givenUser, givenPass string) (bool, error) {
	t := client.Exists(Ctx, givenUser)
	isExist, err := t.Result()
	if err != nil || isExist== 0 {
		return false, err
	}

	s := client.HGet(Ctx, givenUser, "pass")
	pass , err := s.Result()
	if err != nil {
		return false ,err 
	}

	
	if givenPass != pass {
		return false ,nil 
	}
	return true , nil 
}

func CreateUser(username , password, email string) error {
	i := client.Exists(Ctx, username)
	result, err :=i.Result()
	if err != nil {
		return err 
	}
	if result == 1 {
		return errors.New("user already exist")
	}

	i = client.HSet(Ctx, username, "pass" , password, "email" , email)
	if err := i.Err(); err != nil {
		return err
	}
	return nil 
}	

func UpdateUser(username, email, password string) error{
	i := client.Exists(Ctx, username)
	result, err :=i.Result()
	if err != nil {
		return err 
	}
	if result == 0 {
		return errors.New("there is no such a user")
	}

	i = client.Del(Ctx, username)
	if err := i.Err(); err != nil {
		return err
	}
	
	i = client.HSet(Ctx, username,"pass",password, "email", email)
	if err := i.Err(); err != nil {
		return err
	}
	return nil 
}

func DeleteAccount(username string) error {
	i := client.Exists(Ctx, username)
	result, err :=i.Result()
	if err != nil {
		return err 
	}
	if result == 0 {
		return errors.New("there is no such a user")
	}

	i = client.Del(Ctx, username)
	if result, err := i.Result(); err != nil || result== 0{
		return err 
	}
	return nil

}