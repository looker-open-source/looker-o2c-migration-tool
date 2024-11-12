/**
 * @license
 * Copyright 2024 Google LLC.
 *
 * Use of this source code is governed by an MIT-style license that can be
 * found in the LICENSE file at https://opensource.org/licenses/MIT
 */

 // Package lookerusage provides the implementation for the looker-usage command.
package lookerusage

import (
	"encoding/json"
	"fmt"
	"strings"

	"google3/base/go/log"
	"google3/third_party/looker_o2c_cli/looker_usage/types"
	"google3/third_party/looker_sdk_codegen/go/sdk/v4/v4"
)

// LookerUsage is the struct for the looker usage in the output csv file.
type LookerUsage struct {
	Projects          []v4.Project       `csv:"PROJECT INFO"`
	Users             []v4.User          `csv:"USER INFO"`
	Dashboards        []v4.DashboardBase `csv:"DASHBOARD INFO"`
	Schedules         []v4.ScheduledPlan `csv:"SCHEDULE INFO"`
	Queries           types.QueryStats   `csv:"QUERY INFO"`
	LegacyFeatures    []v4.LegacyFeature `csv:"LEGACY FEATURES"`
	Settings          v4.Setting         `csv:"SETTINGS"`
	ConnectionDetails []v4.DBConnection  `csv:"CONNECTION DETAILS"`
}

// ComputeUsage gets the usage for the looker instance.
func (lu *LookerUsage) ComputeUsage(lookersdk *v4.LookerSDK) {
	log.Info("PARSING PROJECTS")
	lu.fetchProjectDetails(lookersdk)
	log.Info("PARSING USERS")
	lu.fetchUserDetails(lookersdk)
	log.Info("PARSING DASHBOARDS")
	lu.fetchDashboardDetails(lookersdk)
	log.Info("PARSING SCHEDULES")
	lu.fetchScheduleDetails(lookersdk)
	log.Info("PARSING QUERY HISTORY")
	lu.fetchQueryDetails(lookersdk)
	log.Info("PARSING LEGACY FEATURES")
	lu.fetchLegacyFeatures(lookersdk)
	log.Info("PARSING SETTINGS")
	lu.fetchSettings(lookersdk)
	log.Info("PARSING CONNECTION DETAILS")
	lu.fetchConnectionDetails(lookersdk)

}

func (lu *LookerUsage) String() string {
	buf := new(strings.Builder)
	buf.WriteString(composeString("PROJECT INFO", types.KeyValue{Key: "Total Projects", Value: len(lu.Projects)}))
	buf.WriteString(composeString("USER INFO", types.KeyValue{Key: "Total Users", Value: len(lu.Users)}))
	buf.WriteString(composeString("DASHBOARD INFO", types.KeyValue{Key: "Total Dashboards", Value: len(lu.Dashboards)}))
	buf.WriteString(composeString("SCHEDULE INFO", types.KeyValue{Key: "Total Schedules", Value: len(lu.Schedules)}))
	buf.WriteString(composeString("QUERY INFO", lu.Queries))
	buf.WriteString(composeString("LEGACY FEATURES", lu.LegacyFeatures))
	buf.WriteString(composeString("FEATURE FLAGS", *lu.Settings.InstanceConfig.FeatureFlags))
	buf.WriteString(composeString("LICENSE FEATURES", *lu.Settings.InstanceConfig.LicenseFeatures))
	buf.WriteString(composeString("OTHER SETTINGS", otherSettings(lu)))
	buf.WriteString(composeString("CONNECTION DETAILS", lu.ConnectionDetails))
	return buf.String()
}

func (lu *LookerUsage) fetchProjectDetails(lookersdk *v4.LookerSDK) {
	var err error
	lu.Projects, err = lookersdk.AllProjects("", nil)
	if err != nil {
		log.Errorf("Failed to get project details: %v", err)
	}
}

func (lu *LookerUsage) fetchUserDetails(lookersdk *v4.LookerSDK) {
	var err error
	lu.Users, err = lookersdk.AllUsers(v4.RequestAllUsers{}, nil)
	if err != nil {
		log.Errorf("Failed to get user details: %v", err)
	}
}

func (lu *LookerUsage) fetchDashboardDetails(lookersdk *v4.LookerSDK) {
	var err error
	lu.Dashboards, err = lookersdk.AllDashboards("", nil)
	if err != nil {
		log.Errorf("Failed to get dashboard details: %v", err)
	}
}

func (lu *LookerUsage) fetchScheduleDetails(lookersdk *v4.LookerSDK) {
	var err error
	lu.Schedules, err = lookersdk.AllScheduledPlans(v4.RequestAllScheduledPlans{}, nil)
	if err != nil {
		log.Errorf("Failed to get schedule details: %v", err)
	}
}

// getQueryDetails returns the query details for the last month.
// This opens the gate for the history table to be queried for this tool.
// I would expect we reiterate on this to make it more useful.
// Starting with this to get something out there.
func (lu *LookerUsage) fetchQueryDetails(lookersdk *v4.LookerSDK) {
	req := v4.RequestRunInlineQuery{}
	filters := map[string]any{}
	filters["history.created_date"] = "1 month"
	req.Body = v4.WriteQuery{
		Model: "system__activity",
		View:  "history",
		Fields: &[]string{
			"history.created_hour",
			"history.query_run_count",
		},
		Filters: &filters,
	}
	req.ResultFormat = "json"
	result, err := lookersdk.RunInlineQuery(req, nil)
	if err != nil {
		log.Errorf("Failed to get query details: %v", err)
	}
	queryCount, queryMax, _ := queryCounts(result)

	lu.Queries = types.QueryStats{
		QueryCount: queryCount,
		QueryMax:   queryMax,
	}
}

func queryCounts(result string) (int, int, error) {
	var res []types.QueryResult
	err := json.Unmarshal([]byte(result), &res)
	if err != nil {
		return 0, 0, err
	}

	queryCount := 0
	queryMax := 0

	for _, r := range res {
		queryCount += r.HistoryQueryRunCount
		if r.HistoryQueryRunCount > queryMax {
			queryMax = r.HistoryQueryRunCount
		}
	}

	return queryCount, queryMax, nil
}

func (lu *LookerUsage) fetchLegacyFeatures(lookersdk *v4.LookerSDK) {
	var err error
	lu.LegacyFeatures, err = lookersdk.AllLegacyFeatures(nil)
	if err != nil {
		log.Errorf("Failed to get legacy features: %v", err)
	}
}

func (lu *LookerUsage) fetchSettings(lookersdk *v4.LookerSDK) {
	var err error
	lu.Settings, err = lookersdk.GetSetting("", nil)
	if err != nil {
		log.Errorf("Failed to get settings: %v", err)
	}
}

func (lu *LookerUsage) fetchConnectionDetails(lookersdk *v4.LookerSDK) {
	var err error
	lu.ConnectionDetails, err = lookersdk.AllConnections("", nil)
	if err != nil {
		log.Errorf("Failed to get connection details: %v", err)
	}
}

// TODO: Make it generic to handle any data type.
func composeString(header string, data any) string {
	buf := new(strings.Builder)
	buf.WriteString(header + "\n\n")

	switch v := data.(type) {
	case map[string]bool:
		for key, value := range v {
			if _, err := buf.WriteString(fmt.Sprintf("%s,%t\n", key, value)); err != nil {
				log.Errorf("Failed to create csv string: %v", err)
			}
		}
	case types.KeyValue:
		if _, err := buf.WriteString(fmt.Sprintf("%s,%d\n", v.Key, v.Value)); err != nil {
			log.Errorf("Failed to create csv string: %v", err)
		}
	case types.QueryStats:
		if _, err := buf.WriteString("Last Month Total Queries," + fmt.Sprintf("%d", v.QueryCount) + "\n"); err != nil {
			log.Errorf("Failed to create csv string: %v", err)
		}
		if _, err := buf.WriteString("Last Month Max Queries in an Hour," + fmt.Sprintf("%d", v.QueryMax) + "\n"); err != nil {
			log.Errorf("Failed to create csv string: %v", err)
		}
	case []v4.LegacyFeature:
		for _, r := range v {
			if _, err := buf.WriteString(fmt.Sprintf("%s\n", *r.Name)); err != nil {
				log.Errorf("Failed to create csv string: %v", err)
			}
		}
	case []v4.DBConnection:
		for _, r := range v {
			if _, err := buf.WriteString(fmt.Sprintf("%s,%s\n", *r.Name, *r.Dialect.Name)); err != nil {
			}
		}
	default:
		log.Errorf("Unsupported data type, %T", v)
	}
	buf.WriteString("\n") // Add an empty line after writing data

	return buf.String()
}

func otherSettings(lu *LookerUsage) map[string]bool {
	settings := map[string]bool{
		"Marketplace Enabled":         *lu.Settings.MarketplaceEnabled,
		"Embed Enabled":               *lu.Settings.EmbedEnabled,
		"Extension Framework Enabled": *lu.Settings.ExtensionFrameworkEnabled,
	}
	return settings
}