package task

import (
	"context"
	"fmt"
	"app/model"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	t := new(model.Task)

	if err := c.Bind(t); err != nil {
		return err
	}

	sql := `
        INSERT INTO tasks (
            id,
            name,
            details,
            done,
            createdat
        ) VALUES (
            uuid_generate_v4(),
            $1,
            $2,
            FALSE,
            CURRENT_TIMESTAMP
        )
    `
	if _, err = conn.Exec(context.Background(), sql, t.Name, t.Details); err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
	}

	return c.String(http.StatusCreated, "Success!!")
}
