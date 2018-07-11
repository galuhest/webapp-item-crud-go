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

var db *crud.MyDb

func init() {
    db, _ = crud.ConnectDb()
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/item/:id", GetHandler)
    router.POST("/item", CreateHandler)
    router.POST("/item/:id", UpdateHandler)
    router.POST("/item/:id/delete", DeleteHandler)
    defer db.CloseDb()
    log.Fatal(http.ListenAndServe(":8080", router))
}   

// Halaman index untuk root
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Error(status, err_msg string) []byte {

    payload := make(map[string]string)
    payload["err_msg"] = err_msg
    response := crud.Response{Status : "500", Payload : payload}
    js_response, _ := json.Marshal(response)
    return js_response
}

// Handler untuk mengambil nama user
func GetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)   {
    temp := ps.ByName("id") 
    id, err := strconv.Atoi(temp)
    if err != nil {
        js_response := Error("500", err.Error())
        w.Header().Set("Content-Type", "application/json")
        w.Write(js_response) 
        return
    }
    response, err := db.GetItem(id)
    if err != nil {
        js_response := Error("500", err.Error())
        w.Header().Set("Content-Type", "application/json")
        w.Write(js_response) 
        return
    }
    js_response, err := json.Marshal(response)
    if err != nil {
        js_response := Error("500", err.Error())
        w.Header().Set("Content-Type", "application/json")
        w.Write(js_response) 
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js_response) 
}

// Handler untuk membuat user baru
func CreateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
    err := r.ParseForm()
    if err != nil {
        js_response := Error("500", err.Error())
        w.Header().Set("Content-Type", "application/json")
        w.Write(js_response) 
        return
    }
    v := r.Form
    name := v.Get("name")
    response, err := db.CreateItem(name)
    if err != nil {
        js_response := Error("500", err.Error())
        w.Header().Set("Content-Type", "application/json")
        w.Write(js_response) 
        return
    }

    js_response, err := json.Marshal(response)
    if err != nil {
        js_response := Error("500", err.Error())
        w.Header().Set("Content-Type", "application/json")
        w.Write(js_response) 
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js_response) 
}

// Handler untuk mengganti nama user 
func UpdateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
    err := r.ParseForm()
    if err != nil {
        js_response := Error("500", err.Error())
        w.Header().Set("Content-Type", "application/json")
        w.Write(js_response) 
        return
    }
    v := r.Form
    temp := ps.ByName("id") 
    id, err := strconv.Atoi(temp)
    name := v.Get("name")
    response, err := db.UpdateItem(id, name)
    if err != nil {
        js_response := Error("500", err.Error())
        w.Header().Set("Content-Type", "application/json")
        w.Write(js_response) 
        return
    }
    js_response, err := json.Marshal(response)
    if err != nil {
        js_response := Error("500", err.Error())
        w.Header().Set("Content-Type", "application/json")
        w.Write(js_response) 
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js_response) 
}

// Handler untuk menghapus user
func DeleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
    err := r.ParseForm()
    if err != nil {
        js_response := Error("500", err.Error())
        w.Header().Set("Content-Type", "application/json")
        w.Write(js_response) 
        return
    }
    temp := ps.ByName("id") 
    id, err := strconv.Atoi(temp)
    response, err := db.DeleteItem(id)
    if err != nil {
        js_response := Error("500", err.Error())
        w.Header().Set("Content-Type", "application/json")
        w.Write(js_response) 
        return
    }
    js_response, err := json.Marshal(response)
    if err != nil {
        js_response := Error("500", err.Error())
        w.Header().Set("Content-Type", "application/json")
        w.Write(js_response) 
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js_response) 
}