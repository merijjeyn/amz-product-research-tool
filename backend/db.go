package main

import (
	"database/sql"
	"fmt"
)

// TODO: Prevent sql injections (noob :D)

func getUserWithCredential(credential string) (*User, error) {
	row := db.QueryRow("SELECT id, email, username FROM users WHERE credential = $1", credential)

	var user User
	err := row.Scan(&user.id, &user.email, &user.username)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("db.getUserWithCredentials: Something went wrong with rows.Scan:\n%v", err)
	}

	return &user, nil
}

func insertUserIntoDB(email string, username string, credential string) error {
	_, err := db.Exec("INSERT INTO users (email, username, credential) VALUES ($1, $2, $3)", email, username, credential)
	if err != nil {
		return fmt.Errorf("db.insertUser: failed inserting user to db:\n%v", err)
	}
	return nil
}
