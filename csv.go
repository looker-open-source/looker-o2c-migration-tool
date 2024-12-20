/**
 * @license
 * Copyright Google LLC All Rights Reserved.
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE file at https://opensource.org/licenses/MIT
 */

// Package csv contains the logic to write data to a csv file.
package csv

import (
	"os"

	"google3/base/(internal)"
)

// WriteDataToCSV writes the looker usage data to a csv file.
func WriteDataToCSV(path string, usage string) {
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
}
