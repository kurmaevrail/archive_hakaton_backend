package main

import (
    "os"
    "fmt"
    "unsafe"
    "strconv"
    "strings"

    "net/http"
    "context"

//  "github.com/ClickHouse/clickhouse-go/v2"
    "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
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

func serveRows(rows []row, w http.ResponseWriter, r * http.Request) {
    ptr     := unsafe.Pointer(&rows[0])
    rowSize := unsafe.Sizeof(rows[0])

    bSize   := uintptr(len(rows)) * rowSize
    bBuffer := unsafe.Slice((*byte)(ptr), bSize)
    w.Header().Set("Content-Type", "text/plain")
    w.Write(bBuffer)
}

type dbhandler struct {
    ctx     context.Context
    db      string
    table   string
    conn    driver.Conn
    buffer  []row
}

func (this dbhandler) ServeHTTP(w http.ResponseWriter, r * http.Request) {
    path := r.URL.Path;
    n := strings.Index(path, ":") // Location of the ':'

    if n < 0 {
        fmt.Printf("Index failed (Couldn't find %s in %s)\n", ":", path)
        w.WriteHeader(404) // Write fail status if failed to read the number of entries
        return;
    }
    firstStr    := path[len("/db/"):n]
    secondStr   := path[n+1:]

    // How many entries to read
    offset, err := strconv.Atoi(firstStr)

    if err != nil {
        fmt.Println("Atoi(failed) [1]")
        w.WriteHeader(404) // Write fail status if failed to read the number of entries
        return;
    }

    count, err := strconv.Atoi(secondStr)

    if err != nil {
        fmt.Println("Atoi(failed) [2]")
        w.WriteHeader(404) // Write fail status if failed to read the number of entries
        return;
    }

    rows := this.buffer[:0:count]
    readTable(this.ctx, this.db, this.table, this.conn, offset, count, rows)
    serveRows(rows[:count], w, r)
}
