INSERT INTO roles (role_name) VALUES ('Admin');
INSERT INTO roles (role_name) VALUES ('Moderator');
INSERT INTO roles (role_name) VALUES ('User');
INSERT INTO roles (role_name) VALUES ('Guest');

-- {
--     "Username": "testUser",
--     "Email": "testUser@example.com",
--     "Password": "testPassword",
--     "ProfilePicture": "https://example.com/profile.jpg",
--     "Biography": "This is a test user",
--     "RoleID": 1
-- }

INSERT INTO users (username, email, password_hash, profile_picture, biography, role_id) 
VALUES ('testUser', 'testUser@testUser@example.com', '$2a$10$sT4z5AHcw5CqATcCBIklqeSKNnW1XVnaQQ9KBCEdL0Q5DGbJoDnU2', 'https://example.com/profile.jpg', 'This is a test user', 1);
