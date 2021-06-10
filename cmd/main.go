package main

import (
	"fmt"
	"log"

	"github.com/MadXIII/todo-app/pkg/handler"

	"github.com/MadXIII/todo-app"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
	fmt.Println("tes")
}

//POST /AUTH/SIGN-UP
//POST /AUTH/SIGN-IN

//GET /LISTS
//GET /LISTS/{ID}
//POST /LISTS
//PUT /LISTS/{ID}
//DELETE /LISTS/{ID}
//GET /LISTS/{ID}/ITEMS
//POST /LISTS/{ID}/ITEMS

//PUT /ITEMS/{ID}
//GET /ITEMS/{ID}
//DELETE /ITEMS/{ID}
