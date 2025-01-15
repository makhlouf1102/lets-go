INSERT INTO role (name) VALUES ('Admin');
INSERT INTO role (name) VALUES ('User');


INSERT INTO problem (problem_id, title, description, difficulty) VALUES
('123e4567-e89b-12d3-a456-426614174000', 'Two Sum', 'Given an array of integers, return indices of the two numbers such that they add up to a specific target.', 'Easy'),
('123e4567-e89b-12d3-a456-426614174001', 'Reverse Integer', 'Given a 32-bit signed integer, reverse digits of an integer.', 'Medium'),
('123e4567-e89b-12d3-a456-426614174002', 'Longest Substring Without Repeating Characters', 'Given a string, find the length of the longest substring without repeating characters.', 'Hard');

INSERT INTO user VALUES('6f2651cc-e02c-451d-8531-72f43c78f99d','mak','mak@mak.com','$2a$14$uN0QQIYsvmnteDQ3ORU4nuFazGUtjLuIVTM0R49AdHLoJuTc/PFGa','mak','mak');
INSERT INTO user_role VALUES('f1c38740-9bed-4869-819e-217231059ccf','6f2651cc-e02c-451d-8531-72f43c78f99d','User');
INSERT INTO user_role VALUES('123e4567-e89b-12d3-a456-426614174002','6f2651cc-e02c-451d-8531-72f43c78f99d','Admin');