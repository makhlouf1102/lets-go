-- Insérer des utilisateurs dans la table `user`
INSERT INTO user (user_id, username, email, password, first_name, last_name)
VALUES 
    ('testuser1', 'testuser1', 'test1@example.com', 'password1', 'Test1', 'User1'),
    ('testuser2', 'testuser2', 'test2@example.com', 'password2', 'Test2', 'User2'),
    ('testuser3', 'testuser3', 'test3@example.com', 'password3', 'Test3', 'User3');

-- Insérer des rôles dans la table `role`
INSERT INTO role (role_id, name)
VALUES 
    ('testrole1', 'Admin'),
    ('testrole2', 'Editor'),
    ('testrole3', 'Viewer');

-- Insérer des relations utilisateur-rôle dans la table `user_role`
INSERT INTO user_role (user_role_id, user_id, role_id)
VALUES 
    ('testuserrole1', 'testuser1', 'testrole1'),
    ('testuserrole2', 'testuser1', 'testrole2'),
    ('testuserrole3', 'testuser2', 'testrole3');
