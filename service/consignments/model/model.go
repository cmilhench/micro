package model

type Container struct {
	ID         string `json:"id,omitempty"`
	CustomerID string `json:"customer_id,omitempty"`
	Origin     string `json:"origin,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}

type Consignment struct {
	ID          string       `json:"id,omitempty"`
	Description string       `json:"description,omitempty"`
	Weight      int32        `json:"weight,omitempty"`
	Containers  []*Container `json:"containers,omitempty"`
	VesselID    string       `json:"vessel_id,omitempty"`
}

// ---

type Repository interface {
	Create(*Consignment) (*Consignment, error)
	GetAll() []*Consignment
}
