package main

import (
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const imageIDLength = 10

type Image struct {
	ID          string
	UserID      string
	Name        string
	Location    string
	Size        int64
	CreatedAt   time.Time
	Description string
}

func NewImage(user *User) *Image {
	return &Image{
		ID:        GenerateID("img", imageIDLength),
		UserID:    user.ID,
		CreatedAt: time.Now(),
	}
}

func (image *Image) StaticRoute() string {
	return "/im/" + image.Location
}

func (image *Image) ShowRoute() string {
	return "/image/" + image.ID
}

// A map of accepted mime types and their file extension
var mimeExtensions = map[string]string{
	"image/png":  ".png",
	"image/jpeg": ".jpg",
	"image/gif":  ".gif",
}

func (image *Image) CreateFromURL(imageURL string) error {
	// Get the response from the URL
	response, err := http.Get(imageURL)
	if err != nil {
		return err
	}

	// Make sure we got a response
	if response.StatusCode != http.StatusOK {
		return errImageURLInvalid
	}

	defer response.Body.Close()

	// Ascertain the type of file we downloaded
	mimeType, _, err := mime.ParseMediaType(response.Header.Get("Content-Type"))
	if err != nil {
		return errInvalidImageType
	}

	// Get an extension for the file
	ext, valid := mimeExtensions[mimeType]
	if !valid {
		return errInvalidImageType
	}

	// Get a name from the URL
	image.Name = filepath.Base(imageURL)
	image.Location = image.ID + ext

	// Open a file at target location
	savedFile, err := os.Create("./data/images/" + image.Location)
	if err != nil {
		return err
	}
	defer savedFile.Close()

	// Copy the entire response to the output flie
	size, err := io.Copy(savedFile, response.Body)
	if err != nil {
		return err
	}
	// The returned value from io.Copy is the number of bytes copied
	image.Size = size

	// Save our image to the store
	return globalImageStore.Save(image)
}

func (image *Image) CreateFromFile(file multipart.File, headers *multipart.FileHeader) error {

	// Move our file to an appropriate place, with an appropriate name
	image.Name = headers.Filename
	image.Location = image.ID + filepath.Ext(image.Name)

	// Open a file at target location
	savedFile, err := os.Create("./data/images/" + image.Location)
	if err != nil {
		return err
	}

	defer savedFile.Close()

	// Copy the uploaded file to the target location
	size, err := io.Copy(savedFile, file)
	if err != nil {
		return err
	}
	image.Size = size

	// Save the image to the database
	return globalImageStore.Save(image)
}
