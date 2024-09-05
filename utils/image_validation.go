package utils

import (
	"errors"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

var (
	// MIME type to Extensions
	mimeType2Extension = map[string]string{
		"image/jpg":  "jpg",
		"image/jpeg": "jpg",
		"image/png":  "png",
		"image/gif":  "gif",
	}
)

func ReadAndValidatePhoto(ctx iris.Context) ([]byte, string, string, error) {

	// Retrieve all form values, which include files
	files, _, err := ctx.FormFiles("media")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "some error happen"})
		return nil, "", "", err
	}
	if len(files) == 0 {
		// Handle no file error
		golog.Infof("No media provided")
		return nil, "", "", nil
	} else if len(files) > 1 {
		// Handle multiple files error
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Multiple media files are not allowed"})
		return nil, "", "", errors.New("multiple media files provided")
	}

	// Retrieve the file
	photoFile, fileHeader, err := ctx.FormFile("media")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Error retrieving media file"})
		return nil, "", "", err
	}
	defer photoFile.Close()

	photoBytes, err := io.ReadAll(photoFile)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Error reading media file"})
		return nil, "", "", err
	}

	// Check the photo's MIME type
	mimeType, err := detectAndValidateContentType(photoBytes)
	if err != nil {
		ctx.StatusCode(iris.StatusUnsupportedMediaType)
		ctx.JSON(iris.Map{"error": "Unsupported file type"})
		return nil, "", "", errors.New("unsupported file type")
	}

	return photoBytes, mimeType, fileHeader.Filename, nil
}

func DownlodAndValidatePhoto(ctx iris.Context, photoUrl string) ([]byte, string, string, error) {

	response, err := http.Get(photoUrl)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "failed to download photo"})
		return nil, "", "", err
	}
	defer response.Body.Close()

	// Get photo into memory
	photoBytes, err := io.ReadAll(response.Body)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Error reading media file"})
		return nil, "", "", err
	}

	mimeType, err := detectAndValidateContentType(photoBytes)
	if err != nil {
		ctx.StatusCode(iris.StatusUnsupportedMediaType)
		ctx.JSON(iris.Map{"error": "Unsupported file type"})
		return nil, "", "", errors.New("unsupported file type")
	}

	fileName, err := buildFileName(photoUrl, mimeType)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Unable to extract filename from url"})
		return nil, "", "", err
	}

	return photoBytes, mimeType, fileName, nil
}

func detectAndValidateContentType(photoBytes []byte) (string, error) {
	// Check the photo's MIME type
	mimeType := http.DetectContentType(photoBytes)
	if _, ok := mimeType2Extension[mimeType]; ok {
		return mimeType, nil
	} else {
		return "", errors.New("unsupported file type")
	}
}

func buildFileName(fullUrlFile string, mimeType string) (string, error) {
	fileUrl, err := url.Parse(fullUrlFile)
	if err != nil {
		return "", err
	}

	path := fileUrl.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	// if file without an extension then append one based on context type
	if filepath.Ext(fileName) == "" {
		ext, _ := mimeType2Extension[mimeType]
		return fileName + "." + ext, nil
	}

	return fileName, nil

}
