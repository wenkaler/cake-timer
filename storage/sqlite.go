package storage

import (
	"fmt"

	"github.com/go-kit/kit/log/level"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wenkaler/cake-timer/xtimer"
)

type Storage struct {
	db     *sqlx.DB
	logger log.Logger
}

func New(pathDB string, logger log.Logger) (*Storage, error) {
	if pathDB == "" {
		return nil, fmt.Errorf("pathDB was empty")
	}
	db, err := sqlx.Open("sqlite3", pathDB)
	if err != nil {
		return nil, err
	}
	s := &Storage{
		db:     db,
		logger: logger,
	}
	err = s.init()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Storage) init() error {
	_, err := s.db.Unsafe().Exec(`CREATE TABLE  IF NOT EXISTS users(
									id INTEGER PRIMARY KEY AUTOINCREMENT,
									user_name VARCHAR(225) NOT NULL,
									date time NOT NULL
						)`)
	if err != nil {
		return fmt.Errorf("failed create records table: %v", err)
	}
	level.Info(s.logger).Log("msg", "create data base, with table.")
	return nil
}

func (s *Storage) NewChat(cid int64) error {
	_, err := s.db.Unsafe().Exec(`INSERT INTO chats(id) VALUES(?)`, cid)
	return err
}

func (s *Storage) GetUsers() (users []xtimer.User, err error) {
	err = s.db.Unsafe().Select(&users, `select * from users`)
	return
}

func (s *Storage) Close() error {
	return s.db.Close()
}
