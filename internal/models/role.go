package models

import (
	"database/sql"
	"errors"
)

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var ErrNoRole = errors.New("no role found")

func GetRoles(db *sql.DB) ([]Role, error) {
	rows, err := db.Query("SELECT id, name FROM roles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []Role{}
	for rows.Next() {
		var r Role
		err := rows.Scan(&r.ID, &r.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}
	return roles, nil
}

func CreateRole(db *sql.DB, r *Role) error {
	err := db.QueryRow("INSERT INTO roles (name) VALUES ($1) RETURNING id", r.Name).Scan(&r.ID)
	return err
}

func GetRole(db *sql.DB, id int) (*Role, error) {
	var r Role
	err := db.QueryRow("SELECT id, name FROM roles WHERE id=$1", id).Scan(&r.ID, &r.Name)
	if err == sql.ErrNoRows {
		return nil, ErrNoRole
	} else if err != nil {
		return nil, err
	}
	return &r, nil
}

func UpdateRole(db *sql.DB, r *Role) error {
	_, err := db.Exec("UPDATE roles SET name=$1 WHERE id=$2", r.Name, r.ID)
	return err
}

func DeleteRole(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM roles WHERE id=$1", id)
	return err
}
