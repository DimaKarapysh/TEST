package domain

type Iter interface {
	GetALLPerson() ([]*Person, error)
	GetByLimit(int) ([]*Person, error)
	GetById(int) (*Person, error)
	GetByName(string) ([]*Person, error)
	GetByOffset(int) ([]*Person, error)
	GetByLimitOffset(int, int) ([]*Person, error)

	InsertPerson(person *NSPPerson) error
	UpdatePerson(person *UpdatePerson) error
	DeletePerson(id int) error
}

type Repos interface {
	Exist(int) (bool, error)

	Get() ([]*Person, error)
	GetByLimit(int) ([]*Person, error)
	GetByOffset(int) ([]*Person, error)
	GetByLimitOffset(int, int) ([]*Person, error)
	GetById(int) (*Person, error)
	GetByName(string) ([]*Person, error)

	Insert(*Person) error
	Update(*UpdatePerson) error
	Delete(int) error
}

type Agify interface {
	Enrich(*NSPPerson) (*Person, error)
}
