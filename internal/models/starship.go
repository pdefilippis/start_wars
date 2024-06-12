package models

type Starship struct{
	Name string `json:"name"`
	Model string `json:"model"`
}

type StarshipResult struct{
	Next *string `json:"next"`
	Results []Starship `json:"results"`
}