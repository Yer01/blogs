package model

import (
	"database/sql"
	"time"
)

type Blog struct {
	ID      int
	Name    string
	Content string
	Created time.Time
	Updated time.Time
}

type BlogModel struct {
	DB *sql.DB
}

func (m *BlogModel) Insert(name string, content string) (int, error) {
	var id int
	quer := `INSERT INTO blogs (name, content, created_at, updated_at)
	VALUES ($1, $2, now(), now())
	RETURNING blog_id`

	res := m.DB.QueryRow(quer, name, content)

	err := res.Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (m *BlogModel) Get(id int) (Blog, error) {
	quer := `SELECT * FROM blogs WHERE blog_id = $1`

	res := m.DB.QueryRow(quer, id)

	var blog Blog

	err := res.Scan(&blog.ID, &blog.Name, &blog.Content, &blog.Created, &blog.Updated)
	if err != nil {
		return blog, err
	}
	return blog, nil
}

func (m *BlogModel) GetAll() ([]Blog, error) {
	quer := `SELECT * FROM blogs`

	res, err := m.DB.Query(quer)

	if err != nil {
		return nil, err
	}

	var blogs []Blog

	for res.Next() {
		var blog Blog

		err = res.Scan(&blog.ID, &blog.Name, &blog.Content, &blog.Created, &blog.Updated)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	if err = res.Err(); err != nil {
		return nil, err
	}
	return blogs, nil
}

func (m *BlogModel) Update(id int, content string) (int, error) {
	quer := `UPDATE blogs SET content = $1 WHERE blog_id = $2
             RETURNING blog_id`

	var updatedID int
	err := m.DB.QueryRow(quer, content, id).Scan(&updatedID)
	if err != nil {
		return -1, err
	}

	return updatedID, nil
}

func (m *BlogModel) Delete(id int) error {
	quer := "DELETE FROM blogs WHERE blog_id = $1"

	res, err := m.DB.Exec(quer, id)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
