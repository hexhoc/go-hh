package usecase

import (
	"github.com/gofrs/uuid"
	"hexhoc/go-hh/internal/entity"
	"hexhoc/go-hh/internal/interfaces"
	"time"
)

type StatisticsUseCase struct {
	statisticsRepository interfaces.StatisticsRepository
	vacancyRepository    interfaces.VacancyRepository
}

func NewStatisticsUseCase(
	statisticsRepository interfaces.StatisticsRepository,
	vacancyRepository interfaces.VacancyRepository) *StatisticsUseCase {

	return &StatisticsUseCase{
		statisticsRepository,
		vacancyRepository,
	}
}

func (s *StatisticsUseCase) CalculateStatistics(statisticName string) {
	uuid4, _ := uuid.NewV4()
	statistics := entity.Statistics{
		Id:                    uuid4,
		Date:                  time.Now(),
		Name:                  statisticName,
		TotalVacancies:        s.vacancyRepository.CountByType(statisticName),
		AverageSalary:         s.vacancyRepository.GetAverageSalary(statisticName),
		SalaryLess100k:        s.vacancyRepository.GetCountWithSalaryLess100k(statisticName),
		SalaryBetween100k150k: s.vacancyRepository.GetCountWithSalaryBetween100k150k(statisticName),
		SalaryBetween150k200k: s.vacancyRepository.GetCountWithSalaryBetween150k200k(statisticName),
		SalaryBetween200k250k: s.vacancyRepository.GetCountWithSalaryBetween200k250k(statisticName),
		SalaryBetween250k300k: s.vacancyRepository.GetCountWithSalaryBetween250k300k(statisticName),
		SalaryBetween300k350k: s.vacancyRepository.GetCountWithSalaryBetween300k350k(statisticName),
		SalaryBetween350k400k: s.vacancyRepository.GetCountWithSalaryBetween350k400k(statisticName),
		SalaryMore400k:        s.vacancyRepository.GetCountWithSalaryMore400k(statisticName),
		WithoutSalary:         s.vacancyRepository.GetCountWithoutSalary(statisticName),
		CountApplicant:        0,
		ApplicantPerVacancy:   0,
	}

	s.statisticsRepository.Save(&statistics)
}
