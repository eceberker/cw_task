package gcpclient

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
	"github.com/eceberker/cw_task/helpers"
	"google.golang.org/api/option"
)

// NewBigQueryClient creates and returns BigQuery Client
func NewBigQueryClient() (*bigquery.Client, error) {

	ctx := context.Background()

	env := helpers.GetEnv()
	cred := env["GOOGLE_CREDETENTIALS_FILE_NAME"]
	projectID := env["PROJECT_ID"]

	client, err := bigquery.NewClient(ctx, projectID, option.WithCredentialsFile(cred))

	if err != nil {
		log.Fatalf("An error occured : %v", err)
		return nil, err
	}

	return client, nil
}
