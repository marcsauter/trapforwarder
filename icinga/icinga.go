package icinga

import "github.com/marcsauter/trapforwarder/trap"

// Icinga
type Icinga struct {
	Name   string    `json:"name"`
	Output trap.Trap `json:"output"`
	Status int       `json:"status"`
}

// NewIcinga creates a new Sensu type
func NewIcinga(t *trap.Trap) *Icinga {
	s := &Icinga{}
	return s
}
