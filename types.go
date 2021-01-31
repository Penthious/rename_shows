package main

type ShowsStruct []struct {
	ShowStruct ShowStruct `json:"show"`
}
type ShowStruct struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
}
