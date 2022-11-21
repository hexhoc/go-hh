package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"hexhoc/go-hh/internal/entity"
	postgres "hexhoc/go-hh/pkg/datasource"
	"log"
)

type VacancyPostgres struct {
	db *postgres.Postgres
}

func NewVacancyPostgres(db *postgres.Postgres) *VacancyPostgres {
	return &VacancyPostgres{db}
}

func (vacancyPostgres *VacancyPostgres) DeleteAll() {
	_, err := vacancyPostgres.db.Pool.Exec(context.Background(), "DELETE FROM vacancy")
	if err != nil {
		log.Fatalf("(vacancyPostgres *VacancyPostgres) DeleteAll(): %s", err)
		return
	}
}

func (vacancyPostgres *VacancyPostgres) SaveAll(vacancies []*entity.Vacancy) {
	if len(vacancies) == 0 {
		return
	}

	batch := &pgx.Batch{}
	for _, v := range vacancies {
		query := `INSERT INTO vacancy(
                    accept_temporary, 
                    alternate_url, 
                    archived, 
                    area, 
                    created_at, 
                    has_test, 
                    name, 
                    premium, 
                    published_at, 
                    response_letter_required, 
                    salary_avg, 
                    salary_currency, 
                    salary_gross, 
                    schedule, 
                    type, 
                    url, 
                    employer_id) 
				  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)`
		batch.Queue(query,
			v.AcceptTemporary,
			v.AlternateUrl,
			v.Archived,
			v.Area,
			v.CreatedAt,
			v.HasTest,
			v.Name,
			v.Premium,
			v.PublishedAt,
			v.ResponseLetterRequired,
			v.SalaryAvg,
			v.SalaryCurrency,
			v.SalaryGross,
			v.Schedule,
			v.DataType,
			v.Url,
			v.EmployerId,
		)
	}
	res := vacancyPostgres.db.Pool.SendBatch(context.Background(), batch)
	defer res.Close()
}

func (vacancyPostgres *VacancyPostgres) CountByType(statisticsName string) int32 {
	return vacancyPostgres.getCount("SELECT COUNT(id) FROM vacancy WHERE vacancy.type = $1::varchar", statisticsName)
}

func (vacancyPostgres *VacancyPostgres) GetAverageSalary(name string) int32 {
	return vacancyPostgres.getCount("select round(avg(salary_avg)) from vacancy where salary_avg > 0 and salary_currency = 'RUR' and type = $1::varchar", name)
}

func (vacancyPostgres *VacancyPostgres) GetCountWithSalaryLess100k(name string) int32 {
	return vacancyPostgres.getCount("select count(salary_avg) from vacancy where salary_avg between 1 and 100000 and salary_currency = 'RUR' and type = $1::varchar", name)
}

func (vacancyPostgres *VacancyPostgres) GetCountWithSalaryBetween100k150k(name string) int32 {
	return vacancyPostgres.getCount("select count(salary_avg) from vacancy where salary_avg between 100000 and 150000 and salary_currency = 'RUR' and type = $1::varchar", name)
}

func (vacancyPostgres *VacancyPostgres) GetCountWithSalaryBetween150k200k(name string) int32 {
	return vacancyPostgres.getCount("select count(salary_avg) from vacancy where salary_avg between 150000 and 200000 and salary_currency = 'RUR' and type = $1::varchar", name)
}

func (vacancyPostgres *VacancyPostgres) GetCountWithSalaryBetween200k250k(name string) int32 {
	return vacancyPostgres.getCount("select count(salary_avg) from vacancy where salary_avg between 200000 and 250000 and salary_currency = 'RUR' and type = $1::varchar", name)
}

func (vacancyPostgres *VacancyPostgres) GetCountWithSalaryBetween250k300k(name string) int32 {
	return vacancyPostgres.getCount("select count(salary_avg) from vacancy where salary_avg between 250000 and 300000 and salary_currency = 'RUR' and type = $1::varchar", name)
}

func (vacancyPostgres *VacancyPostgres) GetCountWithSalaryBetween300k350k(name string) int32 {
	return vacancyPostgres.getCount("select count(salary_avg) from vacancy where salary_avg between 300000 and 350000 and salary_currency = 'RUR' and type = $1::varchar", name)
}

func (vacancyPostgres *VacancyPostgres) GetCountWithSalaryBetween350k400k(name string) int32 {
	return vacancyPostgres.getCount("select count(salary_avg) from vacancy where salary_avg between 350000 and 400000 and salary_currency = 'RUR' and type = $1::varchar", name)
}

func (vacancyPostgres *VacancyPostgres) GetCountWithSalaryMore400k(name string) int32 {
	return vacancyPostgres.getCount("select count(salary_avg) from vacancy where salary_avg > 400000 and salary_currency = 'RUR' and type = $1::varchar", name)
}

func (vacancyPostgres *VacancyPostgres) GetCountWithoutSalary(name string) int32 {
	return vacancyPostgres.getCount("select count(salary_avg) from vacancy where salary_avg = 0 and type = $1::varchar", name)
}

func (vacancyPostgres *VacancyPostgres) getCount(query string, statisticsName string) int32 {
	rows, err := vacancyPostgres.db.Pool.Query(context.Background(), query, statisticsName)

	if err != nil {
		log.Fatalf("(vacancyPostgres *VacancyPostgres) getCount: %s", err)
		return 0
	}
	// iterate through the rows
	var result int32 = 0
	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			log.Println("error while iterating dataset")
		}

		//values, err := rows.Values()
		//if err != nil {
		//	log.Fatal("error while iterating dataset")
		//}
		//
		//result = int32(values[0].(int64))
	}
	defer rows.Close()

	return result

}
