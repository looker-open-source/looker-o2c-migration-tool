/**
 * @license
 * Copyright 2024 Google LLC.
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE file at https://opensource.org/licenses/MIT
 */

// The o2c_cli is for getting information about a Looker O2C customers instance.
package main

import (
	"github.com/looker-open-source/looker_o2c_migration_evaluation/lookerusage"
	"github.com/looker-open-source/looker_o2c_migration_evaluation/utils"

	"flag"
)

var (
	// Used for flags.
	apiIDKey       = flag.String("api-id", "", "Looker generated API ID")
	apiSecretKey   = flag.String("api-secret", "", "Looker generated API Secret")
	lookerLocation = flag.String("addr", "", "Looker instance address, with 'http(s)://' prefix")
	ssl            = flag.Bool("ssl", true, "Should the connection use SSL")
	outputCSVPath  = flag.String("output-csv", "", "Path to output csv file")
)

func main() {
	// Initialize Google libraries.
	flag.Set("alsologtostderr", "true")
	flag.Parse()
	lookersdk := utils.InitSession(*apiIDKey, *apiSecretKey, *lookerLocation, *ssl)
	lu := lookerusage.LookerUsage{}

	lu.ComputeUsage(lookersdk)
	utils.WriteLookerUsageToCSV(*outputCSVPath, &lu)
}
