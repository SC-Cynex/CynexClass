CREATE TABLE IF NOT EXISTS TBLClass (
    id_class SERIAL PRIMARY KEY,
    name_class VARCHAR(100) NOT NULL,
    semester VARCHAR(20) NOT NULL,
    teacher_id INT NOT NULL,
    subject_id INT NOT NULL,
    day_subject VARCHAR(20) NOT NULL,
    course_id INT NOT NULL,
    period VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_teacher FOREIGN KEY (teacher_id) REFERENCES TBLUser(user_id),
    CONSTRAINT fk_subject FOREIGN KEY (subject_id) REFERENCES TBLSubject(id_subject),
    CONSTRAINT fk_course FOREIGN KEY (course_id) REFERENCES TBLCourse(id_course)
);
