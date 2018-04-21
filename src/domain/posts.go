package domain

type Post struct {
  ID      int         `json:"id,omitempty"`
  Title   string      `json:"name,omitempty"`
  Body    string      `json:"body,omitempty"`
}