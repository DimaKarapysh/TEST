package agify

import (
	"TEST2/app/core"
	"TEST2/domain"
)

type Agify struct {
	log core.Logger
}

func NewAgify(log core.Logger) *Agify {
	return &Agify{log: log}
}

func (a *Agify) Enrich(nsp *domain.NSPPerson) (*domain.Person, error) {
	age, err := GetFieldAge(nsp.Name)
	if err != nil {
		return nil, err
	}

	gen, err := GetFieldGender(nsp.Name)
	if err != nil {
		return nil, err
	}

	n := Country{
		Id:          "",
		Probability: 0,
	}

	Nations, err := GetFieldNation(nsp.Name)
	if err != nil {
		return nil, err
	}

	for _, country := range Nations.Nation {
		if country.Probability > n.Probability {
			n = country
		}
	}

	person := &domain.Person{
		Name:       nsp.Name,
		Surname:    nsp.Surname,
		Patronymic: nsp.Patronymic,
		Age:        age.Age,
		Sex:        gen.Gender,
		Nation:     n.Id,
	}

	return person, nil
}
