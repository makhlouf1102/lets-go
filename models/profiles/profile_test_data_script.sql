-- Insert test users (without profiles initially)
INSERT INTO user (user_id, username, email, password) VALUES (1, 'testuser1', 'test1@example.com', 'password1');
INSERT INTO user (user_id, username, email, password) VALUES (2, 'testuser2', 'test2@example.com', 'password2');
INSERT INTO user (user_id, username, email, password) VALUES (3, 'testuser3', 'test3@example.com', 'password3');

-- Insert some test profiles
INSERT INTO profile (user_id, first_name, last_name, date_of_birth, bio) 
VALUES (1, 'Test', 'User', '1990-01-01', 'Test biography');