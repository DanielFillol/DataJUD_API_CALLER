package models

import "time"

type ResponseBodyNextPage struct {
	Took     int         `json:"took"`
	TimedOut bool        `json:"timed_out"`
	Shards   Shard       `json:"_shards"`
	Hit      HitNextPage `json:"hits"`
}

type HitNextPage struct {
	Total    Total          `json:"total"`
	MaxScore float64        `json:"max_score"`
	Hits     []Hit2NextPage `json:"hits"`
}

type Hit2NextPage struct {
	Index  string         `json:"_index"`
	Type   string         `json:"_type"`
	Id     string         `json:"_id"`
	Score  float64        `json:"_score"`
	Source SourceNextPage `json:"_source"`
	Sort   []int64        `json:"sort"`
}

type SourceNextPage struct {
	Class            Class      `json:"classe"`
	LawsuitNumber    string     `json:"numeroProcesso"`
	System           System     `json:"sistema"`
	Format           Format     `json:"formato"`
	Court            string     `json:"tribunal"`
	DateLastUpdate   time.Time  `json:"dataHoraUltimaAtualizacao"`
	Degree           string     `json:"grau"`
	Timestamp        time.Time  `json:"@TimestampCode"`
	DistributionDate time.Time  `json:"dataAjuizamento"`
	Movements        []Movement `json:"movimentos"`
	Id               string     `json:"id"`
	SecrecyLevel     int        `json:"nivelSigilo"`
	CourtInstance    Court      `json:"orgaoJulgador"`
	Subjects         []Subject  `json:"assuntos"`
}
