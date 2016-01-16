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
	"github.com/marcsauter/trapforwarder/logger"
	"github.com/marcsauter/trapforwarder/trap"
	"github.com/spf13/cobra"
)

var prefix, severity string

// logCmd respresents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "write the trap syslog",
	Run: func(cmd *cobra.Command, args []string) {
		// read the trap
		t := trap.NewTrap()
		// new logger
		w, _ := logger.New(logger.Priority(severity), prefix)
		// send raw event
		t.Send(w)
	},
}

func init() {
	RootCmd.AddCommand(logCmd)
	logCmd.Flags().StringVar(&prefix, "prefix", logger.Prefix, "log prefix")
	// TODO(marc): check severity
	logCmd.Flags().StringVarP(&severity, "severity", "s", "info", "log severity")
}
