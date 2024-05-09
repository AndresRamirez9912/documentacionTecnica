DROP TABLE IF EXISTS documents;

CREATE TABLE documents (
    id VARCHAR(255) PRIMARY KEY,
    file_name VARCHAR(200) NOT NULL,
    extension VARCHAR(30) NOT NULL,
    owner_name VARCHAR(100) NOT NULL,
    visibility VARCHAR(100) NOT NULL
);
