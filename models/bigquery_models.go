package models

// BigQueryUserRow corresponds a full user rows
type BigQueryUserRow struct {
	Type      string `bigquery:"type" json:"type"`
	AppID     string `bigquery:"add_id" json:"app_id"`
	SessionID string `bigquery:"session_id" json:"session_id"`
	EventName string `bigquery:"event_name" json:"event_name"`
	EventTime int    `bigquery:"event_time" json:"event_time"`
	Page      string `bigquery:"page" json:"page"`
	Region    string `bigquery:"region" json:"region"`
	City      string `bigquery:"city" json:"city"`
	UserID    string `bigquery:"user_id" json:"user_id"`
}

// BigQueryDailyUsersRow corresponds count of daily active users according to day query rows
type BigQueryDailyUsersRow struct {
	Date                string `bigquery:"date" json:"date"`
	NumberOfUniqueUsers int    `bigquery:"unique_users" json:"unique_users"`
}

// BigQueryDailyAverageDurations corresponds average session duration of daily active users according to day query rows
type BigQueryDailyAverageDurations struct {
	Date      string  `bigquery:"date" json:"date"`
	Durations float32 `bigquery:"avg_duration" json:"avg_durations"`
}

// BigQueryTotalUsers Bigquery users ID and their last online date rows
type BigQueryTotalUsers struct {
	UsersID string `bigquery:"user_id" json:"user_id"`
	Date    string `bigquery:"last_online_date" json:"last_online_date"`
}
