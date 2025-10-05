CREATE DATABASE IF NOT EXISTS freeboard;

USE freeboard;

CREATE TABLE IF NOT EXISTS timesheet (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    day DATE NOT NULL,
    hours FLOAT NOT NULL,
    UNIQUE KEY unique_user_day (username, day)
);