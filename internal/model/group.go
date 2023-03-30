package model

type Group struct {
	ID           int64          `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Participants []*Participant `json:"participants" pg:"rel:has-many"`
}
