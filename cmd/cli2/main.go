package main

import (
  "encoding/json"
  "fmt"
)

type Birds struct {
  Bird [] Bird 
}

type Dimensions struct {
  Height int
  Width int
}

type Bird struct {
  Species string
  Description string
  Dimensions Dimensions
}

func main() {

  // birdsJson := `{"birds": [
    // {"species": "pigeon","description": "likes to perch on rocks"},
    // {"species": "pigeon","description": "likes to perch on rocks"},
    // {"species": "pigeon","description": "likes to perch on rocks"}],
     // "number": 3}`

  birdsJson := `[
    {
      "species": "pigeon",
      "description": "likes to perch on rocks",
      "dimensions": {
        "height": 24,
        "width": 10
      }
    },
    {
      "species": "eagle",
      "description": "birds of prey",
      "dimensions": {
        "height": 24,
        "width": 10
      }

    },
    {
      "species": "owl",
      "description": "nocturnal bird of prey",
      "dimensions": {
        "height": 24,
        "width": 10
      }
    }]
  `

  birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
  var onebird Bird	

  err := json.Unmarshal([]byte(birdJson), &onebird)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("Species: %s, Description: %s\n", onebird.Species, onebird.Description)

  // var birds Birds

  var birds []Bird
  err = json.Unmarshal([]byte(birdsJson), &birds)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("birds: %v\n", birds)

  for i := 0; i < len(birds); i++ {
    fmt.Printf("bird[%d]: %v\n", i+1, birds[i])
  }
}
