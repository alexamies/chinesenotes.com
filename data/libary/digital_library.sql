USE cse_dict;

CREATE TABLE IF NOT EXISTS user (
  UserID INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
  UserName VARCHAR(100) NOT NULL,
  Email VARCHAR(100), 
  FullName VARCHAR(100) NOT NULL,
  Role VARCHAR(100) NOT NULL DEFAULT "user",
  PasswordNeedsReset TINYINT(1) NOT NULL DEFAULT 1,
  Organization VARCHAR(100),
  Position VARCHAR(100),
  Location VARCHAR(100));

CREATE TABLE IF NOT EXISTS passwd (
  UserID INT NOT NULL PRIMARY KEY,
  Password VARCHAR(100),
  CONSTRAINT `fk_user_passwd`
    FOREIGN KEY (UserID) REFERENCES user (UserID)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
  );

CREATE TABLE IF NOT EXISTS session (
  SessionID VARCHAR(100) NOT NULL PRIMARY KEY, 
  UserID INT NOT NULL,
  Active INT NOT NULL DEFAULT 1,
  Authenticated INT NOT NULL DEFAULT 0,
  DailyVisits INT NOT NULL DEFAULT 1,
  Started TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  Updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT `fk_user_session`
    FOREIGN KEY (UserID) REFERENCES user (UserID)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
  );

INSERT INTO user (UserID, UserName, Email, FullName, Role, PasswordNeedsReset, Organization, Position, Location) 
VALUES (1, 'unauthenticated_user', "", "Unauthenticated User", "none", 0, "", "", "");

INSERT INTO user (UserID, UserName, Email, FullName, Role, PasswordNeedsReset, Organization, Position, Location) 
VALUES (2, 'guest', "", "Guest User", "user", 0, "", "", "");

INSERT INTO passwd (UserID, Password) 
VALUES (2, '84983c60f7daadc1cb8698621f802c0d9f9a3c3c295c810748fb048115c186ec');

INSERT INTO user (UserID, UserName, Email, FullName, Role, PasswordNeedsReset, Organization, Position, Location) 
VALUES (3, 'admin', "", "Administrator", "admin", 0, "", "", "");

INSERT INTO passwd (UserID, Password) 
VALUES (3, '84983c60f7daadc1cb8698621f802c0d9f9a3c3c295c810748fb048115c186ec');
