package repositories

import (
	"database/sql"

	"github.com/jabernardo/tugon/core"
)

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoRepository struct {
	Db *sql.DB
}

func SetupTodoRepository(db *sql.DB) {
	tx, err := db.Begin()

	if err != nil {
		core.GetLoggerInstance().Error("[repositories.todo]", "err", err)
	}

	_, err = tx.Exec(`
    CREATE TABLE IF NOT EXISTS todo (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      title TEXT,
      description TEXT
    )
  `)

	if err != nil {
		tx.Rollback()
		core.GetLoggerInstance().Error("[repositories.todo]", "err", err)
	}

	err = tx.Commit()

	if err != nil {
		core.GetLoggerInstance().Error("[repositories.todo]", "err", err)
	}
}

func NewTodoRepository() *TodoRepository {
	db := core.GetDBInstance()

	return &TodoRepository{Db: db}
}

func (repo *TodoRepository) Create(item *Todo) (sql.Result, error) {
	tx, _ := repo.Db.Begin()

	res, err := tx.Exec("INSERT INTO todo(title, description) VALUES (?, ?)", item.Title, item.Description)

	if err != nil {
		core.GetLoggerInstance().Error("[repositories.todo]", "err", err)
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		core.GetLoggerInstance().Error("[repositories.todo]", "err", err)
		return nil, err
	}

	return res, nil
}

func (repo *TodoRepository) Get(id int) *Todo {
	results := Todo{}

	err := repo.Db.QueryRow("SELECT * FROM todo WHERE id = ?", id).Scan(&results.Id, &results.Title, &results.Description)

	if err != nil {
		core.GetLoggerInstance().Error("[repositories.todo]", "err", err)
		return nil
	}

	return &results
}

func (repo *TodoRepository) Delete(id int) (sql.Result, error) {
	tx, err := repo.Db.Begin()

	if err != nil {
		core.GetLoggerInstance().Error("[repositories.todo]", "err", err)
		return nil, err
	}

	res, err := tx.Exec("DELETE FROM todo WHERE id = ?", id)

	if err != nil {
		core.GetLoggerInstance().Error("[repositories.todo]", "err", err)
		return nil, err
	}

	_ = tx.Commit()

	return res, err
}

func (repo *TodoRepository) Update(id int, title string, description string) (sql.Result, error) {
	tx, err := repo.Db.Begin()

	if err != nil {
		core.GetLoggerInstance().Error("[repositories.todo]", "err", err)
		return nil, err
	}

	res, err := tx.Exec("UPDATE todo SET title = ?, description = ? WHERE id = ?", title, description, id)

	if err != nil {
		core.GetLoggerInstance().Error("[repositories.todo]", "err", err)
		return nil, err
	}

	_ = tx.Commit()

	return res, err
}

func (repo *TodoRepository) GetAll() []Todo {
	var results []Todo
	rows, err := repo.Db.Query("SELECT * FROM todo LIMIT 100")

	if err != nil {
		core.GetLoggerInstance().Error("[repositories.todo]", "err", err)
		return results
	}

	for rows.Next() {
		todo := Todo{}

		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description); err != nil {
			core.GetLoggerInstance().Error("[repositories.todo]", "err", err)
		}

		results = append(results, todo)
	}

	return results
}
