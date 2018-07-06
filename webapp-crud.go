package main

import (
  "fmt"
  "strconv"
  "net/http"
  "github.com/galuhest/item-crud-golang"
  "github.com/julienschmidt/httprouter"
  "log"
)

func main() {
  router := httprouter.New()
  router.GET("/", Index)
  router.GET("/item/:id", GetHandler)
  router.POST("/item", CreateHandler)
  router.POST("/item/:id", UpdateHandler)
  router.POST("/item/:id/delete", DeleteHandler)

  log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}


func GetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)	{
  temp := ps.ByName("id") 
  id, err := strconv.Atoi(temp)
  if err != nil {
    log.Fatal("cant convert str to int")
  }
  db := crud.ConnectDb()
  defer db.Close()
  response := crud.GetItem(db, id)
  js_response := []byte(response)
  w.Header().Set("Content-Type", "application/json")
  w.Write(js_response) 
}

func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
  err := r.ParseForm()
  if err != nil {
    panic(err)
  }
  v := r.Form
  name := v.Get("name")
  db := crud.ConnectDb()
  defer db.Close()
  response := crud.CreateItem(db, name)
  js_response := []byte(response)
  w.Header().Set("Content-Type", "application/json")
  w.Write(js_response) 
}

func UpdateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
  err := r.ParseForm()
  if err != nil {
    panic(err)
  }
  v := r.Form
  temp := ps.ByName("id") 
  id, err := strconv.Atoi(temp)
  name := v.Get("name")
  db := crud.ConnectDb()
  defer db.Close()
  response := crud.UpdateItem(db, id, name)
  js_response := []byte(response)
  w.Header().Set("Content-Type", "application/json")
  w.Write(js_response) 
}

func DeleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
  err := r.ParseForm()
  if err != nil {
    panic(err)
  }
  temp := ps.ByName("id") 
  id, err := strconv.Atoi(temp)
  db := crud.ConnectDb()
  defer db.Close()
  response := crud.DeleteItem(db, id)
  js_response := []byte(response)
  w.Header().Set("Content-Type", "application/json")
  w.Write(js_response) 
}