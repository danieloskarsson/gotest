package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Name    string
	Age     int
	Student bool
}

type blah string

const temp = "TEST"

func (v blah) hey() {
	fmt.Println(v)
}

func main() {
	fmt.Println("Hello, 世界")

	x()
	x, y := y(1, 2)
	fmt.Println(x, y)

	var t blah = "test"
	t.hey()
}

func y(x int, y int) (int, int) {
	return x, y
}

func x() {

	str, _ := ioutil.ReadFile("file.json")
	var config Config
	err := json.Unmarshal(str, &config)
	fmt.Println(str)
	fmt.Println(temp)

	if err == nil {
		fmt.Println(config.Name)
	}

}
