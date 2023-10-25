package delivery

import (
	"TEST2/domain"
	"strconv"
	"time"
)

type NSPPerson struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type UpdatePerson struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Sex        string `json:"sex"`
	Nation     string `json:"nation"`
}

type PersonInfo struct {
	ID         *int   `json:"id,omitempty"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Sex        string `json:"sex"`
	Nation     string `json:"nation"`

	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt *time.Time
}
type PersonId struct {
	Id int `json:"id"`
}

func Convert(param string) int {
	p, err := strconv.Atoi(param)
	if err != nil {
		return -1
	}
	return p
}

func (n *NSPPerson) DTONSP() *domain.NSPPerson {
	return &domain.NSPPerson{
		Name:       n.Name,
		Surname:    n.Surname,
		Patronymic: n.Patronymic,
	}
}

func (u *UpdatePerson) DTOUpdatePerson() *domain.UpdatePerson {
	return &domain.UpdatePerson{
		ID:         u.ID,
		Name:       u.Name,
		Surname:    u.Surname,
		Patronymic: u.Patronymic,
		Age:        u.Age,
		Sex:        u.Sex,
		Nation:     u.Nation,
	}
}

func DTOPersonInfo(person *domain.Person) *PersonInfo {
	return &PersonInfo{
		ID:         person.ID,
		Name:       person.Name,
		Surname:    person.Surname,
		Patronymic: person.Patronymic,
		Age:        person.Age,
		Sex:        person.Sex,
		Nation:     person.Nation,
		UpdatedAt:  person.UpdatedAt,
		CreatedAt:  person.CreatedAt,
		DeletedAt:  person.DeletedAt,
	}
}

func DTOPersonsInfo(persons []*domain.Person) []*PersonInfo {
	var taskInfos []*PersonInfo
	for _, t := range persons {
		taskInfos = append(taskInfos, DTOPersonInfo(t))
	}
	return taskInfos
}
