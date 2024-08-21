package sqlbuilder

import (
	"conn/internal/models"
	"log"

	"github.com/Masterminds/squirrel"
)

func Create(req *models.Create) (string, []interface{}, error) {
	query, args, err := squirrel.Insert("books").
		Columns("title", "author", "published_year").
		Values(req.Title, req.Author, req.Published_Year).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func GetbyId(req int) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("books").
		Where(squirrel.Eq{"id": req}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func Get() (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").
		From("books").
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func Update(req *models.Update) (string, []interface{}, error) {
	query, args, err := squirrel.Update("books").
		SetMap(map[string]interface{}{
			"title":          req.Title,
			"author":         req.Author,
			"published_year": req.Published_Year,
		}).Where(squirrel.Eq{"id": req.ID}).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}

func Delete(req int) (string, []interface{}, error) {
	query, args, err := squirrel.Delete("books").Where(squirrel.Eq{"id": req}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	return query, args, nil
}
