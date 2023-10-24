package models

type BodyRequestLawsuit struct {
	Query QueryLawsuit `json:"query,omitempty"`
}

type QueryLawsuit struct {
	Match MatchLawsuit `json:"match,omitempty"`
}

type MatchLawsuit struct {
	CNJNumber string `json:"numeroProcesso,omitempty"`
}
