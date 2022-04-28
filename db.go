package main

import (
       "context"
       "fmt"
       "time"

       "github.com/ClickHouse/clickhouse-go/v2"
       "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func connectDB(addr string, dbName string, ctx context.Context) (error, driver.Conn) {
    ctx         = context.Background()
    conn, err   := clickhouse.Open(&clickhouse.Options{
        Addr: []string{addr},
        Auth: clickhouse.Auth{
            Database: dbName,
            Username: "default",
            Password: "",
        },
        Debug:           true,
        DialTimeout:     time.Second,
        MaxOpenConns:    10,
        MaxIdleConns:    5,
        ConnMaxLifetime: time.Hour,
    })
    return err, conn;
}

// Create the MergeTree table
func createTable(ctx context.Context, table string, conn driver.Conn) (error) {
    query := fmt.Sprintf(`
        CREATE TABLE %s (
            Time                    Float,
            Humidity                Float,
            "Room temperature"      Float,
            "Workspace temperature" Float,
            pH                      Float,
            Mass                    Float,
            Waste                   Float,
            CO2                     Float
        ) ENGINE = MergeTree()
        PRIMARY KEY (Time)`, table);
    return conn.Exec(ctx, query);
}

func insertBatched(ctx context.Context, conn driver.Conn, table string, rows []row) (error) {
    batch, err := conn.PrepareBatch(ctx, fmt.Sprintf("INSERT INTO %s", table));
    if err != nil { return err }

    for i := 0; i < len(rows); i++ {
        err = batch.AppendStruct(&rows[i]);
        if err != nil { return err }
    }

    return batch.Send();
}

// Read count rows from the table
func readTable(ctx context.Context,
             db string,
             table string,
             conn driver.Conn,
             offset int,
             count int,
             dest []row) (error) {
    query := fmt.Sprintf("SELECT * FROM %s.%s LIMIT %d OFFSET %d", db, table, count, offset)
    err := conn.Select(ctx, &dest, query)
    return err
}
