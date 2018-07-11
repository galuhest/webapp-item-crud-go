package main

import (
    "fmt"
    "strconv"
    "encoding/json"
    "net/http"
    "github.com/galuhest/item-crud-go"
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

// Halaman index untuk root
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

// Handler untuk mengambil nama user
func GetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)   {
    temp := ps.ByName("id") 
    id, err := strconv.Atoi(temp)
    if err != nil {
        log.Fatal("cant convert str to int")
    }
    db, err:= crud.ConnectDb()
    if err != nil {
        log.Fatal("cant open database")
    }
    defer db.CloseDb()
    response, err := db.GetItem(id)
    if err != nil {
        return
    }
    js_response, err := json.Marshal(response)
    if err != nil {
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js_response) 
}

// Handler untuk membuat user baru
func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
    err := r.ParseForm()
    if err != nil {
        panic(err)
    }
    v := r.Form
    name := v.Get("name")
    db, err := crud.ConnectDb()
    if err != nil {
        log.Fatal(err.Error())
    }
    defer db.CloseDb()
    response, err := db.CreateItem(name)
    js_response, err := json.Marshal(response)
    if err != nil {
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js_response) 
}

// Handler untuk mengganti nama user 
func UpdateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
    err := r.ParseForm()
    if err != nil {
        panic(err)
    }
    v := r.Form
    temp := ps.ByName("id") 
    id, err := strconv.Atoi(temp)
    name := v.Get("name")
    db, err := crud.ConnectDb()
    if err != nil {
        log.Fatal(err.Error())
    }
    defer db.CloseDb()
    response, err := db.UpdateItem(id, name)
    if err != nil {
        return    
    }
    js_response, err := json.Marshal(response)
    if err != nil {
        return    
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js_response) 
}

// Handler untuk menghapus user
func DeleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
    err := r.ParseForm()
    if err != nil {
        panic(err)
    }
    temp := ps.ByName("id") 
    id, err := strconv.Atoi(temp)
    db, err := crud.ConnectDb()
    if err != nil {
        return    
    }
    defer db.CloseDb()
    response, err := db.DeleteItem(id)
    if err != nil {
        return    
    }
    js_response, err := json.Marshal(response)
    if err != nil {
        return    
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js_response) 
}