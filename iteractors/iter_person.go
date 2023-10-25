package iteractors

import (
	"TEST2/app/core"
	"TEST2/domain"
	"context"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/pkg/errors"
	"strconv"
)

type PersonIter struct {
	log   core.Logger
	repos domain.Repos
	agify domain.Agify
	ctx   context.Context
	trm   *manager.Manager
}

func NewIterPerson(log core.Logger, agify domain.Agify, repos domain.Repos, ctx context.Context, maneger *manager.Manager) *PersonIter {
	return &PersonIter{
		log:   log,
		agify: agify,
		repos: repos,
		ctx:   ctx,
		trm:   maneger,
	}
}

func (i *PersonIter) GetALLPerson() ([]*domain.Person, error) {

	persons, err := i.repos.Get()
	if err != nil {
		return nil, errors.Wrap(err, "IterGetErr")
	}
	return persons, nil
}

func (i *PersonIter) GetById(id int) (*domain.Person, error) {

	ok, err := i.repos.Exist(id)
	if err != nil {
		return nil, errors.Wrap(err, "Not Exist")
	}

	if ok {
		person, err := i.repos.GetById(id)
		if err != nil {
			return nil, errors.Wrap(err, "IterGetByIdErr")
		}
		return person, nil
	}

	return nil, err
}

func (i *PersonIter) GetByName(name string) ([]*domain.Person, error) {
	persons, err := i.repos.GetByName(name)
	if err != nil {
		return nil, errors.Wrap(err, "IterGetByNameErr")
	}

	if len(persons) == 0 {
		return nil, errors.Wrap(errors.New("no Result"), "no result")
	}

	return persons, err
}

func (i *PersonIter) GetByLimit(limit int) ([]*domain.Person, error) {

	persons, err := i.repos.GetByLimit(limit)
	if err != nil {
		return nil, errors.Wrap(err, "IterGetByLimitErr")
	}

	return persons, nil
}

func (i *PersonIter) GetByOffset(offset int) ([]*domain.Person, error) {

	persons, err := i.repos.GetByOffset(offset)
	if err != nil {
		return nil, errors.Wrap(err, "IterGetByLimitErr")
	}

	return persons, nil
}

func (i *PersonIter) GetByLimitOffset(limit, offset int) ([]*domain.Person, error) {

	persons, err := i.repos.GetByLimitOffset(limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "IterGetByLimitErr")
	}

	return persons, nil
}

func (i *PersonIter) InsertPerson(person *domain.NSPPerson) error {
	if len(person.Name) == 0 {
		return errors.New("Name is empty")
	}

	p, err := i.agify.Enrich(person)
	if err != nil {
		return errors.Wrap(err, "IterInsertErrAgify")
	}

	err = i.trm.Do(i.ctx, func(ctx context.Context) error {
		i.log.Info("insert person Into DB")
		return i.trm.Do(ctx, func(ctx context.Context) error {
			return i.repos.Insert(p)
		})
	})
	if err != nil {
		return errors.Wrap(err, "IterInsertErr")
	}
	return nil
}

func (i *PersonIter) UpdatePerson(person *domain.UpdatePerson) error {
	ok, err := i.repos.Exist(person.ID)
	if err != nil {
		return errors.Wrap(err, "Not Exist")
	}

	if ok {
		err = i.trm.Do(i.ctx, func(ctx context.Context) error {
			i.log.Info("update person number is " + strconv.Itoa(person.ID))
			return i.trm.Do(ctx, func(ctx context.Context) error {
				return i.repos.Update(person)
			})
		})

		if err != nil {
			return errors.Wrap(err, "IterPersonErr")
		}

	}

	return err
}

func (i *PersonIter) DeletePerson(id int) error {

	ok, err := i.repos.Exist(id)
	if err != nil {
		return errors.Wrap(err, "Not Exist")
	}

	if ok {
		err = i.trm.Do(i.ctx, func(ctx context.Context) error {
			i.log.Info("delete person where id " + strconv.Itoa(id))
			return i.trm.Do(ctx, func(ctx context.Context) error {
				return i.repos.Delete(id)
			})
		})
		if err != nil {
			return errors.Wrap(err, "IterDeleteErr")
		}
	}

	return err
}
