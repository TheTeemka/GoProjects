CREATE TABLE IF NOT EXISTS attendance (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL REFERENCES students(id),
    subject_id INT NOT NULL,
    visit_date DATE NOT NULL,
    visited BOOLEAN NOT NULL DEFAULT FALSE
);