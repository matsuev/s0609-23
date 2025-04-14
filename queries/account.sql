-- name: UserLogin :one
SELECT * FROM public.account
WHERE "username"=$1 AND "password"=$2
;

-- name: GetUserByID :one
SELECT * FROM public.account
WHERE id=$1
;