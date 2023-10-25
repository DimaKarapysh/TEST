package delivery

import (
	"TEST2/app/core"
	"TEST2/app/rest"
	"TEST2/domain"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"net/http"
)

type PersonDeliveryService struct {
	log  core.Logger
	iter domain.Iter
	v    validator.Validate
}

func NewPersonDeliveryService(log core.Logger, i domain.Iter, v validator.Validate) *PersonDeliveryService {
	return &PersonDeliveryService{
		log:  log,
		iter: i,
		v:    v,
	}
}

func (t *PersonDeliveryService) GetPersons(w http.ResponseWriter, r *http.Request) {
	Persons, err := t.iter.GetALLPerson()
	if err != nil {
		rest.ServerError(w, err)
		return
	}
	rest.ServerSuccessStruct(w, Persons)
}

func (t *PersonDeliveryService) GetPersonsByLimit(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")

	Limit := Convert(limit)
	if Limit == -1 {
		rest.ValidationError(w, "Can not parse limit")
		return
	}

	Persons, err := t.iter.GetByLimit(Limit)
	if err != nil {
		rest.ServerError(w, err)
		return
	}
	rest.ServerSuccessStruct(w, Persons)
}

func (t *PersonDeliveryService) GetPersonsByOffset(w http.ResponseWriter, r *http.Request) {
	offset := r.URL.Query().Get("offset")

	Offset := Convert(offset)
	if Offset == -1 {
		rest.ValidationError(w, "Can not parse offset")
		return
	}

	Persons, err := t.iter.GetByOffset(Offset)
	if err != nil {
		rest.ServerError(w, err)
		return
	}

	rest.ServerSuccessStruct(w, Persons)
}

func (t *PersonDeliveryService) GetPersonsByLimitOffset(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")
	Limit := Convert(limit)
	if Limit == -1 {
		rest.ValidationError(w, "Can not parse limit")
		return
	}

	Offset := Convert(offset)
	if Offset == -1 {
		rest.ValidationError(w, "Can not parse offset")
		return
	}

	Persons, err := t.iter.GetByLimitOffset(Limit, Offset)
	if err != nil {
		rest.ServerError(w, err)
		return
	}

	rest.ServerSuccessStruct(w, Persons)
}

func (t *PersonDeliveryService) GetPersonsById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	Id := Convert(id)
	if Id == -1 {
		rest.ValidationError(w, "Can not parse Id")
		return
	}

	Person, err := t.iter.GetById(Id)
	if err != nil {
		rest.ServerError(w, err)
		return
	}

	rest.ServerSuccessStruct(w, DTOPersonInfo(Person))
}

func (t *PersonDeliveryService) GetPersonsByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	Persons, err := t.iter.GetByName(name)
	if err != nil {
		rest.ServerError(w, err)
		return
	}

	rest.ServerSuccessStruct(w, DTOPersonsInfo(Persons))
}

func (t *PersonDeliveryService) AddPerson(w http.ResponseWriter, r *http.Request) {
	var person NSPPerson
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		rest.ValidationError(w, "Cannot parse json")
		return
	}

	err = t.v.Struct(person)
	if err != nil {
		rest.ValidationError(w, err.Error())
		return
	}

	err = t.iter.InsertPerson(person.DTONSP())
	if err != nil {
		rest.ServerError(w, errors.Wrap(err, "AddTask"))
		return
	}

	rest.ServerSuccessStruct(w, person)
}

func (t *PersonDeliveryService) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var person UpdatePerson
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		rest.ValidationError(w, "Cannot parse json")
		return
	}

	err = t.v.Struct(person)
	if err != nil {
		rest.ValidationError(w, err.Error())
		return
	}

	err = t.iter.UpdatePerson(person.DTOUpdatePerson())
	if err != nil {
		rest.ServerError(w, errors.Wrap(err, "AddTask"))
		return
	}

	rest.ServerSuccessStruct(w, person)
}

func (t *PersonDeliveryService) DeletePerson(w http.ResponseWriter, r *http.Request) {
	var Id PersonId
	err := json.NewDecoder(r.Body).Decode(&Id)
	if err != nil {
		rest.ValidationError(w, "Cannot parse json")
		return
	}

	err = t.v.Struct(Id)
	if err != nil {
		rest.ValidationError(w, err.Error())
		return
	}

	err = t.iter.DeletePerson(Id.Id)
	if err != nil {
		rest.ServerError(w, errors.Wrap(err, "AddTask"))
		return
	}

	rest.ServerSuccessStruct(w, Id)
}
