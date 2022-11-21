package usecase

import (
	"hexhoc/go-hh/internal/entity"
	"hexhoc/go-hh/internal/repository"
	"time"
)

type StatisticsService struct {
	statisticsRepository *repository.StatisticsPostgres
	vacancyRepository    *repository.VacancyPostgres
}

func NewStatisticsService(
	statisticsRepository *repository.StatisticsPostgres,
	vacancyRepository *repository.VacancyPostgres) *StatisticsService {

	return &StatisticsService{
		statisticsRepository,
		vacancyRepository,
	}
}

func (statisticsService *StatisticsService) CalculateStatistics(statisticName string) {

	statistics := entity.Statistics{
		Date:                  time.Now(),
		Name:                  statisticName,
		TotalVacancies:        statisticsService.vacancyRepository.CountByType(statisticName),
		AverageSalary:         statisticsService.vacancyRepository.GetAverageSalary(statisticName),
		SalaryLess100k:        statisticsService.vacancyRepository.GetCountWithSalaryLess100k(statisticName),
		SalaryBetween100k150k: statisticsService.vacancyRepository.GetCountWithSalaryBetween100k150k(statisticName),
		SalaryBetween150k200k: statisticsService.vacancyRepository.GetCountWithSalaryBetween150k200k(statisticName),
		SalaryBetween200k250k: statisticsService.vacancyRepository.GetCountWithSalaryBetween200k250k(statisticName),
		SalaryBetween250k300k: statisticsService.vacancyRepository.GetCountWithSalaryBetween250k300k(statisticName),
		SalaryBetween300k350k: statisticsService.vacancyRepository.GetCountWithSalaryBetween300k350k(statisticName),
		SalaryBetween350k400k: statisticsService.vacancyRepository.GetCountWithSalaryBetween350k400k(statisticName),
		SalaryMore400k:        statisticsService.vacancyRepository.GetCountWithSalaryMore400k(statisticName),
		WithoutSalary:         statisticsService.vacancyRepository.GetCountWithoutSalary(statisticName),
		CountApplicant:        0,
		ApplicantPerVacancy:   0,
	}

	statisticsService.statisticsRepository.Save(&statistics)
}
