package main

import "encoding/json"

type jSON struct {
	//In order to allow the JSON library to access the Name variable, the first letter must be capital
	Name string `json:"name"`
}

func main() {
	var data []byte
	json.Unmarshal(data, &jSON{})
}
