package main

type Employee struct {
	Name    string
	Address string
}

type Details struct {
	ID     string  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
