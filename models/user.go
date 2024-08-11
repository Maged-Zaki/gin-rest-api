package models

import (
	"database/sql"
	"fmt"

	"github.com/Maged-Zaki/gin-rest-api/db"
	"github.com/Maged-Zaki/gin-rest-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?) RETURNING id"

	// Execute query and scan the returned id
	err := db.DB.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}
func (user *User) Delete() error {
	query := "DELETE FROM users WHERE id=?"

	// Execute query and scan the returned id
	result, err := db.DB.Exec(query, user.ID)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("no user found with the specified ID")
	}

	return nil
}

// func GetUserByEmail(email string) (User, error) {
// 	query := "SELECT * FROM users WHERE email = ?"

// 	var u User

// 	err := db.DB.QueryRow(query, email).Scan(&u.ID, &u.Email, &u.Password)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return u, fmt.Errorf("User with email %s does not exist", u.Email)
// 		}
// 		return u, fmt.Errorf("ERROR with Row: %s", err.Error())
// 	}

//		return u, nil
//	}
func (user *User) ValidateCredentials() error {
	query := "SELECT id, email, password FROM users WHERE email = ?"

	var dbUser User

	err := db.DB.QueryRow(query, user.Email).Scan(&dbUser.ID, &dbUser.Email, &dbUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// random hash for non-existent user to avoid timing attacks
			dbUser.Password = "$2a$14$xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		} else {
			return err
		}
	} else {
		// User exists, assign values to the receiver
		user.ID = dbUser.ID
		user.Email = dbUser.Email
	}

	// Compare password
	isCorrectPassword := utils.CheckPasswordHash(dbUser.Password, user.Password)
	if !isCorrectPassword {
		return fmt.Errorf("incorrect email or password")
	}

	return nil
}
