// The o2c_cli is for getting information about a Looker O2C customers instance.
package main

import (
	"google3/base/go/flag"
	"google3/base/go/google"
	"google3/third_party/looker_o2c_cli/looker_usage/lookerusage"
	"google3/third_party/looker_o2c_cli/looker_usage/utils"
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
	google.Init()
	lookersdk := utils.InitSession(*apiIDKey, *apiSecretKey, *lookerLocation, *ssl)
	lu := lookerusage.LookerUsage{}

	lu.ComputeUsage(lookersdk)
	utils.WriteLookerUsageToCSV(*outputCSVPath, &lu)
}