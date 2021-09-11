// fork of https://github.com/zupzup/golang-http-file-upload-download

package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const maxUploadSize = 4 * 1024 * 1024 // 4 MB
const uploadDir = "./public/uploads"

func main() {
	fs := http.FileServer(http.Dir(uploadDir))
	http.Handle("/files/", http.StripPrefix("/files/", fs))
	http.HandleFunc("/upload", uploadFileHandler)

	log.Println("Server started on localhost:8080, routes:")
	log.Println("  /upload")
	log.Println("  /files/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("public/upload.gtpl")
		t.Execute(w, nil)
		return
	}

	// parse the form
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		log.Printf("parse multipart form failure: %v\n", err)
		renderError(w, "INVALID_FORM", http.StatusBadRequest)
		return
	}

	// validate the upload
	//  open upload
	file, fileHeader, err := r.FormFile("f")
	if err != nil {
		renderError(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileSize := fileHeader.Size
	//  check upload size
	if fileSize > maxUploadSize {
		renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
		return
	}
	//  read upload bytes
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		renderInternalError(w, "READ_FILE_FAILED")
		return
	}
	//  check file type, DetectContentType only needs the first 512 bytes
	fileType := http.DetectContentType(fileBytes)
	switch fileType {
	case "image/jpeg":
	case "image/png":
	case "image/gif":
	case "application/pdf":
		break
	default:
		renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
		return
	}
	fileExtensions, err := mime.ExtensionsByType(fileType)
	if err != nil {
		renderInternalError(w, "FILE_EXTENSION")
		return
	}
	fileName := randToken(12) + fileExtensions[len(fileExtensions)-1]

	// write file to upload folder
	filePath := filepath.Join(uploadDir, fileName)
	err = writeFile(filePath, fileBytes)
	if err != nil {
		log.Printf("upload failed: %v\n", err)
		renderInternalError(w, "UPLOAD_FAILED")
		return
	}
	log.Printf("upload %s %s %d bytes\n", fileType, filePath, fileSize)
	fmt.Fprintln(w, "SUCCESS")
}

func writeFile(name string, data []byte) error {
	newFile, err := os.Create(name)
	if err != nil {
		return err
	}
	defer newFile.Close()
	_, err = newFile.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, message)
}

func renderInternalError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, message)
}

func randToken(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
