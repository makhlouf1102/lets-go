CREATE TABLE problems (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    template TEXT NOT NULL,
    difficulty VARCHAR(50) NOT NULL
);

CREATE TABLE tests (
    id SERIAL PRIMARY KEY,
    problem_id INTEGER NOT NULL,
    input TEXT NOT NULL,
    output TEXT NOT NULL
);

-- create reference between problem and test
ALTER TABLE tests ADD FOREIGN KEY (problem_id) REFERENCES problems(id);
