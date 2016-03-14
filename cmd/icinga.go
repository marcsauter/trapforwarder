// Copyright ©2016 Marc Sauter <marc.sauter@bluewin.ch>
//
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package cmd

import (
	"fmt"
	"os"

	"github.com/marcsauter/trapforwarder/icinga"
	"github.com/marcsauter/trapforwarder/logger"
	"github.com/marcsauter/trapforwarder/trap"
	"github.com/spf13/cobra"
)

var pipe string

// icingaCmd respresents the icinga command
var icingaCmd = &cobra.Command{
	Use:   "icinga",
	Short: "convert and send trap to icinga",
	Run: func(cmd *cobra.Command, args []string) {
		// read the trap
		t := trap.NewTrap()
		// icinga
		i := icinga.NewIcinga(t)
		// connect to icinga client
		if stat, err := os.Stat(pipe); os.IsNotExist(err) || stat.Mode()&os.ModeNamedPipe == 0 {
			logger.Fatal(fmt.Sprintf("%s does not exist or is not a named pipe\n", pipe))
		}
		f, err := os.OpenFile(pipe, os.O_RDWR, 0666)
		if err != nil {
			logger.Fatal(err.Error())
		}
		defer f.Close()
		// send event
		if err := i.Send(f); err != nil {
			w, _ := logger.New(logger.Severity, logger.Prefix)
			t.Send(w)
			logger.Fatal(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(icingaCmd)
	icingaCmd.Flags().StringVar(&pipe, "pipe", "/var/run/icinga2/cmd/icinga2.cmd", "pipe where icinga is listening")
}
