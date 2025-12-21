CREATE TABLE IF NOT EXISTS problems (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    template TEXT NOT NULL,
    difficulty VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS tests (
    id SERIAL PRIMARY KEY,
    problem_id INTEGER NOT NULL,
    input TEXT NOT NULL,
    output TEXT NOT NULL
);

-- create reference between problem and test
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_constraint
        WHERE conname = 'tests_problem_id_fkey'
    ) THEN
        ALTER TABLE tests 
        ADD CONSTRAINT tests_problem_id_fkey 
        FOREIGN KEY (problem_id) REFERENCES problems(id) 
        ON DELETE CASCADE;
    END IF;
END $$;