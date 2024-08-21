package mehtods

import (
	sqlbuilder "conn/internal/database/sql"
	"conn/internal/models"
	"database/sql"
	"log"
)

type Database struct {
	Db *sql.DB
}
func (u *Database) Create(req *models.Create) (int, error) {
	query, args, err := sqlbuilder.Create(req)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	var id int
	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println(err)
		return 0, err
	}
	return id, nil
}

func (u *Database) GetbyId(req int) (*models.Update, error) {
	query, args, err := sqlbuilder.GetbyId(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var res models.Update
	if err := u.Db.QueryRow(query, args...).Scan(&res.ID, &res.Title, &res.Author, &res.Published_Year); err != nil {
		log.Println(err)
		return nil, err
	}
	return &res, nil
}

func (u *Database) Get() ([]*models.Update, error) {
	query, args, err := sqlbuilder.Get()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows, err := u.Db.Query(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var res []*models.Update
	for rows.Next() {
		var all models.Update
		if err := rows.Scan(&all.ID, &all.Title, &all.Author, &all.Published_Year); err != nil {
			log.Println(err)
			return nil, err
		}
		res = append(res, &all)
	}
	return res, nil
}

func (u *Database) Update(req *models.Update) (int, error) {
	query, args, err := sqlbuilder.Update(req)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	var id int
	if err := u.Db.QueryRow(query, args...).Scan(&id); err != nil {
		log.Println(err)
		return 0, err
	}
	return id, nil

}

func (u *Database) Delete(req int) error {
	query, args, err := sqlbuilder.Delete(req)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
