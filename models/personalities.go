package models

type Personalitie struct {
	Id        int    `json:"id"`
	P_name    string `json:"name"`
	P_history string `json:"history"`
}

var Personalities []Personalitie
