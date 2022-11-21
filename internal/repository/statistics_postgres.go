package repository

import (
	"context"
	"hexhoc/go-hh/internal/entity"
	postgres "hexhoc/go-hh/pkg/datasource"
	"log"
)

type StatisticsPostgres struct {
	db *postgres.Postgres
}

func NewStatisticsPostgres(db *postgres.Postgres) *StatisticsPostgres {
	return &StatisticsPostgres{db}
}

func (statisticsPostgres *StatisticsPostgres) DeleteAll() {
	_, err := statisticsPostgres.db.Pool.Exec(context.Background(), "DELETE FROM statistics")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (statisticsPostgres *StatisticsPostgres) Save(s *entity.Statistics) {

	query := `INSERT INTO statistics(
                       applicant_per_vacancy, 
                       average_salary, 
                       count_applicant, 
                       date, 
                       name, 
                       salary_between100k150k, 
                       salary_between150k200k, 
                       salary_between200k250k, 
                       salary_between250k300k, 
                       salary_between300k350k, 
                       salary_between350k400k, 
                       salary_less100k, 
                       salary_more400k, 
                       total_vacancies, 
                       without_salary) 
				  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)`
	_, err := statisticsPostgres.db.Pool.Exec(context.Background(), query,
		s.ApplicantPerVacancy,
		s.AverageSalary,
		s.CountApplicant,
		s.Date,
		s.Name,
		s.SalaryBetween100k150k,
		s.SalaryBetween150k200k,
		s.SalaryBetween200k250k,
		s.SalaryBetween250k300k,
		s.SalaryBetween300k350k,
		s.SalaryBetween350k400k,
		s.SalaryLess100k,
		s.SalaryMore400k,
		s.TotalVacancies,
		s.WithoutSalary)
	if err != nil {
		return
	}

	if err != nil {
		log.Fatalf("(statisticsPostgres *StatisticsPostgres) Save: %s", err)
		return
	}
}
