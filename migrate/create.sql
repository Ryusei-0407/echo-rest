CREATE TABLE tasks (
    id uuid NOT NULL,
    PRIMARY KEY (id),
    name TEXT NOT NULL UNIQUE,
    details TEXT NOT NULL,
    done BOOLEAN NOT NULL,
    createdat TIMESTAMPTZ NOT NULL
);
