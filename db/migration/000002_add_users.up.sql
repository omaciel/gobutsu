CREATE TABLE IF NOT EXISTS users (
	username varchar(255) PRIMARY KEY,
	hashed_password varchar(255) NOT NULL,
	email varchar(255) UNIQUE NOT NULL,
	created_on timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE testcases
ADD username varchar(255) NOT NULL;

ALTER TABLE testcases
ADD CONSTRAINT testcases_username_fkey
FOREIGN KEY (username) REFERENCES users (username);