package models

import "time"

type ResponseBody struct {
	Took     int   `json:"took"`
	TimedOut bool  `json:"timed_out"`
	Shards   Shard `json:"_shards"`
	Hit      Hit   `json:"hits"`
}

type Shard struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type Hit struct {
	Total    Total   `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hit2  `json:"hits"`
}

type Total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type Hit2 struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	Id     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source Source  `json:"_source"`
}

type Source struct {
	LawsuitNumber    string     `json:"numeroProcesso"`
	Class            Class      `json:"classe"`
	System           System     `json:"sistema"`
	Format           Format     `json:"formato"`
	Court            string     `json:"tribunal"`
	DateLastUpdate   time.Time  `json:"dataHoraUltimaAtualizacao"`
	Degree           string     `json:"grau"`
	Timestamp        time.Time  `json:"@timestamp"`
	DistributionDate time.Time  `json:"dataAjuizamento"`
	Movements        []Movement `json:"movimentos"`
	Id               string     `json:"id"`
	SecrecyLevel     int        `json:"nivelSigilo"`
	CourtInstance    Court      `json:"orgaoJulgador"`
	Subjects         []Subject  `json:"assuntos"`
}

type Class struct {
	Code int    `json:"codigo"`
	Name string `json:"nome"`
}

type System struct {
	Code int    `json:"codigo"`
	Name string `json:"nome"`
}

type Format struct {
	Code int    `json:"codigo"`
	Name string `json:"nome"`
}

type Movement struct {
	Complement []Complement `json:"complementosTabelados,omitempty"`
	Code       int          `json:"codigo"`
	Name       string       `json:"nome"`
	DateTime   time.Time    `json:"dataHora"`
}

type Complement struct {
	Code        int    `json:"codigo"`
	Value       int    `json:"valor"`
	Name        string `json:"nome"`
	Description string `json:"descricao"`
}

type Court struct {
	CountyCodeIBGE int    `json:"codigoMunicipioIBGE"`
	Code           int    `json:"codigo"`
	Name           string `json:"nome"`
}

type Subject struct {
	Code int    `json:"codigo"`
	Name string `json:"nome"`
}
