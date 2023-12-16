-- name: CreateUser :exec
-- Create a new user
INSERT INTO users (username, email, password_hash, profile_picture, biography, role_id)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetUser :one
-- Get a user by id, no password_hash
SELECT user_id, username, email, registration_date, profile_picture, biography, last_login_date, is_active, role_id FROM users WHERE user_id = $1;

-- name: GetUserWithRoleName :one
-- Get a user by id, no password_hash, with role_name
SELECT users.user_id, users.username, users.email, users.registration_date, users.profile_picture, users.biography, users.last_login_date, users.is_active, users.role_id, roles.role_name
FROM users
INNER JOIN roles ON users.role_id = roles.role_id
WHERE users.user_id = $1;

-- name: GetAllUsers :many
-- Get all users, no password_hash
SELECT user_id, username, email, registration_date, profile_picture, biography, last_login_date, is_active, role_id FROM users;

-- name: UpdateUserExcludingSensitive :exec
-- Update a user's details, no password_hash
UPDATE users SET username = $2, email = $3, profile_picture = $4, biography = $5 WHERE user_id = $1;

-- name: UpdateUserPassword :exec
-- Update a user's password_hash
UPDATE users SET password_hash = $2 WHERE user_id = $1;

-- name: UpdateUserRole :exec
-- Update a user's role
UPDATE users SET role_id = $2 WHERE user_id = $1;

-- name: DeleteUser :exec
-- Delete a user by id
DELETE FROM users WHERE user_id = $1;

-- name: GetUserByUsername :one
-- Get a user by username
SELECT * FROM users WHERE username = $1;

-- name: GetUserByEmail :one
-- Get a user by email
SELECT * FROM users WHERE email = $1;

-- name: UpdateLastLogin :exec
-- Update a user's last login date
UPDATE users SET last_login_date = CURRENT_TIMESTAMP WHERE user_id = $1;

-- name: DeactivateUser :exec
-- Deactivate a user
UPDATE users SET is_active = FALSE WHERE user_id = $1;

-- name: ActivateUser :exec
-- Activate a user
UPDATE users SET is_active = TRUE WHERE user_id = $1;

------------------------------------------------------------------------------------------------------------------------

-- name: CreateRole :exec
-- Create a new role
INSERT INTO roles (role_name) VALUES ($1);

-- name: GetRole :one
-- Get a role by id
SELECT * FROM roles WHERE role_id = $1;

-- name: UpdateRole :exec
-- Update a role by id
UPDATE roles SET role_name = $2 WHERE role_id = $1;

-- name: DeleteRole :exec
-- Delete a role by id
DELETE FROM roles WHERE role_id = $1;

-- name: ListRoles :many
-- Get all roles
SELECT * FROM roles;

------------------------------------------------------------------------------------------------------------------------

-- name: CreateCategory :exec
-- Create a new category
INSERT INTO categories (name, description)
VALUES ($1, $2);

-- name: GetCategory :one
-- Get a category by id
SELECT * FROM categories WHERE category_id = $1;

-- name: GetAllCategories :many
-- Get all categories
SELECT * FROM categories;

-- name: UpdateCategory :exec
-- Update a category by id
UPDATE categories SET name = $2, description = $3 WHERE category_id = $1;

-- name: DeleteCategory :exec
-- Delete a category by id
DELETE FROM categories WHERE category_id = $1;

------------------------------------------------------------------------------------------------------------------------

-- name: CreatePost :exec
INSERT INTO posts (title, content, user_id, post_category_id, additional_notes)
VALUES ($1, $2, $3, $4, $5);

-- name: GetPost :one
SELECT * FROM posts WHERE post_id = $1;

-- name: GetAllPosts :many
SELECT * FROM posts ORDER BY creation_date DESC;

-- name: GetPostsByUser :many
SELECT * FROM posts WHERE user_id = $1 ORDER BY creation_date DESC;

-- name: GetPostsByCategory :many
SELECT * FROM posts WHERE post_category_id = $1 ORDER BY creation_date DESC;

-- name: UpdatePost :exec
UPDATE posts SET title = $2, content = $3, user_id = $4, post_category_id = $5, additional_notes = $6 WHERE post_id = $1;

-- name: DeletePost :exec
DELETE FROM posts WHERE post_id = $1;

-- name: DeletePostByPostIdAndUserId :exec
DELETE FROM posts WHERE post_id = $1 AND user_id = $2;

-- name: GetStickyPosts :many
SELECT * FROM posts WHERE is_sticky = TRUE ORDER BY creation_date DESC;

-- name: GetLockedPosts :many
SELECT * FROM posts WHERE is_locked = TRUE ORDER BY creation_date DESC;

-- name: GetUnlockedPosts :many
SELECT * FROM posts WHERE is_locked = FALSE ORDER BY creation_date DESC;

-- name: LockPost :exec
UPDATE posts SET is_locked = TRUE WHERE post_id = $1;

-- name: UnlockPost :exec
UPDATE posts SET is_locked = FALSE WHERE post_id = $1;

-- name: StickyPost :exec
UPDATE posts SET is_sticky = TRUE WHERE post_id = $1;

-- name: UnstickyPost :exec
UPDATE posts SET is_sticky = FALSE WHERE post_id = $1;

------------------------------------------------------------------------------------------------------------------------

-- name: GetComment :one
-- Get a single comment by its ID
SELECT * FROM comments WHERE comment_id = $1;

-- name: GetAllComments :many
-- Get all comments, ordered by creation date
SELECT * FROM comments ORDER BY creation_date DESC;

-- name: GetCommentsByPost :many
-- Get all comments for a specific post, ordered by creation date
SELECT * FROM comments WHERE post_id = $1 ORDER BY creation_date DESC;

-- name: GetCommentsByUser :many
-- Get all comments made by a specific user, ordered by creation date
SELECT * FROM comments WHERE user_id = $1 ORDER BY creation_date DESC;

-- name: CreateComment :one
-- Create a new comment
INSERT INTO comments (content, post_id, user_id) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateComment :one
-- Update a comment's content
UPDATE comments SET content = $2 WHERE comment_id = $1 RETURNING *;

-- name: UpdateCommentByCommentIdAndUserId :one
-- Update a comment's content by its ID and the ID of the user who made it
UPDATE comments SET content = $3 WHERE comment_id = $1 AND user_id = $2 RETURNING *;

-- name: DeleteComment :exec
-- Delete a comment by its ID
DELETE FROM comments WHERE comment_id = $1;

-- name: DeleteCommentByCommentIdAndUserId :exec
-- Delete a comment by its ID and the ID of the user who made it
DELETE FROM comments WHERE comment_id = $1 AND user_id = $2;

------------------------------------------------------------------------------------------------------------------------

-- name: GetRoutes :many
-- Get all routes, ordered by name
SELECT * FROM routes ORDER BY name;

-- name: GetRouteByID :one
-- Get a specific route by its ID
SELECT * FROM routes WHERE route_id = $1;

-- name: CreateRoute :exec
-- Create a new route
INSERT INTO routes (name, description, start_location, end_location, distance, elevation_gain, route_map_link, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: UpdateRoute :exec
-- Update a specific route
UPDATE routes SET name = $2, description = $3, start_location = $4, end_location = $5, distance = $6, elevation_gain = $7, route_map_link = $8, user_id = $9
WHERE route_id = $1;

-- name: DeleteRoute :exec
-- Delete a specific route
DELETE FROM routes WHERE route_id = $1;

-- name: GetRoutesByUserID :many
-- Get all routes created by a specific user
SELECT * FROM routes WHERE user_id = $1 ORDER BY name;

-- name: GetRoutesByDistance :many
-- Get all routes, ordered by distance
SELECT * FROM routes ORDER BY distance;

------------------------------------------------------------------------------------------------------------------------

-- name: GetEvents :many
-- Get all events, ordered by event date
SELECT * FROM events ORDER BY event_date;

-- name: GetEventByID :one
-- Get a specific event by its ID
SELECT * FROM events WHERE event_id = $1;

-- name: CreateEvent :exec
-- Create a new event
INSERT INTO events (title, description, event_date, meeting_point, route_id, creator_user_id)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UpdateEvent :exec
-- Update a specific event
UPDATE events SET title = $2, description = $3, event_date = $4, meeting_point = $5, route_id = $6, creator_user_id = $7
WHERE event_id = $1;

-- name: DeleteEvent :exec
-- Delete a specific event
DELETE FROM events WHERE event_id = $1;

-- name: GetEventsByUserID :many
-- Get all events created by a specific user
SELECT * FROM events WHERE creator_user_id = $1 ORDER BY event_date;

-- name: GetEventsByRouteID :many
-- Get all events for a specific route
SELECT * FROM events WHERE route_id = $1 ORDER BY event_date;

-- name: GetEventsByDate :many
-- Get all events, ordered by event date
SELECT * FROM events ORDER BY event_date;

-- name: GetEventsByDateRange :many
-- Get all events within a date range
SELECT * FROM events WHERE event_date >= $1 AND event_date <= $2 ORDER BY event_date;

-- name: GetEventsByDateRangeAndRouteID :many
-- Get all events within a date range for a specific route
SELECT * FROM events WHERE event_date >= $1 AND event_date <= $2 AND route_id = $3 ORDER BY event_date;

-- name: GetEventsByDateRangeAndUserID :many
-- Get all events within a date range for a specific user
SELECT * FROM events WHERE event_date >= $1 AND event_date <= $2 AND creator_user_id = $3 ORDER BY event_date;

-- name: GetEventsByDateRangeAndRouteIDAndUserID :many
-- Get all events within a date range for a specific route and user
SELECT * FROM events WHERE event_date >= $1 AND event_date <= $2 AND route_id = $3 AND creator_user_id = $4 ORDER BY event_date;

------------------------------------------------------------------------------------------------------------------------

-- name: GetRSVPs :many
-- Get all RSVPs
SELECT * FROM rsvps;

-- name: GetRSVPByID :one
-- Get a specific RSVP by its ID
SELECT * FROM rsvps WHERE rsvp_id = $1;

-- name: CreateRSVP :exec
-- Create a new RSVP
INSERT INTO rsvps (event_id, user_id) VALUES ($1, $2);

-- name: UpdateRSVP :exec
-- Update a specific RSVP
UPDATE rsvps SET event_id = $2, user_id = $3 WHERE rsvp_id = $1;

-- name: DeleteRSVP :exec
-- Delete a specific RSVP
DELETE FROM rsvps WHERE rsvp_id = $1;

-- name: GetRSVPsByEventID :many
-- Get all RSVPs for a specific event
SELECT * FROM rsvps WHERE event_id = $1;

-- name: GetRSVPsByUserID :many
-- Get all RSVPs for a specific user
SELECT * FROM rsvps WHERE user_id = $1;

-- name: GetRSVPsByEventIDAndUserID :many
-- Get all RSVPs for a specific event and user
SELECT * FROM rsvps WHERE event_id = $1 AND user_id = $2;

------------------------------------------------------------------------------------------------------------------------

-- name: CreatePrivateMessage :exec
-- Create a new private message
INSERT INTO private_messages (content, sender_user_id, receiver_user_id)
VALUES ($1, $2, $3);

-- name: GetAllPrivateMessages :many
-- Get all private messages
SELECT * FROM private_messages;

-- name: GetPrivateMessageById :one
-- Get a private message by its ID
SELECT * FROM private_messages WHERE message_id = $1;

-- name: GetPrivateMessagesBySenderId :many
-- Get all private messages sent by a specific user
SELECT * FROM private_messages WHERE sender_user_id = $1 ORDER BY sent_date DESC;

-- name: GetPrivateMessagesByReceiverId :many
-- Get all private messages received by a specific user
SELECT * FROM private_messages WHERE receiver_user_id = $1 ORDER BY sent_date DESC;

-- name: GetUnreadPrivateMessagesByReceiverId :many
-- Get all unread private messages received by a specific user
SELECT * FROM private_messages WHERE receiver_user_id = $1 AND is_read = FALSE ORDER BY sent_date DESC;

-- name: UpdatePrivateMessage :exec
-- Update the content of a private message
UPDATE private_messages SET content = $2 WHERE message_id = $1;

-- name: MarkPrivateMessageAsRead :exec
-- Mark a private message as read
UPDATE private_messages SET is_read = TRUE WHERE message_id = $1;

-- name: DeletePrivateMessage :exec
-- Delete a private message by its ID
DELETE FROM private_messages WHERE message_id = $1;

-- name: DeletePrivateMessagesBySenderId :exec
-- Delete all private messages sent by a specific user
DELETE FROM private_messages WHERE sender_user_id = $1;

-- name: DeletePrivateMessagesByReceiverId :exec
-- Delete all private messages received by a specific user
DELETE FROM private_messages WHERE receiver_user_id = $1;

-- name: DeletePrivateMessagesBySenderIdAndReceiverId :exec
-- Delete all private messages sent by a specific user to another specific user
DELETE FROM private_messages WHERE sender_user_id = $1 AND receiver_user_id = $2;

------------------------------------------------------------------------------------------------------------------------

-- name: CreateLog :exec
INSERT INTO forum_moderation_log (action, moderator_user_id, affected_user_id, post_id, comment_id, reason)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetLogById :one
SELECT * FROM forum_moderation_log WHERE log_id = $1;

-- name: GetAllLogs :many
SELECT * FROM forum_moderation_log;

-- name: UpdateLog :exec
UPDATE forum_moderation_log
SET action = $2, moderator_user_id = $3, affected_user_id = $4, post_id = $5, comment_id = $6, reason = $7
WHERE log_id = $1;

-- name: DeleteLog :exec
DELETE FROM forum_moderation_log WHERE log_id = $1;

-- name: GetLogsByAction :many
SELECT * FROM forum_moderation_log WHERE action = $1;

-- name: GetLogsByModerator :many
SELECT * FROM forum_moderation_log WHERE moderator_user_id = $1 ORDER BY action_date DESC;

-- name: GetLogsByAffectedUser :many
SELECT * FROM forum_moderation_log WHERE affected_user_id = $1 ORDER BY action_date DESC;

-- name: GetLogsByPost :many
SELECT * FROM forum_moderation_log WHERE post_id = $1;

-- name: GetLogsByComment :many
SELECT * FROM forum_moderation_log WHERE comment_id = $1;

-- name: GetLogsByReason :many
SELECT * FROM forum_moderation_log WHERE reason = $1;

------------------------------------------------------------------------------------------------------------------------

-- name: GetUserSession :one
-- Get a single user session by session_id
SELECT * FROM user_sessions WHERE session_id = $1;

-- name: GetUserSessionAndRoleName :one
-- Get a single user session by session_id, with role_name
SELECT user_sessions.session_id, user_sessions.user_id, user_sessions.expiry_date, user_sessions.ip_address, user_sessions.user_agent, user_sessions.creation_date, users.role_id, roles.role_name
FROM user_sessions
INNER JOIN users ON user_sessions.user_id = users.user_id
INNER JOIN roles ON users.role_id = roles.role_id
WHERE session_id = $1;

-- name: GetUserSessionsByUserId :many
-- Get all sessions for a specific user_id
SELECT * FROM user_sessions WHERE user_id = $1;

-- name: GetUserSessionsByDate :many
-- Get all sessions created on a specific date
SELECT * FROM user_sessions WHERE DATE(creation_date) = $1;

-- name: GetUserSessionsByIP :many
-- Get all sessions from a specific IP address
SELECT * FROM user_sessions WHERE ip_address = $1;

-- name: CreateUserSession :one
-- Insert a new user session and return the created session
INSERT INTO user_sessions (session_id, user_id, expiry_date, ip_address, user_agent) VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateUserSession :exec
-- Update a user session
UPDATE user_sessions SET user_id = $2, expiry_date = $3, ip_address = $4, user_agent = $5 WHERE session_id = $1;

-- name: DeleteUserSession :exec
-- Delete a user session
DELETE FROM user_sessions WHERE session_id = $1;

-- name: GetUserSessionsOrderedByDate :many
-- Get all user sessions ordered by creation_date
SELECT * FROM user_sessions ORDER BY creation_date;

-- name: GetUserSessionsFilteredByDate :many
-- Get all user sessions after a specific date
SELECT * FROM user_sessions WHERE creation_date > $1;

-- name: InvalidateUserSession :exec
-- Invalidate a user session by setting the expiry_date to a past date
UPDATE user_sessions SET expiry_date = TIMESTAMP '1970-01-01 00:00:00' WHERE session_id = $1;

------------------------------------------------------------------------------------------------------------------------

-- @name GetNotifications :many
-- Get all notifications, ordered by creation_date
SELECT * FROM notifications ORDER BY creation_date DESC;

-- @name GetNotificationById :one
-- Get a single notification by its ID
SELECT * FROM notifications WHERE notification_id = $1;

-- @name GetNotificationsByUserId :many
-- Get all notifications for a specific user, ordered by creation_date
SELECT * FROM notifications WHERE user_id = $1 ORDER BY creation_date DESC;

-- @name GetUnreadNotificationsByUserId :many
-- Get all unread notifications for a specific user, ordered by creation_date
SELECT * FROM notifications WHERE user_id = $1 AND is_read = FALSE ORDER BY creation_date DESC;

-- @name CreateNotification :exec
-- Create a new notification
INSERT INTO notifications (user_id, content) VALUES ($1, $2);

-- @name UpdateNotification :exec
-- Update a notification's content
UPDATE notifications SET content = $2 WHERE notification_id = $1;

-- @name MarkNotificationAsRead :exec
-- Mark a notification as read
UPDATE notifications SET is_read = TRUE WHERE notification_id = $1;

-- @name DeleteNotification :exec
-- Delete a notification by its ID
DELETE FROM notifications WHERE notification_id = $1;

-- @name DeleteNotificationsByUserId :exec
-- Delete all notifications for a specific user
DELETE FROM notifications WHERE user_id = $1;

------------------------------------------------------------------------------------------------------------------------

-- name: CreateBookmark :exec
-- Create a new bookmark
INSERT INTO bookmarks (user_id, post_id) VALUES ($1, $2);

-- name: GetBookmark :one
-- Get a bookmark by its id
SELECT * FROM bookmarks WHERE bookmark_id = $1;

-- name: GetBookmarksByUser :many
-- Get all bookmarks of a user, ordered by bookmark_id
SELECT * FROM bookmarks WHERE user_id = $1 ORDER BY bookmark_id;

-- name: GetBookmarksByPost :many
-- Get all bookmarks for a post, ordered by user_id
SELECT * FROM bookmarks WHERE post_id = $1 ORDER BY user_id;

-- name: UpdateBookmark :exec
-- Update a bookmark
UPDATE bookmarks SET user_id = $2, post_id = $3 WHERE bookmark_id = $1;

-- name: DeleteBookmark :exec
-- Delete a bookmark by its id
DELETE FROM bookmarks WHERE bookmark_id = $1;

-- name: DeleteBookmarksByUser :exec
-- Delete all bookmarks of a user
DELETE FROM bookmarks WHERE user_id = $1;
