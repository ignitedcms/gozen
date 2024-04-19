package users

import (
	"gozen/db"
	"time"
)

// User represents a User in the system
type User struct {
	Id         string
	Name       string
	Email      string
	Password   string
	Token      string
	Created_at string
	Updated_at string
}

// Insert inserts a new User into the database
func Create(name string, email string, password string) (int64, error) {
	stmt, err := db.DB.Prepare("INSERT INTO users(name, email, password,token, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(name, email, password, "fdsfds", time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

// Update updates an existing User in the database
func Update(id string, name string, email string, password string) error {
	stmt, err := db.DB.Prepare("UPDATE users SET name = ?, email = ?, password = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, email, password, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes an existing User from the database
func Delete(id string) error {
	stmt, err := db.DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// All returns all Users from the database
func All() ([]User, error) {
	rows, err := db.DB.Query("SELECT id, name, email, password, token,created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Token, &u.Created_at, &u.Updated_at)
		if err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

// ReadUser reads a single User from the database by its ID
func Read(id string) (*User, error) {
	var result User
	err := db.DB.QueryRow("SELECT id, name, email, password,token, created_at, updated_at FROM users WHERE id = ?", id).
		Scan(&result.Id, &result.Name, &result.Email, &result.Password, &result.Token, &result.Created_at, &result.Updated_at)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//get the hash from email address

func GetHash(email string) (*User, error) {
	var result User
	err := db.DB.QueryRow("SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?", email).
		Scan(&result.Id, &result.Name, &result.Email, &result.Password, &result.Created_at, &result.Updated_at)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
