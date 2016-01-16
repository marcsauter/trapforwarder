package sensu

import (
	"encoding/json"
	"io"

	"github.com/marcsauter/trapforwarder/trap"
)

// Sensu
type Sensu struct {
	Name   string     `json:"name"`
	Output *trap.Trap `json:"output"`
	Status int        `json:"status"`
}

// NewSensu creates a new Sensu type
func NewSensu(t *trap.Trap) *Sensu {
	s := &Sensu{
		Name:   "application name", // get from trap
		Output: t,
		Status: 0, // get from trap
	}
	return s
}

func (s *Sensu) Send(w io.Writer) error {
	d, err := json.Marshal(s)
	if err != nil {
		return err
	}
	_, err = w.Write(d)
	return err
}
