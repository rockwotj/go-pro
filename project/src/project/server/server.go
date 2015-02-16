package server

import (
	"html/template"
	"net/http"
	"fmt"
	"os"
	"io"
	"strings"
	"strconv"
	"path/filepath"
	"archive/zip"
)

type page struct {
	Count int
	Body  string
}

var photosUploaded = 0

var templates = template.Must(template.ParseFiles("./static/index.html", "./static/submit.html"))

func makeHandler(templateName string, body string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, templateName, page{Count: photosUploaded, Body: body})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

 func uploadHandler(w http.ResponseWriter, r *http.Request) {
 		photosUploaded++
        // the FormFile function takes in the POST input id file
        file, header, err := r.FormFile("file")
        if err != nil {
            fmt.Fprintln(w, err)
            return
        }
        defer file.Close()
        uploadedPath := "./tmp/uploadedfile" + strconv.Itoa(photosUploaded)
        if strings.HasSuffix(header.Filename, ".zip") {
        	uploadedPath += ".zip"
        } else if strings.HasSuffix(header.Filename, ".jpg") || strings.HasSuffix(header.Filename, ".jpeg") {
        	uploadedPath += ".jpg"
        } else {
        	fmt.Fprintf(w, "Wrong type of file uploaded")
            return
        }
        out, err := os.Create(uploadedPath)
        if err != nil {
            fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
            return
        }
        defer out.Close()
        // write the content from POST to the file
        _, err = io.Copy(out, file)
        if err != nil {
            fmt.Fprintln(w, err)
        }
        fmt.Fprintf(w, "File uploaded successfully : ")
        fmt.Fprintf(w, header.Filename)
 }

func Start() {
	http.HandleFunc("/receive", uploadHandler)
	http.HandleFunc("/submit", makeHandler("submit.html", "unknown! The project is not finished yet, check back later"))
	http.HandleFunc("/", makeHandler("index.html", "Please submit your photo, and we'll tell you if it is a sunset or not!"))
	http.ListenAndServe(":8080", nil)
}

func unzip(zipfile string) {
    reader, err := zip.OpenReader(zipfile)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer reader.Close()
    for _, f := range reader.Reader.File {
        zipped, err := f.Open()
        if err != nil {
            fmt.Println(err)
            return
        }
        defer zipped.Close()
        // get the individual file name and extract the current directory
        path := filepath.Join("./tmp/", f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())
            fmt.Println("Creating directory", path)
        } else {
            writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, f.Mode())
            if err != nil {
                fmt.Println(err)
                return
        	}
            defer writer.Close()
            if _, err = io.Copy(writer, zipped); err != nil {
                fmt.Println(err)
                return
            }
        fmt.Println("Decompressing : ", path)
        }
    }
}