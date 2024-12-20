/**
 * @license
 * Copyright Google LLC All Rights Reserved.
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE file at https://opensource.org/licenses/MIT
 */

// The o2c_cli is for getting information about a Looker O2C customers instance.
package main

import (
	"fmt"
	"os/exec"

	"google3/base/(internal)"
	"google3/base/(internal)"
	"google3/third_party/looker_o2c_cli/csv"
	"google3/third_party/looker_o2c_cli/lookerusage"
	"google3/third_party/looker_o2c_cli/session"
)

var (
	// Used for flags.
	command        = flag.String("command", "", "Command to execute")
	apiIDKey       = flag.String("api-id", "", "Looker generated API ID")
	apiSecretKey   = flag.String("api-secret", "", "Looker generated API Secret")
	lookerLocation = flag.String("addr", "", "Looker instance address, with 'http(s)://' prefix")
	ssl            = flag.Bool("ssl", true, "Should the connection use SSL")
	outputCSVPath  = flag.String("output-csv", "", "Path to output csv file")
)

// Defined for commands.
const computeUsageCommand = "compute-usage"
const fileSystemPerformanceCommand = "file-system-performance"

func main() {
	// Initialize Google libraries.
	flag.Set("alsologtostderr", "true")
	google.Init()

	switch *command {
	case computeUsageCommand:
		lookersdk := session.InitSession(*apiIDKey, *apiSecretKey, *lookerLocation, *ssl)
		lu := lookerusage.LookerUsage{}

		lu.ComputeUsage(lookersdk)
		csv.WriteDataToCSV(*outputCSVPath, lu.String())

	case fileSystemPerformanceCommand:
		cmd := exec.Command("./looker-file-system-performance")
		output, _ := cmd.Output()

		csv.WriteDataToCSV(*outputCSVPath, string(output))

	default:
		fmt.Println("Invalid command")
	}
}
