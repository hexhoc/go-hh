package entity

type Employer struct {
	Id           int8   `json:"id"`
	Name         string `json:"name"`
	Url          string `json:"url"`
	AlternateUrl string `json:"alternateUrl"`
	VacanciesUrl string `json:"vacanciesUrl"`
	Trusted      string `json:"trusted"`
}
