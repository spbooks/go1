package main

import "database/sql"

const pageSize = 25

var globalImageStore ImageStore

type ImageStore interface {
	Save(image *Image) error
	Find(id string) (*Image, error)
	FindAll(offset int) ([]Image, error)
	FindAllByUser(user *User, offset int) ([]Image, error)
}

type DBImageStore struct {
	db *sql.DB
}

func NewDBImageStore() ImageStore {
	return &DBImageStore{
		db: globalMySQLDB,
	}
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

func (store *DBImageStore) FindAllByUser(user *User, offset int) ([]Image, error) {
	rows, err := store.db.Query(
		`
		SELECT id, user_id, name, location, description, size, created_at
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
