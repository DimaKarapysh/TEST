package agify

import (
	"encoding/json"
	"net/http"
)

func GetFieldAge(name string) (age *AgilyAge, err error) {
	r, err := http.Get("https://api.agify.io/?name=" + name)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(r.Body).Decode(&age)
	if err != nil {
		return nil, err
	}

	return age, err
}

func GetFieldNation(name string) (nationals *Nationals, err error) {
	r, err := http.Get("https://api.nationalize.io/?name=" + name)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(r.Body).Decode(&nationals)
	if err != nil {
		return nil, err
	}

	return nationals, err
}

func GetFieldGender(name string) (gen *Genderaze, err error) {
	r, err := http.Get("https://api.genderize.io/?name=" + name)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(r.Body).Decode(&gen)
	if err != nil {
		return nil, err
	}

	return gen, err
}
