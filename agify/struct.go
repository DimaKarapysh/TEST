package agify

type Genderaze struct {
	Id          int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

type Nationals struct {
	Id     int       `json:"count"`
	Name   string    `json:"name"`
	Nation []Country `json:"country"`
}

type Country struct {
	Id          string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type AgilyAge struct {
	Id   int    `json:"count"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
