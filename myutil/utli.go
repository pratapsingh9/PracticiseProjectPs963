package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Employee struct {
	Id int
	name string
	salary float64
	Upper bool
}

type Joke struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func SendPostRequest(url string, payload []byte) {
    Request, err := http.Post(url, "application/json", bytes.NewReader(payload))
    if err != nil {
        panic(err)
    }
    defer Request.Body.Close()
    fmt.Println("Request status Code Sent Succesfully ", Request.StatusCode)
}

func PostJsonRequest() {
	url := "http://localhost:3000/api"
	requestbody := strings.NewReader(`
	{
		"course":"Pratap Singh Course",
		"name":"Pratap",
		"age":25,
		"increments":10
	}`)
	fmt.Println(requestbody)
	// responese,err = http.Post(url)
	response, err := http.Post(url, "application/json", requestbody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Response successfully done by golang", response.Status)
}

func main() {
	SendPostRequest()
	url := "https://official-joke-api.appspot.com/random_joke"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	var joke Joke
	if err := json.Unmarshal(body, &joke); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	} else {
		fmt.Println("Json perfectly working")
	}

	fmt.Println(joke.Setup)
	fmt.Println(joke.Punchline)
	fmt.Println(joke.Type)

	PostJsonRequest()
}
