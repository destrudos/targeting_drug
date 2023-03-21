package model

// Template represents a row from the `templates` table.
type Template struct {
	Name   string `db:"name" json:"name"`
	Config []byte `db:"config" json:"config"`
}
