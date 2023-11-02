package model

type Todo struct {
	Id        uint64 `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func CreateTodo(title string) error {
	statement := `insert into todos (title, completed) values ($1, $2);`

	_, err := db.Query(statement, title, false)

	return err
}

func GetAllTodos() ([]Todo, error) {
	todos := []Todo{}

	statement := `select id, title, completed from todos order by id;`

	rows, err := db.Query(statement)

	if err != nil {
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var title string
		var completed bool

		err := rows.Scan(&id, &title, &completed)

		if err != nil {
			return todos, err
		}

		todo := Todo{
			Id:        id,
			Title:     title,
			Completed: completed,
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func GetTodo(id uint64) (Todo, error) {
	statement := `select title, completed from todos where id = $1;`

	todo := Todo{}
	todo.Id = id

	row, err := db.Query(statement, id)

	if err != nil {
		return todo, err
	}

	defer row.Close()

	for row.Next() {
		var title string
		var completed bool

		err := row.Scan(&title, &completed)

		if err != nil {
			return todo, err
		}

		todo.Title = title
		todo.Completed = completed
	}

	return todo, nil
}

func MarkDone(id uint64) error {
	todo, err := GetTodo(id)

	if err != nil {
		return err
	}

	statement := `update todos set completed = $2 where id = $1;`

	_, err = db.Query(statement, id, !todo.Completed)

	return err
}

func Delete(id uint64) error {
	statement := `delete from todos where id = $1;`

	_, err := db.Query(statement, id)

	return err
}
