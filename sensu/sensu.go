package sensu

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/marcsauter/trapforwarder/trap"
)

// Sensu
type Sensu struct {
	Name   string `json:"name"`
	Output string `json:"output"`
	Status int    `json:"status"`
}

// NewSensu creates a new Sensu type
func NewSensu(t *trap.Trap) *Sensu {
	var msg string
	s := &Sensu{}
	s.Name = t.Vars["SNMPv2-SMI::enterprises.5471.2.20.15"]
	switch t.Vars["SNMPv2-SMI::enterprises.5471.2.20.10"] {
	case "0":
		switch strings.ToUpper(t.Vars["SNMPv2-SMI::enterprises.5471.2.20.20"]) {
		case "NORMAL":
			s.Status = 0
		case "WARNING":
			s.Status = 1
		case "CRITICAL":
			s.Status = 2
		default:
			s.Status = 3
		}
		msg = fmt.Sprintf("node=%s", t.Vars["SNMPv2-SMI::enterprises.5471.2.20.15"])
		msg = fmt.Sprintf("%s;severity=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.20"])
		msg = fmt.Sprintf("%s;application=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.25"])
		msg = fmt.Sprintf("%s;object=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.30"])
		msg = fmt.Sprintf("%s;msg_grp=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.35"])
		msg = fmt.Sprintf("%s;msg_text=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.40"])
		msg = fmt.Sprintf("%s;service_id=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.45"])
		s.Output = msg
	case "1":
		switch strings.ToUpper(t.Vars["SNMPv2-SMI::enterprises.5471.2.20.20"]) {
		case "OK":
			s.Status = 0
		case "WARNING":
			s.Status = 1
		case "CRITICAL":
			s.Status = 2
		default:
			s.Status = 3
		}
		msg = fmt.Sprintf("host=%s", t.Vars["SNMPv2-SMI::enterprises.5471.2.20.15"])
		msg = fmt.Sprintf("%s;severity=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.20"])
		msg = fmt.Sprintf("%s;object_class=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.25"])
		msg = fmt.Sprintf("%s;object=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.30"])
		msg = fmt.Sprintf("%s;parameter=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.35"])
		msg = fmt.Sprintf("%s;parameter_value=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.40"])
		msg = fmt.Sprintf("%s;message=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.45"])
		msg = fmt.Sprintf("%s;tc=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.50"])
		msg = fmt.Sprintf("%s;ac=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.55"])
		msg = fmt.Sprintf("%s;it_service=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.60"])
		msg = fmt.Sprintf("%s;alarm_id=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.65"])
		msg = fmt.Sprintf("%s;dedup_key=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.70"])
		msg = fmt.Sprintf("%s;security_event=%s", msg, t.Vars["SNMPv2-SMI::enterprises.5471.2.20.75"])
		s.Output = msg
	default:
		s.Output = fmt.Sprintf("UNKNOWN VERSION: %s", t.Vars["SNMPv2-SMI::enterprises.5471.2.20.10"])
		s.Status = 1
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
