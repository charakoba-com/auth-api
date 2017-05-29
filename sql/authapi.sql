--- Drop existing tables
DROP TABLE IF EXISTS users;

--- Create new tables
CREATE TABLE IF NOT EXISTS users (
        username VARCHAR(256) NOT NULL,
        password VARCHAR(1024) NOT NULL,
        created_on DATETIME NOT NULL,
        modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY(username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
