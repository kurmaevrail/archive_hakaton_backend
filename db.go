package main

import (
       "context"
       "fmt"
       "time"

       "github.com/ClickHouse/clickhouse-go/v2"
       "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)


func ConnectDB(addr string, dbName string, ctx context.Context) (error, driver.Conn) {
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
func CreateTable(ctx context.Context, table string, conn driver.Conn) (error) {
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

func InsertBatched(ctx context.Context, conn driver.Conn, table string, rows []row) (error) {
    batch, err := conn.PrepareBatch(ctx, fmt.Sprintf("INSERT INTO %s", table));
    if err != nil { return err }

    for i := 0; i < len(rows); i++ {
        err = batch.AppendStruct(&rows[i]);
        if err != nil { return err }
    }

    return batch.Send();
}

/*
func main () {
    ctx := context.Background();
    err, conn := ConnectDB("localhost:9000", DBNAME, ctx);
    if err != nil { panic (err) }

    CreateTable(ctx, TABLENAME, conn)

    row1 := row{ 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0 }

    err = InsertBatched(ctx, conn, TABLENAME, []row{row1});
    if err != nil { panic(err) }
    // err = Insert(ctx, conn, TABLENAME, "id, Name, Symbol", values);
}
*/
