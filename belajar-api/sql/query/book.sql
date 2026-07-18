-- name: CreateBook :one
INSERT INTO books (title, author, isbn)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetBook :one
SELECT * FROM books
WHERE id = $1;

-- name: ListBooks :many
SELECT * FROM books
WHERE is_available = TRUE
ORDER BY created_at DESC;

-- name: UpdateBookAvailability :one
UPDATE books 
SET 
  is_available = $2, 
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;