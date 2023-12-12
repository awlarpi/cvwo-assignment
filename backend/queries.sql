-- name: CreateUser :exec
INSERT INTO Users (Username, Email, PasswordHash, RegistrationDate, IsActive, UserRoleID)
VALUES (?, ?, ?, NOW(), TRUE, ?);

-- name: CreatePost :exec
INSERT INTO Posts (Title, Content, CreationDate, UserID, IsSticky, IsLocked, PostCategoryID, AdditionalNotes)
VALUES (?, ?, NOW(), ?, ?, ?, ?, ?);

-- name: GetPostsByUser :many
SELECT * FROM Posts WHERE UserID = ?;

-- name: GetUserComments :many
SELECT * FROM Comments WHERE UserID = ?;

-- name: CreateEvent :exec
INSERT INTO Events (Title, Description, EventDate, MeetingPoint, RouteID, CreatorUserID)
VALUES (?, ?, ?, ?, ?, ?);

-- name: RSVPToEvent :exec
INSERT INTO RSVPs (EventID, UserID, RSVPStatus, RSVPDate)
VALUES (?, ?, ?, NOW());

-- name: GetEventsByUser :many
SELECT * FROM Events WHERE CreatorUserID = ?;

-- name: GetRSVPsForEvent :many
SELECT * FROM RSVPs WHERE EventID = ?;
