package model

type Count struct {
	FirstName string  `json:"frist_name"`
	LastName  string  `json:"last_name"`
	Password  string  `json:"password"`
	Money     string  `json:"money"`
	Identity  int64   `json:"identity"`
	Dart      []Goods `json:"dart"`
	Ticket    string  `json:"ticket"`
}

type Goods struct {
	Name      string `json:"name"`
	Price     string `json:"price"`
	Describe  string `json:"describe"`
	RepoCount int    `json:"repo_count"`
}
