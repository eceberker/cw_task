package models

// PublishResponseModel corresponds PubSub client's response as JSON result
type PublishResponseModel struct {
	Status      int    `json:"status,omitempty"`
	MessageID   string `json:"message_id,omitempty"`
	MessageText string `json:"message_text,omitempty"`
}

// UserResponseModel corresponds User query row as JSON result
type UserResponseModel struct {
	Status      string            `json:"status,omitempty"`
	MessageText string            `json:"message_text,omitempty"`
	Rows        []BigQueryUserRow `json:"users,omitempty"`
}

// DailyUsersResponseModel corresponds daily user rows as JSON result
type DailyUsersResponseModel struct {
	Status      int                     `json:"status,omitempty"`
	MessageText string                  `json:"message_text,omitempty"`
	Rows        []BigQueryDailyUsersRow `json:"number_of_unique_users,omitempty"`
}

// DailyAverageDurationsResponse corresponds daily average duration query rows as JSON result
type DailyAverageDurationsResponse struct {
	Status      int                             `json:"status,omitempty"`
	MessageText string                          `json:"message_text,omitempty"`
	Rows        []BigQueryDailyAverageDurations `json:"daily_average_durations,omitempty"`
}

// TotalUsersResponse for total user query rows as JSON result
type TotalUsersResponse struct {
	Status      int                  `json:"status,omitempty"`
	MessageText string               `json:"message_text,omitempty"`
	Rows        []BigQueryTotalUsers `json:"users,omitempty"`
}
