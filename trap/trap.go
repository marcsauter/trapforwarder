package trap

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
)

// Trap represents a SNMP trap as provided via traphandle
type Trap struct {
	Host string            `json:"host"`
	IP   string            `json:"ip"`
	Vars map[string]string `json:"vars"`
}

// NewTrap reads from stdin an creates a new trap
func NewTrap() *Trap {
	t := &Trap{Vars: make(map[string]string)}
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		switch count {
		case 0:
			t.Host = line
		case 1:
			t.IP = line
		default:
			d := strings.SplitN(line, " ", 2)
			if len(d) == 2 {
				t.Vars[d[0]] = strings.Trim(d[1], "\"")
			}
		}
		count++
	}
	return t
}

func (t *Trap) Send(w io.Writer) error {
	d, err := json.Marshal(t)
	if err != nil {
		return err
	}
	_, err = w.Write(d)
	return err
}
