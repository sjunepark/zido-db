package database

import "database/sql"

func CreateUserTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		email TEXT NOT NULL
	)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func DropUserTable(db *sql.DB) {
	query := `DROP TABLE IF EXISTS users`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

type User struct {
	ID       int
	Name     string
	Password string
	Email    string
}

func SelectUsers(db *sql.DB) {
	query := `SELECT * FROM "users"`
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Email)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	for _, user := range users {
		println(user.ID, user.Name, user.Password, user.Email)
	}
}

func InsertUser(db *sql.DB, username, password, email string) {
	query := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`
	_, err := db.Exec(query, username, password, email)
	if err != nil {
		panic(err)
	}
}
