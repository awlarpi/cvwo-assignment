INSERT INTO roles (role_id, role_name) VALUES (1, 'Admin');
INSERT INTO roles (role_id, role_name) VALUES (2, 'Moderator');
INSERT INTO roles (role_id, role_name) VALUES (3, 'User');
INSERT INTO roles (role_id, role_name) VALUES (4, 'Guest');

INSERT INTO users (username, email, password_hash, profile_picture, biography, role_id) 
VALUES ('testUser', 'testUser@testUser@example.com', '$2a$10$sT4z5AHcw5CqATcCBIklqeSKNnW1XVnaQQ9KBCEdL0Q5DGbJoDnU2', 'https://example.com/profile.jpg', 'This is a test user', 1);

INSERT INTO categories (name, description) VALUES
('Technology', 'Posts about various technology topics'),
('Science', 'Posts about various science topics'),
('Art', 'Posts about various art topics');

INSERT INTO posts (title, content, user_id, post_category_id) VALUES
('Test Post 1', 'This is a test post', 1, 1),
('Test Post 2', 'This is a test post', 1, 2),
('Test Post 3', 'This is a test post', 1, 3);

INSERT INTO comments (content, post_id, user_id) VALUES
('This is a test comment', 1, 1),
('This is a test comment', 2, 1),
('This is a test comment', 3, 1);
