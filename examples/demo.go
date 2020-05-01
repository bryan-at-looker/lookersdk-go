package main

import (
	"context"
	"log"

	"github.com/bryan-at-looker/lookersdkgo"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := lookersdkgo.NewAPIClient()
	ctx := lookersdkgo.LoginContext(context.TODO(), client)
	me, _, err := client.UserApi.Me(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*me.DisplayName)
}
