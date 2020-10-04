package shorten

import (
	"database/sql"
)

type repository struct {
	db *sql.DB
}

const (
	createTable = "CREATE TABLE IF NOT EXISTS shorten_urls (id TEXT PRIMARY KEY NOT NULL, shorten_url TEXT)"
)

// NewRepository ...
func NewRepository(dialect, dsn string, idleConn, maxConn int) (Repository, error) {

	db, err := sql.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	db.Exec(createTable)

	return &repository{db}, nil
}

func (r repository) Close() {
	r.db.Close()
}

func (r repository) FindByID(id string) (*ShortURLModel, error) {
	var (
		model ShortURLModel
	)

	err := r.db.QueryRow(`SELECT id, shorten_url FROM shorten_urls WHERE id = ?`, id).Scan(&model.ID, &model.Location)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (r repository) Find() ([]*ShortURLModel, error) {
	return nil, nil
}

func (r repository) Create(model *ShortURLModel) error {
	_, err := r.db.Exec("INSERT INTO shorten_urls (id, shorten_url) VALUES(?, ?)", model.ID, model.Location)
	return err
}

func (r repository) Update(model *ShortURLModel) error {
	return nil
}

func (r repository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM shorten_urls WHERE id = ?", id)
	return err
}
