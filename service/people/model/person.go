package model

// Person is a entity that represents a person
type Person struct {
	// ID is the identifier of the Entity, the ID is shared for all sub domains
	ID string `json:"id,omitempty"`
	// Name is the name of the person
	Name string `json:"name,omitempty"`
}
