package model

import "gorm.io/gorm"

type Vulnerability struct {
	gorm.Model
	SchemaVersion    string
	VulnerabilityID  string
	Modified         string
	Published        string
	Withdrawn        string
	Aliases          []Alias
	Related          []Related
	Summary          string
	Details          string
	Severity         []Severity `gorm:"many2many:vulnerability_severity;"`
	Affected         []Affected
	References       []Reference
	Credits          []Credit
	DatabaseSpecific []byte
}

type Alias struct {
	gorm.Model
	VulnerabilityID uint `gorm:"foreignKey:ID"`
	Value           string
}

type Related struct {
	gorm.Model
	VulnerabilityID uint `gorm:"foreignKey:ID"`
	Value           string
}

type Severity struct {
	gorm.Model
	Type  string
	Score string
}

type Affected struct {
	gorm.Model
	VulnerabilityID   uint `gorm:"foreignKey:ID"`
	Package           Package
	Severity          []Severity `gorm:"many2many:affected_severity;"`
	Ranges            []Range
	Versions          []Version
	EcosystemSpecific []byte
	DatabaseSpecific  []byte
}

type Package struct {
	gorm.Model
	AffectedID uint
	Ecosystem  string
	Name       string
	PURL       string
}

type Range struct {
	gorm.Model
	AffectedID       uint
	Type             string
	Repo             string
	Events           []Event
	DatabaseSpecific []byte
}

type Version struct {
	gorm.Model
	AffectedID uint
	Value      string
}

type Event struct {
	gorm.Model
	RangeID      uint
	Introduced   string
	Fixed        string
	LastAffected string
	Limit        string
}

type Reference struct {
	gorm.Model
	VulnerabilityID uint `gorm:"foreignKey:ID"`
	Type            string
	URL             string
}

type Credit struct {
	gorm.Model
	VulnerabilityID uint `gorm:"foreignKey:ID"`
	Name            string
	Contact         []Contact
	Type            string
}

type Contact struct {
	gorm.Model
	CreditID uint
	Value    string
}
