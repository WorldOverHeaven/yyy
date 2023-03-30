package model

type Participant struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Wish      string       `json:"wish"`
	Recipient *Participant `json:"recipient" pg:"rel:has-one"`
}
