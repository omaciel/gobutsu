-- name: CreateTestCase :one
INSERT INTO testcases (
  classname,
  filename,
  linenumber,
  testcasename,
  duration
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetTestCase :one
SELECT * FROM testcases
WHERE testcaseid = $1 LIMIT 1;

-- name: ListTestCases :many
SELECT * FROM testcases
ORDER BY testcaseid
LIMIT $1
OFFSET $2;

-- name: DeleteTestCase :exec
DELETE FROM testcases
WHERE testcaseid = $1;