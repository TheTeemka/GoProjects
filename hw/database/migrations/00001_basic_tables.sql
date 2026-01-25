-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS faculties (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    faculty_id INT NOT NULL REFERENCES faculties(id)
);

CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    gender VARCHAR(10) NOT NULL,
    group_id INT NOT NULL REFERENCES groups(id)
);

CREATE TABLE IF NOT EXISTS schedules (
    id SERIAL PRIMARY KEY,
    subject TEXT NOT NULL,
    day_of_week TEXT NOT NULL,
    time TEXT NOT NULL,
    group_id INT NOT NULL REFERENCES groups(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS schedules;
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS faculties;
-- +goose StatementEnd
