package market

import "github.com/google/uuid"

type Product struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Price int       `json:"price,omitempty"`
}
