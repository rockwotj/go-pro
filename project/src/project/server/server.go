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
    "project/imageProcessor"
    "project/svm"
)

type page struct {
	Count int
	Body  string
}


func Start() {
	// load data to train SVM
    sunsets := imageProcessor.ProcessDirectory("../../../TrainSunset/*.jpg")
    nonsunsets := imageProcessor.ProcessDirectory("../../../TrainNonsunsets/*.jpg")
    labelsSunset := make([]float64,len(sunsets))
	// create labels
    for i :=0; i < len(sunsets); i++{
	labelsSunset[i] = 1
    }
    labelsNonsunset := make([]float64,len(nonsunsets))
    for i :=0; i < len(nonsunsets); i++{
	labelsNonsunset[i] = -1
    }
	// append everything
    labels := append(labelsSunset, labelsNonsunset...)
    data2 := append(sunsets, nonsunsets...)
	// normalize and train
    data := svm.NormalizeAll(data2)
    svm.Train(data,labels)


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

        fmt.Fprintf(w, "<table border='1' cellspacing='5' cellpadding='3'><tr><th>File</th><th>Sunset?</th></tr>")
        data := imageProcessor.Process(uploadedPath)
        if data == nil {
            fmt.Fprintf(w, "Invalid image format!")
            return
        }
        result := svm.Predict(svm.Normalize(data))
        isSunset := "Unkown"
        if result == 1 {
            isSunset = "Yes"
        } else {
            isSunset = "No"
        }
        fmt.Fprintf(w, "<tr><td>" + rawURL + "</td><td>" + isSunset + "</td></tr>")
        fmt.Fprintf(w, "</table>")
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
        file.Close()
        out.Close()
        if strings.HasSuffix(header.Filename, ".zip") {
        	unzip(uploadedPath, w)
            os.RemoveAll(uploadedPath)
        } else {
            fmt.Fprintf(w, "<table border='1' cellspacing='5' cellpadding='3'><tr><th>File</th><th>Sunset?</th></tr>")
            data := imageProcessor.Process(uploadedPath)
            if data == nil {
                fmt.Fprintf(w, "Invalid image format!")
                return
            }
            result := svm.Predict(data)
            isSunset := "Unkown"
            if result == 1 {
                isSunset = "Yes"
            } else {
                isSunset = "No"
            }
            fmt.Fprintf(w, "<tr><td>" + header.Filename + "</td><td>" + isSunset + "</td></tr>")
            fmt.Fprintf(w, "</table>")
        }
 }

func unzip(zipfile string, w http.ResponseWriter) {
    reader, err := zip.OpenReader(zipfile)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer reader.Close()
    jipsUploaded++
    newDir := "./tmp/images" + strconv.Itoa(jipsUploaded) +"/"
    os.MkdirAll(newDir, 0777)
    fmt.Fprintf(w, "<table border='1' cellspacing='5' cellpadding='3'><tr><th>File</th><th>Sunset?</th></tr>")
    var wg sync.WaitGroup
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
            unzippedFile := filepath.Join(newDir, f.Name)
            wg.Add(1)
            go func() {
                defer wg.Done()
                copyFile(unzippedFile, path)
                data := imageProcessor.Process(uploadedPath)
                if data == nil {
                    fmt.Fprintf(w, "<tr><td>" + f.Name + "</td><td>Unknown</td></tr>")
                    return
                }
                result := svm.Predict(data)
                isSunset := "Unkown"
                if result == 1 {
                    isSunset = "Yes"
                } else {
                    isSunset = "No"
                }
                fmt.Fprintf(w, "<tr><td>" + f.Name + "</td><td>" + isSunset + "</td></tr>")
                os.Remove(path)
            }
        }
    }
    wg.Wait()
    fmt.Fprintf(w, "</table>")
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
