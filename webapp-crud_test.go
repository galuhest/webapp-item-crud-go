package main

import (
	"fmt"
	"strings"
  "net/http"
  "net/http/httptest"
  "encoding/json"
  "testing"
  "net/url"
  "github.com/galuhest/item-crud-golang"
  "github.com/julienschmidt/httprouter"
)


func ItemFactory()	*MockResponse {
	db := crud.ConnectDb()
  defer db.Close()
  response := crud.CreateItem(db, "item N")
  js_response := []byte(response)
  data := &MockResponse{Status: "", Payload : nil}
  err := json.Unmarshal(js_response, data)
  if err != nil {
  	panic("cant create data for test")
  }
  return data
}

type MockResponse struct {
	Status string `json:"status"` 
	Payload map[string]string `json:"payload,omitempty"`
}

type HandleTester func(
    method string,
    urls string,
    params url.Values,
) *httptest.ResponseRecorder

// Given the current test runner and an http.Handler, generate a
// HandleTester which will test its given input against the
// handler.

func GenerateHandleTester(
    t *testing.T,
    handleFunc http.Handler,
) HandleTester {

    // Given a method type ("GET", "POST", etc) and
    // parameters, serve the response against the handler and
    // return the ResponseRecorder.

    return func(
        method string,
        urls string,
        params url.Values,
    ) *httptest.ResponseRecorder {

        req, err := http.NewRequest(
            method,
            urls,
            strings.NewReader(params.Encode()),
        )
        if err != nil {
            t.Errorf("%v", err)
        }
        req.Header.Set(
            "Content-Type",
            "application/x-www-form-urlencoded; param=value",
        )
        w := httptest.NewRecorder()
        handleFunc.ServeHTTP(w, req)
        return w
    }
}

func TestGetItem(t *testing.T)	{
	router := httprouter.New()
  router.GET("/item/:id", GetHandler)
  test := GenerateHandleTester(t, router)
	id := ItemFactory().Payload["id"]
	urls := fmt.Sprintf("/item/%s",id)
  w := test("GET", urls, url.Values{})
  // fmt.Println(w)
  if w.Code != http.StatusOK	{
      t.Errorf("Get Item is not %v", http.StatusOK)
  }
}

func TestCreateItem(t *testing.T)	{
	router := httprouter.New()
  router.POST("/item", CreateHandler)
  test := GenerateHandleTester(t, router)
	urls := fmt.Sprintf("/item")
	body := url.Values{}
	body.Set("name", "item N")
  w := test("POST", urls, body)
  // fmt.Println(w)
  if w.Code != http.StatusOK	{
      t.Errorf("Create Item is not %v", http.StatusOK)
  }
}

func TestUpdateItem(t *testing.T)	{
	id := ItemFactory().Payload["id"]
	router := httprouter.New()
  router.POST("/item/:id", UpdateHandler)
  test := GenerateHandleTester(t, router)
	urls := fmt.Sprintf("/item/%s",id)
	body := url.Values{}
	body.Set("name", "item N")
  w := test("POST", urls, body)
  // fmt.Println(w)
  if w.Code != http.StatusOK	{
      t.Errorf("Create Item is not %v", http.StatusOK)
  }
}

func TestDeleteItem(t *testing.T)	{
	id := ItemFactory().Payload["id"]
	router := httprouter.New()
  router.POST("/item/:id/delete", DeleteHandler)
  test := GenerateHandleTester(t, router)
	urls := fmt.Sprintf("/item/%s/delete",id)
	body := url.Values{}
  w := test("POST", urls, body)
  // fmt.Println(w)
  if w.Code != http.StatusOK	{
      t.Errorf("Create Item is not %v", http.StatusOK)
  }
}