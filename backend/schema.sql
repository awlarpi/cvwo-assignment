-- User Roles
CREATE TABLE user_roles (
  user_role_id SERIAL PRIMARY KEY,
  role_name VARCHAR(255) NOT NULL
);

-- Permissions
CREATE TABLE permissions (
  permission_id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT
);

-- Users
CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  registration_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  profile_picture TEXT,
  biography TEXT,
  last_login_date TIMESTAMP WITH TIME ZONE,
  is_active BOOLEAN DEFAULT TRUE,
  user_role_id INT REFERENCES user_roles(user_role_id)
);

-- Role Permissions (junction table between user_roles and permissions)
CREATE TABLE role_permissions (
  user_role_id INT REFERENCES user_roles(user_role_id),
  permission_id INT REFERENCES permissions(permission_id),
  PRIMARY KEY (user_role_id, permission_id)
);

-- Categories
CREATE TABLE categories (
  category_id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT
);

-- Posts
CREATE TABLE posts (
  post_id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  creation_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  user_id INT REFERENCES users(user_id),
  is_sticky BOOLEAN DEFAULT FALSE,
  is_locked BOOLEAN DEFAULT FALSE,
  post_category_id INT REFERENCES categories(category_id),
  additional_notes TEXT
);

-- Comments
CREATE TABLE comments (
  comment_id SERIAL PRIMARY KEY,
  content TEXT NOT NULL,
  creation_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  post_id INT REFERENCES posts(post_id),
  user_id INT REFERENCES users(user_id)
);

-- Routes
CREATE TABLE routes (
  route_id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  start_location TEXT NOT NULL,
  end_location TEXT NOT NULL,
  distance FLOAT,
  elevation_gain INT,
  route_map_link TEXT,
  user_id INT REFERENCES users(user_id)
);

-- Events
CREATE TABLE events (
  event_id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  event_date TIMESTAMP WITH TIME ZONE NOT NULL,
  meeting_point TEXT NOT NULL,
  route_id INT REFERENCES routes(route_id),
  creator_user_id INT REFERENCES users(user_id)
);

-- RSVPs
CREATE TABLE rsvps (
  rsvp_id SERIAL PRIMARY KEY,
  event_id INT REFERENCES events(event_id),
  user_id INT REFERENCES users(user_id),
  rsvp_status VARCHAR(255) NOT NULL,
  rsvp_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Private Messages
CREATE TABLE private_messages (
  message_id SERIAL PRIMARY KEY,
  content TEXT NOT NULL,
  sender_user_id INT REFERENCES users(user_id),
  receiver_user_id INT REFERENCES users(user_id),
  sent_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  is_read BOOLEAN DEFAULT FALSE
);

-- Forum Moderation Log
CREATE TABLE forum_moderation_log (
  log_id SERIAL PRIMARY KEY,
  action VARCHAR(255) NOT NULL,
  action_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  moderator_user_id INT REFERENCES users(user_id),
  affected_user_id INT REFERENCES users(user_id),
  post_id INT REFERENCES posts(post_id),
  comment_id INT REFERENCES comments(comment_id),
  reason TEXT
);

-- User Sessions
CREATE TABLE user_sessions (
  session_id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(user_id),
  creation_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  expiry_date TIMESTAMP WITH TIME ZONE,
  ip_address INET,
  user_agent TEXT
);

-- Notifications
CREATE TABLE notifications (
  notification_id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(user_id),
  content TEXT NOT NULL,
  creation_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  is_read BOOLEAN DEFAULT FALSE
);

-- Bookmarks
CREATE TABLE bookmarks (
  bookmark_id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(user_id),
  post_id INT REFERENCES posts(post_id)
);
