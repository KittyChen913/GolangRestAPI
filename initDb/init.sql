-- create DB
CREATE DATABASE DemoDb;
GO

USE DemoDb;
GO
-- create [Admins] tables
CREATE TABLE Admins (
    AdminId tinyint IDENTITY(1,1) PRIMARY KEY NOT NULL,
    Name NVARCHAR(50) NOT NULL,
    Password VARCHAR(500) NOT NULL,
    Email NVARCHAR(200)
);
GO

-- create [Users] tables
CREATE TABLE Users (
    Id tinyint IDENTITY(1,1) PRIMARY KEY NOT NULL,
    Name NVARCHAR(50) NOT NULL,
    Age INT NULL,
    CreateDateTime DATETIME NULL DEFAULT GETDATE()
);
GO