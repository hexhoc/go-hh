package entity

type Vacancy struct {
	Id                     int64  `json:"id"`
	Premium                bool   `json:"premium"`
	Name                   string `json:"name"`
	HasTest                bool   `json:"hasTest"`
	ResponseLetterRequired bool   `json:"responseLetterRequired"`
	Area                   string `json:"area"`
	SalaryFrom             int32  `json:"salaryFrom"`
	SalaryTo               int32  `json:"salaryTo"`
	SalaryAvg              int32  `json:"salaryAvg"`
	SalaryCurrency         string `json:"salaryCurrency"`
	SalaryGross            bool   `json:"salaryGross"`
	PublishedAt            string `json:"publishedAt"`
	CreatedAt              string `json:"createdAt"`
	Archived               bool   `json:"archived"`
	Url                    string `json:"url"`
	AlternateUrl           string `json:"alternateUrl"`
	Schedule               string `json:"schedule"`
	AcceptTemporary        bool   `json:"acceptTemporary"`
	EmployerId             int64  `json:"employer_id"`
	DataType               string `json:"type"`
}
