package queries

const GetUserData = "SELECT username FROM users WHERE id = :id"

const AddUserData = `
	INSERT INTO
		users
	VALUES (:id, ':nickname')
`
