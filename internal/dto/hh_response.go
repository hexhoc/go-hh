package dto

type HhResponse struct {
	Items        []Item      `json:"items"`
	Found        int         `json:"found"`
	Pages        int         `json:"pages"`
	PerPage      int         `json:"per_page"`
	Page         int         `json:"page"`
	Clusters     interface{} `json:"clusters"`
	Arguments    interface{} `json:"arguments"`
	AlternateUrl string      `json:"alternate_url"`
}

type Item struct {
	Id                     string        `json:"id"`
	Premium                bool          `json:"premium"`
	Name                   string        `json:"name"`
	Department             interface{}   `json:"department"`
	HasTest                bool          `json:"has_test"`
	ResponseLetterRequired bool          `json:"response_letter_required"`
	Area                   Area          `json:"area"`
	Salary                 Salary        `json:"salary"`
	DataType               Type          `json:"type"`
	Address                Address       `json:"address"`
	ResponseUrl            interface{}   `json:"response_url"`
	SortPointDistance      interface{}   `json:"sort_point_distance"`
	PublishedAt            string        `json:"published_at"`
	CreatedAt              string        `json:"created_at"`
	Archived               bool          `json:"archived"`
	ApplyAlternateUrl      string        `json:"apply_alternate_url"`
	InsiderInterview       interface{}   `json:"insider_interview"`
	Url                    string        `json:"url"`
	AdvResponseUrl         string        `json:"adv_response_url"`
	AlternateUrl           string        `json:"alternate_url"`
	Relations              []interface{} `json:"relations"`
	Employer               Employer      `json:"employer"`
	Snippet                Snippet       `json:"snippet"`
	Contacts               interface{}   `json:"contacts"`
	Schedule               Schedule      `json:"schedule"`
	WorkingDays            []interface{} `json:"working_days"`
	WorkingTimeIntervals   []interface{} `json:"working_time_intervals"`
	WorkingTimeModes       []interface{} `json:"working_time_modes"`
	AcceptTemporary        bool          `json:"accept_temporary"`
}

type Area struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Salary struct {
	From     int32  `json:"from"`
	To       int32  `json:"to"`
	Currency string `json:"currency"`
	Gross    bool   `json:"gross"`
}

type Type struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Snippet struct {
	Requirement    string `json:"requirement"`
	Responsibility string `json:"responsibility"`
}
type Schedule struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Address struct {
	City          string      `json:"city"`
	Street        string      `json:"street"`
	Building      string      `json:"building"`
	Description   interface{} `json:"description"`
	Lat           float64     `json:"lat"`
	Lng           float64     `json:"lng"`
	Raw           string      `json:"raw"`
	Metro         Metro       `json:"metro"`
	MetroStations []Metro     `json:"metro_stations"`
	Id            string      `json:"id"`
}

type Metro struct {
	StationName string  `json:"station_name"`
	LineName    string  `json:"line_name"`
	StationId   string  `json:"station_id"`
	LineId      string  `json:"line_id"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
}

type Employer struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Url          string   `json:"url"`
	AlternateUrl string   `json:"alternate_url"`
	LogoUrls     LogoUrls `json:"logo_urls"`
	VacanciesUrl string   `json:"vacancies_url"`
	Trusted      bool     `json:"trusted"`
}

type LogoUrls struct {
	Url90    string `json:"90"`
	Url240   string `json:"240"`
	Original string `json:"original"`
}
