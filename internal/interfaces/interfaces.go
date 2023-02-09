package interfaces

import "hexhoc/go-hh/internal/entity"

type StatisticsRepository interface {
	DeleteAll()
	Save(s *entity.Statistics)
}

type VacancyRepository interface {
	DeleteAll()
	SaveAll(vacancies []*entity.Vacancy)
	CountByType(statisticsName string) int32
	GetAverageSalary(name string) int32
	GetCountWithSalaryLess100k(name string) int32
	GetCountWithSalaryBetween100k150k(name string) int32
	GetCountWithSalaryBetween150k200k(name string) int32
	GetCountWithSalaryBetween200k250k(name string) int32
	GetCountWithSalaryBetween250k300k(name string) int32
	GetCountWithSalaryBetween300k350k(name string) int32
	GetCountWithSalaryBetween350k400k(name string) int32
	GetCountWithSalaryMore400k(name string) int32
	GetCountWithoutSalary(name string) int32
}

type VacancyUseCase interface {
	Load()
	DeleteAll()
	SaveAll(vacancies []*entity.Vacancy)
	CountByType(statisticsName string) int32
	GetAverageSalary(name string) int32
	GetCountWithSalaryLess100k(name string) int32
	GetCountWithSalaryBetween100k150k(name string) int32
	GetCountWithSalaryBetween150k200k(name string) int32
	GetCountWithSalaryBetween200k250k(name string) int32
	GetCountWithSalaryBetween250k300k(name string) int32
	GetCountWithSalaryBetween300k350k(name string) int32
	GetCountWithSalaryBetween350k400k(name string) int32
	GetCountWithSalaryMore400k(name string) int32
	GetCountWithoutSalary(name string) int32
}

type StatisticsUseCase interface {
	CalculateStatistics(statisticName string)
}
