CREATE DATABASE Nemuda;
USE Nemuda;

CREATE TABLE Users(
	Username VARCHAR(20) PRIMARY KEY CHECK (Username <> ""),
	Password VARCHAR(20) NOT NULL CHECK (Password <> ""),
	Token CHAR(8) UNIQUE NOT NULL
);