package models

import (
	"database/sql"
	"errors"
)

type Server struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IPAddress string `json:"ip_address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var ErrNoServer = errors.New("no server found")

func GetServers(db *sql.DB) ([]Server, error) {
	rows, err := db.Query("SELECT id, name, ip_address, username, password FROM servers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	servers := []Server{}
	for rows.Next() {
		var s Server
		err := rows.Scan(&s.ID, &s.Name, &s.IPAddress, &s.Username, &s.Password)
		if err != nil {
			return nil, err
		}
		servers = append(servers, s)
	}
	return servers, nil
}

func CreateServer(db *sql.DB, s *Server) error {
	err := db.QueryRow("INSERT INTO servers (name, ip_address, username, password) VALUES ($1, $2, $3, $4) RETURNING id",
		s.Name, s.IPAddress, s.Username, s.Password).Scan(&s.ID)
	return err
}

func GetServer(db *sql.DB, id int) (*Server, error) {
	var s Server
	err := db.QueryRow("SELECT id, name, ip_address, username, password FROM servers WHERE id=$1", id).Scan(&s.ID, &s.Name, &s.IPAddress, &s.Username, &s.Password)
	if err == sql.ErrNoRows {
		return nil, ErrNoServer
	} else if err != nil {
		return nil, err
	}
	return &s, nil
}

func UpdateServer(db *sql.DB, s *Server) error {
	_, err := db.Exec("UPDATE servers SET name=$1, ip_address=$2, username=$3, password=$4 WHERE id=$5",
		s.Name, s.IPAddress, s.Username, s.Password, s.ID)
	return err
}

func DeleteServer(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM servers WHERE id=$1", id)
	return err
}

