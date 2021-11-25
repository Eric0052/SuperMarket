package main

type Order struct {
	Id  string `json:"id"`
	Add int    `ison:add`
}
