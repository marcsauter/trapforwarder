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
	Host string     `json:"host"`
	IP   string     `json:"ip"`
	Vars []Variable `json:"vars"`
}

// Variable represents a OID/Value pair
type Variable struct {
	OID   string `json:"oid"`
	Value string `json:"value"`
}

// NewTrap reads from stdin an creates a new trap
func NewTrap() *Trap {
	t := &Trap{}
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
				v := Variable{OID: d[0], Value: strings.Trim(d[1], "\"")}
				t.Vars = append(t.Vars, v)
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
