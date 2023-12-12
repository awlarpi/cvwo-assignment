-- Users table
CREATE TABLE Users (
    UserID INT PRIMARY KEY AUTO_INCREMENT,
    Username VARCHAR(50) UNIQUE NOT NULL,
    Email VARCHAR(255) UNIQUE NOT NULL,
    PasswordHash VARCHAR(255) NOT NULL,
    RegistrationDate DATETIME NOT NULL,
    ProfilePicture VARCHAR(255),
    Biography TEXT,
    LastLoginDate DATETIME,
    IsActive BOOLEAN NOT NULL,
    UserRoleID INT,
    FOREIGN KEY (UserRoleID) REFERENCES UserRoles(UserRoleID)
);

-- Posts table
CREATE TABLE Posts (
    PostID INT PRIMARY KEY AUTO_INCREMENT,
    Title VARCHAR(255) NOT NULL,
    Content TEXT NOT NULL,
    CreationDate DATETIME NOT NULL,
    UserID INT,
    IsSticky BOOLEAN NOT NULL,
    IsLocked BOOLEAN NOT NULL,
    PostCategoryID INT,
    AdditionalNotes TEXT,
    FOREIGN KEY (UserID) REFERENCES Users(UserID),
    FOREIGN KEY (PostCategoryID) REFERENCES Categories(CategoryID)
);

-- Comments table
CREATE TABLE Comments (
    CommentID INT PRIMARY KEY AUTO_INCREMENT,
    Content TEXT NOT NULL,
    CreationDate DATETIME NOT NULL,
    PostID INT,
    UserID INT,
    FOREIGN KEY (PostID) REFERENCES Posts(PostID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- Routes table
CREATE TABLE Routes (
    RouteID INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(255) NOT NULL,
    Description TEXT,
    StartLocation VARCHAR(255) NOT NULL,
    EndLocation VARCHAR(255) NOT NULL,
    Distance FLOAT NOT NULL,
    ElevationGain INT,
    RouteMapLink VARCHAR(255),
    UserID INT,
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- Events table
CREATE TABLE Events (
    EventID INT PRIMARY KEY AUTO_INCREMENT,
    Title VARCHAR(255) NOT NULL,
    Description TEXT,
    EventDate DATETIME NOT NULL,
    MeetingPoint VARCHAR(255) NOT NULL,
    RouteID INT,
    CreatorUserID INT,
    FOREIGN KEY (RouteID) REFERENCES Routes(RouteID),
    FOREIGN KEY (CreatorUserID) REFERENCES Users(UserID)
);

-- RSVPs table
CREATE TABLE RSVPs (
    RSVPID INT PRIMARY KEY AUTO_INCREMENT,
    EventID INT,
    UserID INT,
    RSVPStatus VARCHAR(50) NOT NULL,
    RSVPDate DATETIME NOT NULL,
    FOREIGN KEY (EventID) REFERENCES Events(EventID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- Categories table (optional)
CREATE TABLE Categories (
    CategoryID INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(50) NOT NULL,
    Description TEXT
);

-- UserRoles table
CREATE TABLE UserRoles (
    UserRoleID INT PRIMARY KEY AUTO_INCREMENT,
    RoleName VARCHAR(50) NOT NULL
);

-- Permissions table
CREATE TABLE Permissions (
    PermissionID INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(50) NOT NULL,
    Description TEXT,
    UserRoleID INT,
    FOREIGN KEY (UserRoleID) REFERENCES UserRoles(UserRoleID)
);

-- PrivateMessages table
CREATE TABLE PrivateMessages (
    MessageID INT PRIMARY KEY AUTO_INCREMENT,
    Content TEXT NOT NULL,
    SenderUserID INT,
    ReceiverUserID INT,
    SentDate DATETIME NOT NULL,
    IsRead BOOLEAN NOT NULL,
    FOREIGN KEY (SenderUserID) REFERENCES Users(UserID),
    FOREIGN KEY (ReceiverUserID) REFERENCES Users(UserID)
);

-- ForumModerationLog table
CREATE TABLE ForumModerationLog (
    LogID INT PRIMARY KEY AUTO_INCREMENT,
    Action VARCHAR(255) NOT NULL,
    ActionDate DATETIME NOT NULL,
    ModeratorUserID INT,
    AffectedUserID INT,
    PostID INT,
    CommentID INT,
    Reason TEXT,
    FOREIGN KEY (ModeratorUserID) REFERENCES Users(UserID),
    FOREIGN KEY (AffectedUserID) REFERENCES Users(UserID),
    FOREIGN KEY (PostID) REFERENCES Posts(PostID),
    FOREIGN KEY (CommentID) REFERENCES Comments(CommentID)
);

-- UserSessions table
CREATE TABLE UserSessions (
    SessionID INT PRIMARY KEY AUTO_INCREMENT,
    UserID INT,
    CreationDate DATETIME NOT NULL,
    ExpiryDate DATETIME NOT NULL,
    IPAddress VARCHAR(50),
    UserAgent VARCHAR(255),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- Notifications table
CREATE TABLE Notifications (
    NotificationID INT PRIMARY KEY AUTO_INCREMENT,
    UserID INT,
    Content TEXT NOT NULL,
    CreationDate DATETIME NOT NULL,
    IsRead BOOLEAN NOT NULL,
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- Bookmarks table
CREATE TABLE Bookmarks (
    BookmarkID INT PRIMARY KEY AUTO_INCREMENT,
    UserID INT,
    PostID INT,
    FOREIGN KEY (UserID) REFERENCES Users(UserID),
    FOREIGN KEY (PostID) REFERENCES Posts(PostID)
);
