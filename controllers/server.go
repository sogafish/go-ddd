package controllers

import (
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  _ "io/ioutil"
  "log"
  "go-go/config"
  "go-go/models"
  "net/http"
  _ "strconv"
)


func root(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Running")
}

func fetchItems(w http.ResponseWriter, r *http.Request) {
  var items []models.Item
  models.GetAllItems(&items)
  responseBody, err := json.Marshal(items)
  if err != nil {
    log.Fatal(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(responseBody)
}

func StartWebServer() error {
  fmt.Println("Rest API with Mux Routers")
  router := mux.NewRouter().StrictSlash(true)

  router.HandleFunc("/", root)
  router.HandleFunc("/items", fetchItems).Methods("GET")

  return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.ServerPort), router)
}
