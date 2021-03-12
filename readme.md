# GCP Logging and Querying User Events Pipeline API Project

This repository holds code written in Go that provides endpoints to Google Cloud Platform Pub/Sub and BigQuery Services.

## Getting Started

These instructions will cover usage information and for the Docker Container

### Prerequisites

In order to run this container you'll need Docker installed on your machine.

#### Installataion

Steps to build a Docker image:

1. Clone this repo

```
git clone https://github.com/eceberker/cw_task.git
```

2. In the project directory 

```
docker-compose build
```
This will take a few minutes.

3. In the project directory

```
docker-compose up
```

4. Once everything has started up, you should be able to access API project via http://localhost:8080

##  API Endpoints

1. ```
POST http://localhost:8080/logs
```
Publishes the log in request body in GCP Pub/Sub service

##### Example request body:
 ```json
{
  "type": "event",
  "app_id": "com.codeway.test",
  "session_id": "ttttttttb",
  "event_name": "about",
  "event_time": 1615546818859,
  "page": "settings",
  "country": "US",
  "region": "New Jersey",
  "city": "Newark",
  "user_id": "9t0lrnuLQr"
}
```
If response success, it will return 
 ```json
{
    "status": 200,
    "message_id": "2085483466664289",
    "message_text": "Log is sent succesfully."
}
```
Otherwise, request will return Internal Server Error(500) and logs the cause.

2. 
```
GET http://localhost:8080/logs/durations/average
```
No payload or parameter is required.
Gets average session durations of daily active users according to day from GCP BigQuery Service.
If response success, it will return 
 ```json
{
    "message_text": "Daily average session durations in minutes",
    "daily_average_durations": [
        {
            "date": "Aug-23-2020",
            "avg_durations": 5.2
        },
        ...
     ]
}    
```
Otherwise, request will return Internal Server Error(500) and logs the cause.

3. 
```
GET http://localhost:8080/logs/daily/active
```
No payload or parameter is required.
Gets count of daily active unique users according to day from GCP BigQuery Service.
If response success, it will return 
 ```json
{
    "status": 200,
    "message_text": "Number of unique users per day",
    "number_of_unique_users": [
        {
            "date": "Aug-23-2020",
            "unique_users": 10
        },
        ...
    ]    
}
```
Otherwise, request will return Internal Server Error(500) and logs the cause.

4. 
```
GET http://localhost:8080/logs/users
```
No payload or parameter is required.
Gets total unique users ID and their last online date from GCP BigQuery Service.
If response success, it will return 
 ```json
{
    "status": 200,
    "message_text": "Total users of app",
    "users": [
        {
            "user_id": "9t0lrnuLQr",
            "last_online_date": "Thu Mar 11 13:03:52 2021"
        },
        ...
    ]
}        
```
Otherwise, request will return Internal Server Error(500) and logs the cause.


















# cw_task
# cw_task
# cw_task
