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
	"net"
	"time"

	"github.com/marcsauter/trapforwarder/logger"
	"github.com/marcsauter/trapforwarder/sensu"
	"github.com/marcsauter/trapforwarder/trap"
	"github.com/spf13/cobra"
)

const (
	connectTimeout = 3 // seconds
	writeTimeout   = 3 // seconds
)

var host, port string

// sensuCmd respresents the sensu command
var sensuCmd = &cobra.Command{
	Use:   "sensu",
	Short: "converted and send trap to sensu",
	Run: func(cmd *cobra.Command, args []string) {
		// read the trap
		t := trap.NewTrap()
		// sensu
		s := sensu.NewSensu(t)
		// connect to sensu client
		c, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), time.Second*connectTimeout)
		if err != nil {
			w, _ := logger.New(logger.Severity, logger.Prefix)
			t.Send(w)
			logger.Fatal(err.Error())
		}
		c.SetWriteDeadline(time.Now().Add(time.Second * writeTimeout))
		// send event
		if err := s.Send(c); err != nil {
			w, _ := logger.New(logger.Severity, logger.Prefix)
			t.Send(w)
			logger.Fatal(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(sensuCmd)
	sensuCmd.Flags().StringVar(&host, "host", "localhost", "host where sensu client is running")
	sensuCmd.Flags().StringVar(&port, "port", "3030", "port of the sensu client")
}
