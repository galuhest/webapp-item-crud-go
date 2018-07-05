package main

import (
  "encoding/json"
  "net/http"
  "github.com/galuhest/item-crud-golang"
  "github.com/julienschmidt/httprouter"
  "log"
)

func main() {
  router := httprouter.New()
  router.GET("/", Index)
  router.GET("/item/:id", Hello)
  router.POST("/item", Hello)
  router.POST("/item/:id", Hello)
  router.POST("/item/:id/delete", Hello)

  log.Fatal(http.ListenAndServe(":8080", router))
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
      http.NotFound(w, r)
      return
    }
    fn(w, r, m[2])
  }
}

func getHandler(w http.ResponseWriter, r *http.Request, title string)	{

  js = crud.GetItem()
  w.Header().Set("Content-Type", "application/json")
  w.Write(js) 
}