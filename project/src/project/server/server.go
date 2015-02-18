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


func Start() {
    http.HandleFunc("/receiveUrl", downloadHandler)
    http.HandleFunc("/receive", uploadHandler)
    http.HandleFunc("/submit", makeHandler("submit.html", "unknown! The project is not finished yet, check back later"))
    http.HandleFunc("/", makeHandler("index.html", "Please submit your photo, and we'll tell you if it is a sunset or not!"))
    http.ListenAndServe(":8080", nil)
}

var photosUploaded = 0
var jipsUploaded = 0
var templates = template.Must(template.ParseFiles("./static/index.html", "./static/submit.html"))

func makeHandler(templateName string, body string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, templateName, page{Count: photosUploaded, Body: body})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
        photosUploaded++
        //Call to ParseForm makes form fields available.
        r.ParseForm()
        rawURL := r.PostFormValue("url")
        fmt.Println("Received URL: ", rawURL);
        uploadedPath := "./tmp/uploadedfile" + strconv.Itoa(photosUploaded) + ".jpg"
        file, err := os.Create(uploadedPath)
        if err != nil {
            fmt.Println(err)
            panic(err)
        }
        defer file.Close()
        check := http.Client{
            CheckRedirect: func(r *http.Request, via []*http.Request) error {
                r.URL.Opaque = r.URL.Path
                return nil
            },
        }
        resp, err := check.Get(rawURL) // add a filter to check redirect
        if err != nil {
            fmt.Println(err)
            panic(err)
        }
        defer resp.Body.Close()
        fmt.Println(resp.Status)
        _, err = io.Copy(file, resp.Body)
        if err != nil {
            panic(err)
        }
        fmt.Fprintf(w, "File uploaded successfully : ")
        fmt.Fprintf(w, rawURL)
 }

 func uploadHandler(w http.ResponseWriter, r *http.Request) {
 		photosUploaded++
        // the FormFile function takes in the POST input id file
        file, header, err := r.FormFile("file")
        if err != nil {
            fmt.Fprintln(w, err)
            return
        }
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
        // write the content from POST to the file
        _, err = io.Copy(out, file)
        if err != nil {
            fmt.Fprintln(w, err)
        }
        fmt.Fprintf(w, "File uploaded successfully : ")
        fmt.Fprintf(w, header.Filename)
        file.Close()
        out.Close()
        if strings.HasSuffix(header.Filename, ".zip") {
        	unzip(uploadedPath)
            os.RemoveAll(uploadedPath)
        }
 }

func unzip(zipfile string) {
    reader, err := zip.OpenReader(zipfile)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer reader.Close()
    jipsUploaded++
    newDir := "./tmp/images" + strconv.Itoa(jipsUploaded) +"/"
    os.MkdirAll(newDir, 0777)
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
            
            if _, err = io.Copy(writer, zipped); err != nil {
                fmt.Println(err)
                return
            }
        writer.Close()
        fmt.Println("Decompressing : ", path)
        photosUploaded++
        copyFile(filepath.Join(newDir, f.Name), path)
        os.Remove(path)
        }
    }
}

func copyFile(dst, src string) error {
    in, err := os.Open(src)
    if err != nil { return err }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil { return err }
    defer out.Close()
    _, err = io.Copy(out, in)
    cerr := out.Close()
    if err != nil { return err }
    return cerr
}
