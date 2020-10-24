package main

import (
	"API/comdirect"
	"fmt"
)

func main()  {
	fmt.Println("Hallo welt");
	creds := &comdirect.Credentials{
		Username:     "test",
		Password:     "password",
		ClientId:     "clientId",
		ClientSecret: "secret",
	}
	client := comdirect.NewClient(*creds)
	err := client.RequestToken()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(client.Credentials.ClientId)

}
