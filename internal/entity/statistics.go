package entity

import "time"

type Statistics struct {
	Id                    int64     `json:"id"`
	Date                  time.Time `json:"date"`
	Name                  string    `json:"name"`
	TotalVacancies        int32     `json:"totalVacancies"`
	AverageSalary         int32     `json:"averageSalary"`
	SalaryLess100k        int32     `json:"salaryLess100k"`
	SalaryBetween100k150k int32     `json:"salaryBetween100k150k"`
	SalaryBetween150k200k int32     `json:"salaryBetween150k200k"`
	SalaryBetween200k250k int32     `json:"salaryBetween200k250k"`
	SalaryBetween250k300k int32     `json:"salaryBetween250k300k"`
	SalaryBetween300k350k int32     `json:"salaryBetween300k350k"`
	SalaryBetween350k400k int32     `json:"salaryBetween350k400k"`
	SalaryMore400k        int32     `json:"salaryMore400k"`
	WithoutSalary         int32     `json:"withoutSalary"`
	CountApplicant        int32     `json:"countApplicant"`
	ApplicantPerVacancy   int32     `json:"applicantPerVacancy"`
}
