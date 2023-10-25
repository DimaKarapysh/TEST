package repository

import (
	"TEST2/app/core"
	"TEST2/domain"
	"context"
	"database/sql"
	trmsql "github.com/avito-tech/go-transaction-manager/sql"
	"github.com/pkg/errors"
)

type PersonRepos struct {
	log    core.Logger
	db     *sql.DB
	ctx    context.Context
	getter *trmsql.CtxGetter
}

func NewPersonRepos(log core.Logger, db *sql.DB, ctx context.Context, tr *trmsql.CtxGetter) *PersonRepos {
	return &PersonRepos{
		log:    log,
		db:     db,
		ctx:    ctx,
		getter: tr,
	}
}

func (p *PersonRepos) Exist(id int) (bool, error) {

	query := `select id from persons where id=$1 and deleted_at is null limit 1`
	err := p.db.QueryRow(query, id).Scan(&id)
	if err != nil {
		return false, err
	}

	switch err {
	case nil:
		return true, nil
	case sql.ErrNoRows:
		return false, nil
	default:
		return false, err
	}
}

func (p *PersonRepos) Get() ([]*domain.Person, error) {
	var persons []*domain.Person

	query := `SELECT id, name,surname ,patronymic,age,sex,nation FROM persons WHERE  deleted_at IS NULL ORDER BY id  ;`
	raws, err := p.getter.DefaultTrOrDB(p.ctx, p.db).QueryContext(p.ctx, query)

	if err != nil {
		return nil, errors.Wrap(err, "ReposGetErr")
	}

	for raws.Next() {
		var person domain.Person
		err = raws.Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Sex, &person.Nation)
		persons = append(persons, &domain.Person{
			ID:         person.ID,
			Name:       person.Name,
			Surname:    person.Surname,
			Patronymic: person.Patronymic,
			Age:        person.Age,
			Sex:        person.Sex,
			Nation:     person.Nation,
		})
	}
	return persons, nil
}

func (p *PersonRepos) GetByLimit(limit int) ([]*domain.Person, error) {
	var persons []*domain.Person

	query := `SELECT id, name,surname ,patronymic,age,sex,nation FROM persons WHERE  deleted_at IS NULL  ORDER BY id limit $1 ;`
	raws, err := p.getter.DefaultTrOrDB(p.ctx, p.db).QueryContext(p.ctx, query, limit)
	if err != nil {
		return nil, errors.Wrap(err, "ReposGetErr")
	}

	for raws.Next() {
		var person domain.Person
		err = raws.Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Sex, &person.Nation)
		persons = append(persons, &domain.Person{
			ID:         person.ID,
			Name:       person.Name,
			Surname:    person.Surname,
			Patronymic: person.Patronymic,
			Age:        person.Age,
			Sex:        person.Sex,
			Nation:     person.Nation,
		})
	}
	return persons, nil
}

func (p *PersonRepos) GetByOffset(offset int) ([]*domain.Person, error) {
	var persons []*domain.Person

	query := `SELECT id, name,surname ,patronymic,age,sex,nation FROM persons WHERE  deleted_at IS NULL  ORDER BY id offset $1 ;`

	raws, err := p.getter.DefaultTrOrDB(p.ctx, p.db).QueryContext(p.ctx, query, offset)
	if err != nil {
		return nil, errors.Wrap(err, "ReposGetErr")
	}

	for raws.Next() {
		var person domain.Person
		err = raws.Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Sex, &person.Nation)
		persons = append(persons, &domain.Person{
			ID:         person.ID,
			Name:       person.Name,
			Surname:    person.Surname,
			Patronymic: person.Patronymic,
			Age:        person.Age,
			Sex:        person.Sex,
			Nation:     person.Nation,
		})
	}
	return persons, nil
}

func (p *PersonRepos) GetByLimitOffset(limit, offset int) ([]*domain.Person, error) {
	var persons []*domain.Person

	query := `SELECT id, name,surname ,patronymic,age,sex,nation FROM persons WHERE  deleted_at IS NULL  ORDER BY id limit $1 offset $2 ;`

	raws, err := p.getter.DefaultTrOrDB(p.ctx, p.db).QueryContext(p.ctx, query, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "ReposGetErr")
	}

	for raws.Next() {
		var person domain.Person
		err = raws.Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Sex, &person.Nation)
		persons = append(persons, &domain.Person{
			ID:         person.ID,
			Name:       person.Name,
			Surname:    person.Surname,
			Patronymic: person.Patronymic,
			Age:        person.Age,
			Sex:        person.Sex,
			Nation:     person.Nation,
		})
	}
	return persons, nil
}

func (p *PersonRepos) GetById(id int) (*domain.Person, error) {
	var item = &domain.Person{
		ID: &id,
	}
	query := `SELECT 
    			name,
    			surname,
    			patronymic,
    			age,
    			sex,
    			nation,
    			created_at,
    			updated_at,
    			deleted_at
			FROM persons
			WHERE id=$1  AND deleted_at IS NULL 
			limit 1`
	err := p.getter.DefaultTrOrDB(p.ctx, p.db).QueryRowContext(p.ctx, query, id).Scan(&item.Name,
		&item.Surname,
		&item.Patronymic,
		&item.Age,
		&item.Sex,
		&item.Nation,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.DeletedAt)

	switch err {
	case nil:
		return item, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (p *PersonRepos) GetByName(name string) ([]*domain.Person, error) {
	var persons []*domain.Person

	query := `SELECT id, name,surname ,patronymic,age,sex,nation FROM persons WHERE name=$1  AND deleted_at IS NULL;`

	raws, err := p.getter.DefaultTrOrDB(p.ctx, p.db).QueryContext(p.ctx, query, name)
	if err != nil {
		return nil, errors.Wrap(err, "ReposGetByNameErr")
	}

	for raws.Next() {
		var person domain.Person
		err = raws.Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Sex, &person.Nation)
		persons = append(persons, &domain.Person{
			ID:         person.ID,
			Name:       person.Name,
			Surname:    person.Surname,
			Patronymic: person.Patronymic,
			Age:        person.Age,
			Sex:        person.Sex,
			Nation:     person.Nation,
		})
	}
	return persons, nil
}

func (p *PersonRepos) Insert(person *domain.Person) error {

	query := `INSERT INTO persons (name, surname, patronymic,age,sex,nation) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`

	_, err := p.getter.DefaultTrOrDB(p.ctx, p.db).ExecContext(p.ctx, query, person.Name, person.Surname, person.Patronymic, person.Age, person.Sex, person.Nation)
	if err != nil {
		return errors.Wrap(err, "ReposInsertPostgresErr")
	}

	return err
}

func (p *PersonRepos) Update(person *domain.UpdatePerson) error {

	query := `UPDATE persons SET name=$2,surname=$3 ,patronymic=$4,age=$5,sex=$6,nation=$7,updated_at=now() WHERE id=$1`

	_, err := p.getter.DefaultTrOrDB(p.ctx, p.db).ExecContext(p.ctx, query, person.ID, person.Name, person.Surname, person.Patronymic, person.Age, person.Sex, person.Nation)
	if err != nil {
		return errors.Wrap(err, "ReposUpdatePostgresErr")
	}

	return err

}

func (p *PersonRepos) Delete(id int) error {
	query := `UPDATE persons SET deleted_at=now() WHERE id=$1 `

	_, err := p.getter.DefaultTrOrDB(p.ctx, p.db).ExecContext(p.ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "ReposPersonErr")
	}

	return err

}
