CREATE DATABASE IF NOT EXISTS driveDocuments;
USE driveDocuments;
DROP TABLE IF EXISTS documents;

CREATE TABLE documents (
    id VARCHAR(255) PRIMARY KEY,
    file_name VARCHAR(200) NOT NULL,
    extension VARCHAR(30) NOT NULL,
    owner_email VARCHAR(100) NOT NULL,
    visibility VARCHAR(100) NOT NULL
);
