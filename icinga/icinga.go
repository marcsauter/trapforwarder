package icinga

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/marcsauter/trapforwarder/trap"
)

// Icinga
type Icinga struct {
	Name    string
	Service string
	Status  int
	Output  string
}

// NewIcinga creates a new Sensu type
func NewIcinga(t *trap.Trap) *Icinga {
	var msg string
	s := &Icinga{}
	s.Name = t.Vars["SNMPv2-SMI::enterprises.5471.2.20.15"]
	switch t.Vars["SNMPv2-SMI::enterprises.5471.2.20.10"] {
	case "0":
		s.Service = t.Vars["SNMPv2-SMI::enterprises.5471.2.20.30"]
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
		s.Service = t.Vars["SNMPv2-SMI::enterprises.5471.2.20.30"]
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

func (s *Icinga) Send(w io.Writer) error {
	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("[%d] PROCESS_SERVICE_CHECK_RESULT;", time.Now().Unix()))
	buf.WriteString(fmt.Sprintf("%s;", s.Name))
	buf.WriteString(fmt.Sprintf("%s;", s.Service))
	buf.WriteString(fmt.Sprintf("%d;", s.Status))
	buf.WriteString(fmt.Sprintf("%s;", s.Output))
	buf.WriteString("\n")
	_, err := w.Write(buf.Bytes())
	return err
}
