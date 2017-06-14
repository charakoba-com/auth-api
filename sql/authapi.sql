-- Drop existing tables
DROP TABLE IF EXISTS users;

-- Create new tables
CREATE TABLE users (
        id VARCHAR(64) NOT NULL,
        username VARCHAR(128) NOT NULL,
        password VARCHAR(1024) NOT NULL,
        created_on DATETIME NOT NULL,
        modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
