CREATE TABLE user (
    user_id CHAR(36) PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT
);

CREATE TABLE role (
    name TEXT PRIMARY KEY
);

CREATE TABLE user_role (
    user_role_id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    role_id CHAR(36) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (role_id) REFERENCES role(role_id)
);

CREATE TABLE language (
    language_id CHAR(36) PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE problem (
    problem_id CHAR(36) PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    difficulty TEXT NOT NULL
);

CREATE TABLE solved_problem (
    solved_problem_id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    problem_id CHAR(36) NOT NULL,
    language_id CHAR(36) NOT NULL,
    solution TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(user_id),
    FOREIGN KEY (problem_id) REFERENCES problem(problem_id),
    FOREIGN KEY (language_id) REFERENCES language(language_id)
);
