package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

const (
	folderId = "1tGoORaO3mA7rFuU95idqUv7QBjDyi6rF"
	fileName = "credentials.json"
)

func main() {

	// create context
	ctx := context.Background()
	// create drive service with credentials, use default if provided
	driveService, err := drive.NewService(ctx, option.WithCredentialsFile(fileName), option.WithScopes(drive.DriveScope))
	if err != nil {
		panic(err)
	}

	// list files
	files, err := driveService.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		panic(err)
	}

	// create folder
	folderCreated, err := driveService.Files.Create(&drive.File{
		Name:     "test",
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{folderId},
	}).Do()
	if err != nil {
		panic(err)
	}
	folderCreated.Shared = true

	// read local file
	localFile, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer localFile.Close()

	// upload file
	if _, err := localFile.Seek(0, 0); err != nil {
		panic(err)
	}
	fileToCreate, err := driveService.Files.Create(&drive.File{
		Name:    "test.txt",
		Parents: []string{folderCreated.Id},
	}).Do()
	if err != nil {
		panic(err)
	}
	fileToCreate.Shared = true

	file, err := driveService.Files.Update(fileToCreate.Id, &drive.File{
		MimeType: "text/plain",
	}).Media(bufio.NewReader(localFile)).Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(files)
	fmt.Println(folderCreated)
	fmt.Println(file)
}
