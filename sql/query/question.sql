-- name: GetQuestionById :one
SELECT * FROM `question`
WHERE `id` = ? AND `isDelete` = 0 LIMIT 1;

-- name: PageListQuestions :many
SELECT * FROM `question`
WHERE `isDelete` = 0
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: ListQuestions :many
SELECT * FROM `question`
WHERE `isDelete` = 0
ORDER BY id;

-- name: AddQuestion :execresult
insert into `question` (
    `title`, `content`, `tags`, `answer`, `judgeCase`, `judgeConfig`, `userId`
    ) values (
        ?, ?, ?, ?, ?, ?, ?
    );

-- name: UpdateQuestion :exec
UPDATE `question` set title=?, content=?, tags=?, answer=?, `judgeCase`=?, `judgeConfig`=?
WHERE id = ?;

-- name: DeleteQuestion :exec
UPDATE `question` set `isDelete` = 1 
WHERE id = ?;