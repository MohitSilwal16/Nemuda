CREATE DATABASE Nemuda;
USE Nemuda;

CREATE TABLE Users(
	Username VARCHAR(20) PRIMARY KEY ,
	Password VARCHAR(20) NOT NULL ,
	Token TEXT UNIQUE NOT NULL CHECK (Token <> "")
);