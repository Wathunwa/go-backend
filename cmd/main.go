package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Wathunwa/go-backend.git/cmd/api"
	mydb "github.com/Wathunwa/go-backend.git/db"
)

func main() {

	ctx := context.Background()

	pool, err := mydb.Connect(ctx, "postgres://postgres:fqu5gx7U7mEiMnhIlrlRIA1O4OoQFW2qhmEDuFtTq8rqfcgkcDMCCvor4T4uMsWf@157.245.202.214:27001/postgres")

	if err != nil {
		panic(err)
	}

	fmt.Println("Connect db success ! ")

	defer mydb.Close(pool)

	server := api.NewAPIServer(":22001", pool) // Send pool for connect
	server.Run()                               // สั่ง Run
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
