package main

import (
    "fmt"
    "log"
    "flag"
    "strconv"
    "strings"
    "github.com/joho/godotenv"
    "github.com/galuhest/item-crud-go"
)

type Response struct    {
    Status string
    Payload map[string]string
}

var (
    envPath = flag.String("env", "", "Path to .env file (default is root)")
)

func main() {
    var err error
    flag.Parse()
    // fmt.Println(*envPath)
    if *envPath != ""   {
        err = godotenv.Load(*envPath)
    } else  {
        err = godotenv.Load()
    }

    if err != nil   {
        log.Fatal("Fail to load .env")
    }

    var args []string
    if flag.NArg() < 2 {
        log.Fatal("Need at least 2 parameters")
    } else  {
        args = flag.Args()
    }

    crud := strings.ToLower(args[0])
    switch crud{
    case "get":
        GetItem(args[1:])
    case "create":
        CreateItem(args[1:])
    case "update":
        UpdateItem(args[1:])
    case "delete":
        DeleteItem(args[1:])
    default:
        fmt.Println("invalid argument")
    }
}

func GetItem(args []string) {
    if len(args) != 1   {
        log.Fatal("Invalid amount of parameters")
    }
    id, err := strconv.Atoi(args[0])
    if err != nil   {
        log.Fatal("invalid id parameter")
    }
    db, err := crud.ConnectDb()
    if err != nil {
        panic("cant connect db")
    }
    crud_response, err := db.GetItem(id)
    if err != nil {
        panic("cant gt item")
    }
    fmt.Println(crud_response.Payload["name"])
}

func CreateItem(args []string)  {
    if len(args) != 1   {
        log.Fatal("Invalid amount of parameters")
    }
    name := args[0]
    db, err := crud.ConnectDb()
    if err != nil {
        panic("cant create data")
    }
    crud_response, err := db.CreateItem(name)
    if err != nil {
        panic("cant create data")
    }
    fmt.Println(crud_response.Payload["id"])
}

func UpdateItem(args []string)  {
    if len(args) != 2   {
        log.Fatal("Invalid amount of parameters")
    }
    id, err := strconv.Atoi(args[0])
    if err != nil   {
        log.Fatal("invalid id parameter")
    }
    name := args[1]
    db,err := crud.ConnectDb()
    if err != nil {
        panic("cant create data")
    }
    crud_response, err := db.UpdateItem(id, name)
    if err != nil {
        panic("cant create data")
    }
    fmt.Println(crud_response.Status)
}

func DeleteItem(args []string)  {
    if len(args) != 1   {
        log.Fatal("Invalid amount of parameters")
    }
    id, err := strconv.Atoi(args[0])
    if err != nil   {
        log.Fatal("invalid id parameter")
    }
    db, err := crud.ConnectDb()
    crud_response, err := db.DeleteItem(id)
    if err != nil {
        panic("cant create data")
    }
    fmt.Println(crud_response.Status)
}