package osv

import "encoding/json"

type Document struct {
	SchemaVersion    string          `json:"schema_version"`
	Id               string          `json:"id"`
	Modified         string          `json:"modified"`
	Published        string          `json:"published"`
	Withdrawn        string          `json:"withdrawn"`
	Aliases          []string        `json:"aliases"`
	Related          []string        `json:"related"`
	Summary          string          `json:"summary"`
	Details          string          `json:"details"`
	Severity         []Severity      `json:"severity"`
	Affected         []Affected      `json:"affected"`
	References       []Reference     `json:"references"`
	Credits          []Credit        `json:"credits"`
	DatabaseSpecific json.RawMessage `json:"database_specific"`
}

type Severity struct {
	Type  string `json:"type"`
	Score string `json:"score"`
}

type Affected struct {
	Package           Package         `json:"package"`
	Severity          []Severity      `json:"severity"`
	Ranges            []Range         `json:"ranges"`
	Versions          []string        `json:"versions"`
	EcosystemSpecific json.RawMessage `json:"ecosystem_specific"`
	DatabaseSpecific  json.RawMessage `json:"database_specific"`
}

type Package struct {
	Ecosystem string `json:"ecosystem"`
	Name      string `json:"name"`
	Purl      string `json:"purl"`
}

type Range struct {
	Type             string          `json:"type"`
	Repo             string          `json:"repo"`
	Events           []Event         `json:"events"`
	DatabaseSpecific json.RawMessage `json:"database_specific"`
}

type Event struct {
	Introduced   string `json:"introduced"`
	Fixed        string `json:"fixed"`
	LastAffected string `json:"last_affected"`
	Limit        string `json:"limit"`
}

type Reference struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type Credit struct {
	Name    string   `json:"name"`
	Contact []string `json:"contact"`
	Type    string   `json:"type"`
}
