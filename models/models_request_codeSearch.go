package models

type BodyRequestCodeNextPage struct {
	Size        int     `json:"size"`
	Query       Query   `json:"query"`
	Sort        []Sort  `json:"sort"`
	SearchAfter []int64 `json:"search_after"`
}

type BodyRequestCode struct {
	Size  int    `json:"size"`
	Query Query  `json:"query"`
	Sort  []Sort `json:"sort"`
}

type Query struct {
	Bool Bool `json:"bool"`
}

type Bool struct {
	Must []Must `json:"must"`
}

type Must struct {
	Match Match `json:"match"`
}

type Match struct {
	ClasseCodigo        int `json:"classe.codigo,omitempty"`
	OrgaoJulgadorCodigo int `json:"orgaoJulgador.codigo,omitempty"`
}

type Sort struct {
	Timestamp `json:"@timestamp"`
}

type Timestamp struct {
	Order string `json:"order"`
}
