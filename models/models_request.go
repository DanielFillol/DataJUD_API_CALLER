package models

type BodyRequest struct {
	Query Query `json:"query,omitempty"`
}

type Query struct {
	Match Match `json:"match,omitempty"`
}

type Match struct {
	CNJNumber string `json:"numeroProcesso,omitempty"`
}
