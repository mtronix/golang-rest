package main

import (
  "net/http"
  "log"

  "github.com/gorilla/mux"
  "github.com/mtronix/golang-rest/src/routes"
)

func main() {
  router := mux.NewRouter()

  router.HandleFunc("/", routes.IndexPage)
  router.HandleFunc("/posts", routes.GetPosts)

  log.Fatal(http.ListenAndServe(":8080", router))
}