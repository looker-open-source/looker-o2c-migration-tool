// Package types contains the types used by the other commands.
package types

import ()

// Project is the struct for the project in the output csv file.
type Project struct {
	ProjectName     string `csv:"Project Name"`
	ProjectID       string `csv:"Project ID"`
	ProjectBranches int    `csv:"Project Branches"`
	ProjectFiles    int    `csv:"Project Files"`
}

// QueryResult is the struct for the query stats in the Looker system activity model inline query results.
type QueryResult struct {
	HistoryCreatedHour   string `json:"history.created_hour"`
	HistoryQueryRunCount int    `json:"history.query_run_count"`
}

// QueryStats is the struct for the queries in the output csv file.
type QueryStats struct {
	QueryCount int `csv:"Query Count"`
	QueryMax   int `csv:"Max Queries in an Hour"`
}

// KeyValue struct is used to facilitate generic data handling for flushing to csv.
type KeyValue struct {
	Key   string `csv:"Key"`
	Value int    `csv:"Value"`
}