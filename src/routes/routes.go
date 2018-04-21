package routes

import (
  "fmt"
  "net/http"
  "encoding/json"
  "strconv"

  "github.com/mtronix/golang-rest/src/domain"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprintf(w, "<h1>Hello World</h1>\n<br /><br />\n<a href=\"/posts\">Posts</a> ")
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
  var posts []domain.Post

  for i := 1; i <= 1000; i++ {
    iAsString := strconv.Itoa(i)

    posts = append(posts, domain.Post{ID: i, Title: "Title of post: " + iAsString, Body: "Body of post with number: " + iAsString})
  }

  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(posts)
}