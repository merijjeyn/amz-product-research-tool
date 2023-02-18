CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email TEXT NOT NULL,
  username TEXT NOT NULL,
  credential TEXT NOT NULL
);

INSERT INTO users(email, username, credential)
	VALUES ('testemail', 'testusername', 'testcredential');