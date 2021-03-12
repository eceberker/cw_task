package services

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	GCPClients "github.com/eceberker/cw_task/gcp_clients"
	"github.com/eceberker/cw_task/models"
	"google.golang.org/api/iterator"
)

// GetDailyAverageDurations queries average session duration of daily active users according to day from GCP BigQuery
func GetDailyAverageDurations() models.DailyAverageDurationsResponse {
	ctx := context.Background()

	q := `SELECT
			FORMAT_TIMESTAMP("%b-%d-%Y", start_time ) AS date,
			ROUND(AVG(TIMESTAMP_DIFF(end_time, start_time, MINUTE)), 2) AS avg_duration
 		  FROM (
			SELECT
				session_id,
				MIN(TIMESTAMP_MILLIS(event_time)) AS start_time,
				MAX(TIMESTAMP_MILLIS(event_time)) AS end_time,
			FROM ` + "`cwtask-307017.cw_task_dataset.user_logst_test`" + `
			GROUP BY
	  		  session_id )
  		  GROUP BY
		    date
		  ORDER BY 
  			date;`

	var response models.DailyAverageDurationsResponse
	response.MessageText = "Daily average session durations in minutes"
	job, err := query(q)

	if err != nil {
		response.MessageText = fmt.Sprintf("An error occured %v", err)
		response.Status = 500
		return response
	}

	it, err := job.Read(ctx)

	for {
		var row models.BigQueryDailyAverageDurations
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			response.Status = 500
			response.MessageText = fmt.Sprintf("An error occured: %v.", err)
			return response
		}
		response.Rows = append(response.Rows, row)
	}
	return response
}

// GetUsersPerDay queries count of daily active users according to day from GCP BigQuery
func GetUsersPerDay() models.DailyUsersResponseModel {

	ctx := context.Background()

	q := `SELECT
			DISTINCT FORMAT_TIMESTAMP("%b-%d-%Y", TIMESTAMP_MILLIS(event_time)) AS date,
			COUNT(DISTINCT(user_id)) OVER(PARTITION BY FORMAT_TIMESTAMP("%b-%d-%Y", TIMESTAMP_MILLIS(event_time))) AS unique_users
  		  FROM ` + "`cwtask-307017.cw_task_dataset.user_logst_test`" + `
  		  ORDER BY
			date;`

	// OLD query for user list for each day
	// q := `SELECT
	// 		DISTINCT FORMAT_TIMESTAMP("%b-%d-%Y", TIMESTAMP_MILLIS(event_time)) as date,
	// 		user_id
	// 	  FROM ` + "`cwtask-307017.cw_task_dataset.user_logst_test`" + `
	// 	  GROUP BY
	// 	    date, user_id;`

	var response models.DailyUsersResponseModel
	response.MessageText = "Number of unique users per day"
	response.Status = 200

	job, err := query(q)

	if err != nil {
		response.Status = 500
		response.MessageText = fmt.Sprintf("An error occured: %v.", err)
		return response
	}

	it, err := job.Read(ctx)

	for {
		var row models.BigQueryDailyUsersRow
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			response.Status = 500
			response.MessageText = fmt.Sprintf("An error occured: %v.", err)
			return response
		}
		response.Rows = append(response.Rows, row)
	}
	return response
}

// GetTotalUsers queries users ID and their last online date from GCP BigQuery
func GetTotalUsers() models.TotalUsersResponse {
	ctx := context.Background()

	q := `SELECT
			user_id,
			last_online AS last_online_date
  		  FROM (
			SELECT
	  		  user_id,
	  		  MAX(FORMAT_TIMESTAMP("%c", TIMESTAMP_MILLIS(event_time))) AS last_online,
			FROM ` + "`cwtask-307017.cw_task_dataset.user_logst_test`" + `
			GROUP BY
	  		user_id)
  		  GROUP BY
			user_id,
			last_online
		  ORDER BY
		     last_online;`

	var response models.TotalUsersResponse
	response.MessageText = "Total users of app"
	response.Status = 200

	job, err := query(q)

	if err != nil {
		response.Status = 500
		response.MessageText = fmt.Sprintf("An error occured: %v.", err)
		return response
	}

	it, err := job.Read(ctx)

	for {
		var row models.BigQueryTotalUsers
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			response.Status = 500
			response.MessageText = fmt.Sprintf("An error occured: %v.", err)
			return response
		}
		response.Rows = append(response.Rows, row)
	}
	return response

}

func query(query string) (*bigquery.Job, error) {

	ctx := context.Background()

	client, er := GCPClients.NewBigQueryClient()

	if er != nil {
		return nil, er
	}

	defer client.Close()

	q := client.Query(query)

	q.Location = "US"

	job, err := q.Run(ctx)
	if err != nil {
		return nil, err
	}

	status, err := job.Wait(ctx)

	if err != nil {
		return nil, err
	}

	if err := status.Err(); err != nil {
		return nil, err
	}

	return job, nil
}
