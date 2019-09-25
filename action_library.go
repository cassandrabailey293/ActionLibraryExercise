package main

import (
    "fmt"
    "encoding/json"
)

type Action struct {
    Action string `json:"action"`
    Time json.Number `json:"time"`
}

func main() {
  act := `{ "action": "jump", "time": 100 }`
  addAction(act)
}

func addAction(s string) (error) {
  action := &Action{}
  err := json.Unmarshal([]byte(s), action)
    fmt.Println(err)
    fmt.Println(action)
    return err
}
