package main

import (
    "context"

    "net/http"
    "log"
)

const DBNAME        = "helloworld"
const DBADDR        = "localhost:9000"
const TABLENAME     = "information"

func main() {
    // Setup the server
    http.HandleFunc("/page/", handlerFile)
    http.HandleFunc("/db/", handlerDB)
    log.Fatal(http.ListenAndServe(":8080", nil))

    // Setup the db
    ctx := context.Background()
    err, conn := ConnectDB(DBADDR, DBNAME, ctx)
    if err != nil { panic(err) }

    // Create if not exists
    CreateTable(ctx, TABLENAME, conn)
}
