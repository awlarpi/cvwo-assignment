-- name: CreateUser :exec
-- Inserts a new user into the users table.
INSERT INTO users (
  username, 
  email, 
  password_hash, 
  user_role_id
) VALUES (
  $1, $2, $3, $4
);

-- name: CreatePost :exec
-- Inserts a new post into the posts table.
INSERT INTO posts (
  title, 
  content, 
  user_id, 
  is_sticky, 
  is_locked, 
  post_category_id, 
  additional_notes
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
);

-- name: CreateComment :exec
-- Inserts a new comment into the comments table.
INSERT INTO comments (
  content, 
  post_id, 
  user_id
) VALUES (
  $1, $2, $3
);

-- name: CreateEvent :exec
-- Inserts a new event into the events table.
INSERT INTO events (
  title, 
  description, 
  event_date, 
  meeting_point, 
  route_id, 
  creator_user_id
) VALUES (
  $1, $2, $3, $4, $5, $6
);

-- name: RSVPToEvent :exec
-- Inserts a new RSVP record for an event.
INSERT INTO rsvps (
  event_id, 
  user_id, 
  rsvp_status
) VALUES (
  $1, $2, $3
);
