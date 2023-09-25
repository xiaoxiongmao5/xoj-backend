-- name: GetUserInfoById :one
SELECT * FROM `user`
WHERE `id` = ? AND `isDelete` = 0 LIMIT 1;

-- name: GetUserInfoByUniUserAccount :one
SELECT * FROM `user`
WHERE `userAccount` = ? AND `isDelete` = 0 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM `user`
WHERE `isDelete` = 0
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CreateUser :execresult
insert into `user` (
    `userAccount`, `userPassword`
    ) values (
        ?, ?
    );

-- name: DeleteUser :exec
UPDATE `user` set `isDelete` = 1 
WHERE id = ?;