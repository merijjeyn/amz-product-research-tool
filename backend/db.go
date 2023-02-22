package main

import (
	"database/sql"
	"fmt"
)

// TODO: Prevent sql injections (noob :D)

func getUserWithGid(gid string) (*User, error) {
	if gid == "" {
		return nil, fmt.Errorf("db.getUserWithCredentials: received empty credential")
	}

	row := db.QueryRow("SELECT * FROM users WHERE gid = $1;", gid)

	var user User
	err := row.Scan(&user.id, &user.email, &user.username, &user.gid)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("db.getUserWithGid: Something went wrong with rows.Scan:\n%v", err)
	}

	return &user, nil
}

func insertUserIntoDB(email string, username string, gid string) error {
	if email == "" || username == "" || gid == "" {
		return fmt.Errorf("db.insertUserIntoDB: received empty credentials")
	}

	_, err := db.Exec("INSERT INTO users (email, username, gid) VALUES ($1, $2, $3)", email, username, gid)
	if err != nil {
		return fmt.Errorf("db.insertUserIntoDB: failed inserting user to db:\n%v", err)
	}
	return nil
}
