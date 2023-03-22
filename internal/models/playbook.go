package models

import (
	"database/sql"
	"errors"
)

type Playbook struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

var ErrNoPlaybook = errors.New("no playbook found")

func (db *DB) GetPlaybook(id int) (Playbook, error) {
	var p Playbook
	err := db.QueryRow("SELECT id, name, content FROM playbooks WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, ErrNoPlaybook
		}
		return p, err
	}
	return p, nil
}

func (db *DB) CreatePlaybook(p *Playbook) error {
	err := db.QueryRow("INSERT INTO playbooks(name, content) VALUES($1, $2) RETURNING id", p.Name, p.Content).Scan(&p.ID)
	return err
}

func (db *DB) UpdatePlaybook(p *Playbook) error {
	_, err := db.Exec("UPDATE playbooks SET name=$1, content=$2 WHERE id=$3", p.Name, p.Content, p.ID)
	return err
}

func (db *DB) DeletePlaybook(id int) error {
	_, err := db.Exec("DELETE FROM playbooks WHERE id=$1", id)
	return err
}

