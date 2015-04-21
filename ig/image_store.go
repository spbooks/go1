package main

import (
	"database/sql"
)

const (
	pageSize = 25
)

type ImageStore interface {
	Find(id string) (*Image, error)
	FindAllByUser(user *User, offset int) ([]Image, error)
	FindAll(offset int) ([]Image, error)
	Save(image *Image) error
}

var globalImageStore ImageStore

func setupImageStore() {
	globalImageStore = NewDBImageStore()
}

type DBImageStore struct {
	db *sql.DB
}

func NewDBImageStore() ImageStore {
	return &DBImageStore{
		db: globalMySQLDB,
	}
}

/*
func NewDBImageStore(address, username, password, database string) (ImageStore, error) {
	if address == "" {
		address = "127.0.0.1:3306"
	}

	dsn := username
	if password != "" {
		dsn += ":" + password
	}

	dsn += "@"
	if address != "" {
		dsn += "tcp(" + address + ")"
	}
	dsn += "/?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("USE " + database)

	return &DBImageStore{
		DB: db,
	}, err
}
*/

/*
CREATE TABLE `images` (
  `id` varchar(255) NOT NULL DEFAULT '',
  `user_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL DEFAULT '',
  `location` varchar(255) NOT NULL DEFAULT '',
  `description` text NOT NULL,
  `size` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id_idx` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
func (store *DBImageStore) FindAllByUser(user *User, offset int) ([]Image, error) {
	rows, err := store.db.Query(
		`
		SELECT id, user_id, name, location, description, created_at
		FROM images
		WHERE user_id = ?
		ORDER BY created_at DESC
		LIMIT ?
		OFFSET ?`,
		user.ID,
		pageSize,
		offset,
	)
	if err != nil {
		return nil, err
	}

	images := []Image{}
	for rows.Next() {
		image := Image{}
		err := rows.Scan(
			&image.ID,
			&image.UserID,
			&image.Name,
			&image.Location,
			&image.Description,
			&image.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}

func (store *DBImageStore) Find(id string) (*Image, error) {
	row := store.db.QueryRow(
		`
		SELECT id, user_id, name, location, description, size, created_at
		FROM images
		WHERE id = ?`,
		id,
	)

	image := Image{}
	err := row.Scan(
		&image.ID,
		&image.UserID,
		&image.Name,
		&image.Location,
		&image.Description,
		&image.Size,
		&image.CreatedAt,
	)
	return &image, err
}

func (store *DBImageStore) FindAll(offset int) ([]Image, error) {
	rows, err := store.db.Query(
		`
		SELECT id, user_id, name, location, description, size, created_at
		FROM images
		ORDER BY created_at DESC
		LIMIT ?
		OFFSET ?
		`,
		pageSize,
		offset,
	)
	if err != nil {
		return nil, err
	}

	images := []Image{}
	for rows.Next() {
		image := Image{}
		err := rows.Scan(
			&image.ID,
			&image.UserID,
			&image.Name,
			&image.Location,
			&image.Description,
			&image.Size,
			&image.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}

func (store *DBImageStore) Save(image *Image) error {
	_, err := store.db.Exec(
		`
		REPLACE INTO images
			(id, user_id, name, location, description, size, created_at)
		VALUES
			(?, ?, ?, ?, ?, ?, ?)
		`,
		image.ID,
		image.UserID,
		image.Name,
		image.Location,
		image.Description,
		image.Size,
		image.CreatedAt,
	)
	return err
}
