// Chronodium - Keeping Time in Series
//
// Copyright 2016-2017 Dolf Schimmel
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
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	BuildTag  string
	BuildTime string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of chronodium",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"chronodium - Keeping time in series - %s\n\n"+
				"%s\nCopyright (c) 2016-2017, Dolf Schimmel\n"+
				"License: Apache License, Version 2.0\n\n"+
				"Time of Build: %s\n\n",
			BuildTag,
			"https://github.com/Freeaqingme/chronodium",
			BuildTime)
	},
}
