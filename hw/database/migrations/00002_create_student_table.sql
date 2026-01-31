-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY,
    group_id INT NOT NULL REFERENCES groups(id),
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    birthday DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);


CREATE TRIGGER update_students_updated_at
    BEFORE UPDATE ON students
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();


CREATE INDEX idx_students_group_id ON students(group_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_students_group_id;
DROP TABLE IF EXISTS students;
-- +goose StatementEnd
