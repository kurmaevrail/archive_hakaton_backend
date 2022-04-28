package main

import (
    "os"
    "fmt"
    "unsafe"

    "net/http"
)

func handlerFile(w http.ResponseWriter, r * http.Request) {
    filename := r.URL.Path[len("/"):]
    handleFile(filename, w, r)
}

func handler404(w http.ResponseWriter, r * http.Request) {
    fmt.Fprintf(w, "404 page not found")
}

func handleFile(filename string, w http.ResponseWriter, r * http.Request) {
    // Load the content of a file
    content, err := os.ReadFile(filename);

    // Page not found
    if err != nil { handler404(w, r); return }

    // Write content of the file to the page
    fmt.Fprintf(w, string(content));
}

func handlerDB(w http.ResponseWriter, r * http.Request) {
    fBuffer := []float32 { 1.0, 2.0, 1.5 }
    ptr     := unsafe.Pointer(&fBuffer[0])

    f32size := 4 // Size of a float in bytes

    bBuffer := unsafe.Slice((*byte)(ptr), len(fBuffer) * f32size)
    w.Header().Set("Content-Type", "text/plain")
    w.Write(bBuffer)
}
