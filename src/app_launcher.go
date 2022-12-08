package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Constants
const (
	Inputs_path   = "inputs/"
	Programs_path = "src/"
)

// Struct
type config struct {
	Input string `yaml:"input"`
	Code  string `yaml:"code"`
}

// Methods
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// main
func main() {
	// Read yaml file
	file, err1 := os.ReadFile("data.yaml")
	CheckError(err1)

	// Parse yaml data
	var parsedData config
	err2 := yaml.Unmarshal(file, &parsedData)
	CheckError(err2)

	input := Inputs_path + parsedData.Input

	switch parsedData.Code {
	case "day_1":
		Day_1(string(input))
	case "day_2":
		Day_2(string(input))
	case "day_3":
		Day_3(string(input))
	case "day_4":
		Day_4(string(input))
	case "day_5":
		Day_5(string(input))
	case "day_6":
		Day_6(string(input))
	case "day_7":
		Day_7(string(input))
	case "day_8":
		Day_8(string(input))
	}
}
