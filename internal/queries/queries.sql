-- name: ListTodos :many
SELECT * FROM todos;

-- name: CreateTodo :exec
INSERT INTO todos (
  name
) VALUES (
  $1
);

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;