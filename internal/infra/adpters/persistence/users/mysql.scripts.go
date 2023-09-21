package users

const (
	SqlUserExistsById = `SELECT count(1) 
    FROM users WHERE date_deleted IS NULL AND code = ?;`

	SqlUserCreate = `INSERT INTO users 
    (code,name)
	VALUES (?,?);`

	SqlUserDeleteById = `DELETE FROM users 
    WHERE code = ?;`

	SqlUserDeletedLogicalById = `UPDATE users 
    SET date_deleted = now() WHERE code = ?;`
)

const (
	CollectionName = "users"
)
