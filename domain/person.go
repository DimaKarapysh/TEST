package domain

import "time"

type Person struct {
	ID         *int
	Name       string
	Surname    string
	Patronymic string
	Age        int
	Sex        string
	Nation     string

	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt *time.Time
}

type UpdatePerson struct {
	ID         int
	Name       string
	Surname    string
	Patronymic string
	Age        int
	Sex        string
	Nation     string
}

type NSPPerson struct {
	Name       string
	Surname    string
	Patronymic string
}
