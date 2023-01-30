# Google Drive API v3

This is a simple implementation of the Google Drive API v3 in GO

## Installation

```bash
go mod tidy
```

## Usage

```bash
go run main.go
```

## Instructions

1. Create a new project in Google Cloud Platform
2. Enable Google Drive API
3. Create a new Google Credentials (OAuth Client ID/API Key/Service Account)
4. For this project I used a Service Account for local implementation and for production I used default credentials by Google Cloud Platform (GCP) because we set the server in GCP
5. Download the credentials file and put it in the root of the project
6. Create a new folder in Google Drive and get the ID of the folder (Please add `google service account` as a member of the folder with the role of Editor)
7. Enjoy to try the API

# References

https://developers.google.com/drive/api/quickstart/go
https://pkg.go.dev/google.golang.org/api/drive/v3
