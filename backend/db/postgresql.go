package db

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func InitializePostgreDB() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=admin dbname=postgres sslmode=disable")
	if err != nil {
		panic(fmt.Errorf("db/postgresql.initializeDB: Error opening sql connection:\n%v", err))
	}
}

// TODO: Prevent sql injections (noob :D)

func GetUserWithGid(gid string) (*User, error) {
	if gid == "" {
		return nil, fmt.Errorf("db/postgresql.getUserWithCredentials received empty credential")
	}

	row := db.QueryRow("SELECT * FROM users WHERE gid = $1;", gid)

	var user User
	err := row.Scan(&user.id, &user.email, &user.username, &user.gid)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("db/postgresql.getUserWithGid error on rows.Scan:\n%v", err)
	}

	return &user, nil
}

func InsertUserIntoDB(email string, username string, gid string) error {
	if email == "" || username == "" || gid == "" {
		return fmt.Errorf("db/postgresql.insertUserIntoDB received empty credentials")
	}

	_, err := db.Exec("INSERT INTO users (email, username, gid) VALUES ($1, $2, $3)", email, username, gid)
	if err != nil {
		return fmt.Errorf("db/postgresql.insertUserIntoDB failed inserting user to db:\n%v", err)
	}
	return nil
}

func InsertUserSearchEntryIntoDB_SQL(gid string, searchText string) error {
	if gid == "" || searchText == "" {
		return fmt.Errorf("db/postgresql.InsertUserSearchEntryIntoDB_SQL: Invalid parameters")
	}

	_, err := db.Exec("INSERT INTO user_searches (user_id, search_text) VALUES ($1, $2)", gid, searchText)
	if err != nil {
		return fmt.Errorf("db/postgresql.InsertUserSearchEntryIntoDB_SQL: failed inserting user_search into db:\n%v", err)
	}
	return nil
}

func GetUserSearches_SQL(gid string) ([]string, error) {
	if gid == "" {
		return nil, fmt.Errorf("postgresql.GetUserSearches_SQL: Invalid parameters")
	}

	rows, err := db.Query("select search_text from user_searches where user_id = $1", gid)
	if err == sql.ErrNoRows {
		return []string{}, nil
	} else if err != nil {
		return nil, fmt.Errorf("postgresql.GetUserSearches_SQL: error executing sql query:\n%v", err)
	}
	defer rows.Close()

	result := []string{}
	for rows.Next() {
		var searchText string

		if err := rows.Scan(&searchText); err != nil {
			return nil, fmt.Errorf("postgresql.GetUserSearches_SQL: error scanning row:\n%v", err)
		}
		result = append(result, searchText)
	}

	return result, nil
}
