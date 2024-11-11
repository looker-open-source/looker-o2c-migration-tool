// Package utils provides utility functions for the looker usage tool.
package utils

import (
	"fmt"
	"os"

	"google3/base/go/log"
	"google3/third_party/looker_o2c_cli/looker_usage/lookerusage"
	"google3/third_party/looker_sdk_codegen/go/rtl/rtl"
	"google3/third_party/looker_sdk_codegen/go/sdk/v4/v4"
)

// InitSession initializes the looker sdk session.
func InitSession(apiIDKey string, apiSecretKey string, lookerLocation string, ssl bool) *v4.LookerSDK {
	settings := rtl.ApiSettings{
		BaseUrl:      lookerLocation,
		VerifySsl:    ssl,
		Timeout:      2000,
		ClientId:     apiIDKey,
		ClientSecret: apiSecretKey,
		ApiVersion:   "4.0",
	}

	authSession := rtl.NewAuthSession(settings)
	return v4.NewLookerSDK(authSession)
}

// WriteLookerUsageToCSV writes the looker usage to a csv file.
func WriteLookerUsageToCSV(path string, lu *lookerusage.LookerUsage) {
	if path == "" {
		path = "looker-usage.csv"
	}
	log.Infof("CREATING CSV FILE")
	outputCSV, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Errorf("Failed to create csv file: %v", err)
		return
	}
	defer outputCSV.Close()

	log.Infof("WRITING TO CSV FILE")
	fmt.Fprintln(outputCSV, lu.String())
}