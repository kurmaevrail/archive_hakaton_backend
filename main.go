package main

import (
    "context"

    "net/http"
    "log"
)

const DBNAME        = "helloworld"
const DBADDR        = "localhost:9000"
const TABLENAME     = "information"

const BUFSIZE       = 100

func main() {
    // Setup the db
    ctx := context.Background()

    err, conn := connectDB(DBADDR, DBNAME, ctx)
    if err != nil { panic(err) }

    // Create if not exists
    createTable(ctx, TABLENAME, conn)

    // Setup the dbhandler
    handlerDB := dbhandler { ctx: ctx, db: DBNAME, table: TABLENAME, conn: conn, buffer: make([]row, 0, BUFSIZE) }

    // Setup the server
    http.HandleFunc("/page/", handlerFile)
    http.Handle("/db/", handlerDB)
    log.Fatal(http.ListenAndServe(":8080", nil))


}
