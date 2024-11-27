CREATE TABLE TBLInscription (
    user_id INT NOT NULL,
    id_class INT NOT NULL,
    CONSTRAINT fk_teacher FOREIGN KEY (user_id) REFERENCES TBLUser(user_id),
    CONSTRAINT fk_class FOREIGN KEY (id_class) REFERENCES TBLClass(id_class)
);
