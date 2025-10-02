CREATE DATABASE IF NOT EXISTS freeboard;

USE freeboard;

CREATE TABLE IF NOT EXISTS timesheet (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    day DATE NOT NULL,
    hours FLOAT NOT NULL,
    UNIQUE KEY unique_user_day (user_id, day)
);