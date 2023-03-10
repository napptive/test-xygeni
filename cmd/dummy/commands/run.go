/**
 * Copyright 2020 Napptive
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package commands

import (
	"github.com/napptive/test-xygeni/internal/app/dummy"
	"github.com/spf13/cobra"
)

var runCmdLongHelp = "Launch the Dummy service"
var runCmdShortHelp = "Lauch the service"
var runCmdExample = `$ dummy run`
var runCmdUse = "run"

var runCmd = &cobra.Command{
	Use:     runCmdUse,
	Long:    runCmdLongHelp,
	Example: runCmdExample,
	Short:   runCmdShortHelp,
	Run: func(cmd *cobra.Command, args []string) {
		cfg.Debug = debugLevel
		s := dummy.NewService(cfg)
		s.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
