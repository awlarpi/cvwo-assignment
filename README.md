Some Docker commands:
- docker compose down
- docker compose build
- docker compose up -d
- docker compose up -d --build backend

ERO:
**Entities:**

1. User

   - Attributes: UserID (PK), Username, Email, PasswordHash, RegistrationDate, ProfilePicture, Biography, LastLoginDate, IsActive, UserRoleID (FK)
   - Relationships: Users can create Posts, Comments, Routes, and Events. Users can also attend Events.

2. Post

   - Attributes: PostID (PK), Title, Content, CreationDate, UserID (FK), IsSticky, IsLocked, PostCategoryID (FK), AdditionalNotes
   - Relationships: A Post is created by one User. Posts can have multiple Comments.

3. Comment

   - Attributes: CommentID (PK), Content, CreationDate, PostID (FK), UserID (FK)
   - Relationships: A Comment is made on a Post by a User.

4. Route

   - Attributes: RouteID (PK), Name, Description, StartLocation, EndLocation, Distance, ElevationGain, RouteMapLink, UserID (FK)
   - Relationships: Routes are created by Users. They can be referenced by Events.

5. Event

   - Attributes: EventID (PK), Title, Description, EventDate, MeetingPoint, RouteID (FK, nullable), CreatorUserID (FK)
   - Relationships: Events can be organized by Users and may reference a Route. Users can RSVP to Events.

6. RSVP

   - Attributes: RSVPID (PK), EventID (FK), UserID (FK), RSVPStatus, RSVPDate
   - Relationships: Indicates which Users have RSVP'd to which Events and their attendance status.

7. Category (optional, not mentioned but useful for organizing discussions)

   - Attributes: CategoryID (PK), Name, Description
   - Relationships: Posts can belong to Categories.

8. **UserRole**

   - Attributes: UserRoleID (PK), RoleName
   - Relationships: Users have one or many UserRoles.

9. **Permission**

   - Attributes: PermissionID (PK), Name, Description
   - Relationships: Roles have many Permissions.

10. **PrivateMessage**

- Attributes: MessageID (PK), Content, SenderUserID (FK), ReceiverUserID (FK), SentDate, IsRead
- Relationships: Sent from one User to another User.

11. **ForumModerationLog**

- Attributes: LogID (PK), Action, ActionDate, ModeratorUserID (FK), AffectedUserID (FK, nullable), PostID (FK, nullable), CommentID (FK, nullable), Reason
- Relationships: Tracks moderation actions taken by Users with Moderator permissions on Posts, Comments, and Users.

12. **UserSessions**

- Attributes: SessionID (PK), UserID (FK), CreationDate, ExpiryDate, IPAddress, UserAgent
- Relationships: Tracks User login sessions for security purposes.

13. **Notification**

- Attributes: NotificationID (PK), UserID (FK), Content, CreationDate, IsRead
- Relationships: Users receive Notifications about various activities/events.

14. **Bookmark**

- Attributes: BookmarkID (PK), UserID (FK), PostID (FK)
- Relationships: Users can bookmark Posts to find them easily later.

**Relationships:**

- One-to-Many:

  - One User can create many Posts.
  - One User can create many Comments.
  - One Post can have many Comments.
  - One User can create many Routes.
  - One User can create many Events.
  - One User can have many RSVPs to different Events.
  - One Event can have many RSVPs from Users.
  - One Category can include many Posts (if using categories).
  - One User can send many PrivateMessages.
  - One User can receive many PrivateMessages.
  - One User can have many UserSessions.
  - One user can receive many Notifications.
  - One User can have many Bookmarks.
  - One User can have many Roles.
  - One Role can have many Permissions

- Many-to-One:
  - Many Comments belong to one Post.
  - Many Posts may be associated with one Route (if a post can be about a specific route).
  - Many RSVPs correspond to one Event.

**Additional Features to Consider:**

1. User Profile:
   - Allow users to fill out additional information like favorite bike brands, cycling disciplines, etc.
2. Social Media Integration:
   - Enable sharing of Posts, Routes, and Events on social media.
3. Search Functionality:
   - Implement a robust search engine to allow users to search for Posts, Routes, Events, and Users.
4. Analytics:
   - Track user engagement and activity through analytics for community growth metrics.

**Security and Data Integrity Considerations:**

1. All passwords should be encrypted using strong cryptographic algorithms like bcrypt.
2. Enable two-factor authentication (2FA) for additional security.
3. Implement Cross-Site Request Forgery (CSRF) tokens in forms to prevent cross-site attacks.
4. Apply Cross-Origin Resource Sharing (CORS) policies to control which domains are allowed to access your API.
5. Validate and sanitize all inputs to prevent SQL injection and XSS attacks.
6. Permissions and roles should be used to control access to various features and administrative functions.
