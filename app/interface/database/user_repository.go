package database

import "github.com/hiromisuzuki/clean-arch-example/app/model"

//UserRepository UserRepository
type UserRepository struct {
	SQLHandler
}

//Store Store
func (repo *UserRepository) Store(u model.User) (id int, err error) {
	res, err := repo.Execute(`
			INSERT INTO 
				users (mail_address) 
			VALUES (?)
		`, u.MailAddress)
	if err != nil {
		return
	}

	id64, err := res.LastInsertID()
	if err != nil {
		return
	}

	id = int(id64)
	return
}

//FindByID FindByID
func (repo *UserRepository) FindByID(id int) (user model.User, err error) {

	row, err := repo.Query(`
			SELECT
				id, mail_address, created_at, updated_at
			FROM
				users
			WHERE
				id = ?
		`, id)
	defer row.Close()
	if err != nil {
		return
	}

	if row.Next() {
		err = row.Scan(&user.ID, &user.MailAddress, &user.CreatedAt, &user.UpdatedAt)
	}

	return
}

//FindAll FindAll
func (repo *UserRepository) FindAll(p model.Pagination) (users model.Users, err error) {
	rows, err := repo.Query(`
		SELECT 
			id, mail_address, created_at, updated_at
		FROM
			users
		LIMIT
			` + p.Limit() + `
		OFFSET
			` + p.Offset())
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.ID, &user.MailAddress, &user.CreatedAt, &user.UpdatedAt)
		users.User = append(users.User, &user)
	}
	return
}
