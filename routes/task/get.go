package task

import (
	"context"
	"fmt"
    "net/http"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

    var greeting string
    err = conn.QueryRow(context.Background(), "SELECT 'Hello, World!!'").Scan(&greeting)
    if err != nil {
        fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
        os.Exit(1)
    }

    return c.String(http.StatusOK, greeting)
}
