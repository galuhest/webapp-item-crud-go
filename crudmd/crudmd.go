package main

import (
	"os"
	"fmt"
	"log"
  "strconv"
  "encoding/json"
	"github.com/galuhest/item-crud-golang"
)

type Response struct	{
	Status string
	Payload map[string]string
}

func main() {
  args := os.Args[1:]
  if len(args) < 2 {
  	log.Fatal("Need at least 2 parameters")
  }

  crud := args[0]
  switch crud{
  case "get":
  	GetItem(args[1:])
  case "create":
  	fmt.Println("create")
  case "update":
  	fmt.Println("update")
  case "delete":
  	fmt.Println("delete")
  default:
  	fmt.Println("invalid argument")
  }
}

func GetItem(args []string)	{
	id, err := strconv.Atoi(args[0])
	if err != nil	{
		log.Fatal("invalid id parameter")
	}
	db := crud.ConnectDb()
	crud_response := crud.GetItem(db, id)
	data := &Response{Status: "", Payload : nil}
  err = json.Unmarshal([]byte(crud_response), data)
  if err != nil {
  	panic("cant create data")
  }
  fmt.Println(data.Payload["name"])
}

func CreateItem(args []string)	{

}