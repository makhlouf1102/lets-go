-- Insert test users
INSERT INTO user (user_id, username, email, password, first_name, last_name)
VALUES ('testuser1', 'testuser1', 'test1@example.com', 'password1', 'Test1', 'User1');

-- Insert test sessions
INSERT INTO session (session_id, user_id, refreshToken)
VALUES ('testsession1', 'testuser1', 'refreshtoken1'),
       ('testsession2', 'testuser1', 'refreshtoken2');
