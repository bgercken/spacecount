package main

import (
  "encoding/json"
  "fmt"
)

type roster struct {
  Craft string `json:"string"`
  Name string `json:"string"`
}

type people struct {
  Roster []roster
  Message string `json:"string"`
  Number int `json:"number"`
}

func main() {

  text := `{"people": [
    {"craft": "ISS", "name": "Sergey Rizhikov"},
    {"craft": "ISS", "name": "Andrey Borisenko"},
    {"craft": "ISS", "name": "Shane Kimbrough"},
    {"craft": "ISS", "name": "Oleg Novitskiy"},
    {"craft": "ISS", "name": "Thomas Pesquet"},
    {"craft": "ISS", "name": "Peggy Whitson"}],
      "message": "success", "number": 6}`

  textBytes := []byte(text)

  var roster []roster

  people1 := people{}

  // Unmarshal from the hardcoded text bytes into a struct
  err := json.Unmarshal(textBytes, &people1)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(people1.Number)

  err = json.Unmarshal(textBytes, &roster)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("roster: %v\n", roster)

  n := people1.Number

  for i := 0; i < n; i++ {
    fmt.Printf("This is %d.\n", i)
  }

  fmt.Printf("people: %v\n", people1)

}
