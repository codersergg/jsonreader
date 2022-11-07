package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Property struct {
	GlobalId int `json:"global_id"`
}

type Response struct {
	Sum int
}

func main() {
	var (
		properties []Property
	)

	file, _ := os.Open("data-20190514T0100.json")
	dataFromFile, _ := io.ReadAll(file)
	err := json.Unmarshal(dataFromFile, &properties)
	if err != nil {
		panic("Ошибка JSON")
	}

	printResult(getSumId(properties))
}

func printResult(result int) {
	response := Response{
		Sum: result,
	}
	data, _ := json.MarshalIndent(response, "", "    ")
	fmt.Printf("%s", string(data))
}

func getSumId(properties []Property) int {
	var count int
	for _, property := range properties {
		count += property.GlobalId
	}
	return count
}
