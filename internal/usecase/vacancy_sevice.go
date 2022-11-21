package usecase

import (
	"encoding/json"
	"fmt"
	"hexhoc/go-hh/internal/dto"
	"hexhoc/go-hh/internal/entity"
	"hexhoc/go-hh/internal/mapper"
	"hexhoc/go-hh/internal/repository"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type VacancyService struct {
	vacancyRepository *repository.VacancyPostgres
	statisticsService *StatisticsService
}

func NewVacancyService(
	vacancyRepository *repository.VacancyPostgres,
	statisticsService *StatisticsService) *VacancyService {

	return &VacancyService{
		vacancyRepository,
		statisticsService,
	}
}

func (vacancyService *VacancyService) Load() {

	vacancyService.DeleteAll()

	m := make(map[string]string)
	m["1C"] = "https://api.hh.ru/vacancies?search_field=name&text=name:(1%D0%A1%20AND%20NOT%20bitrix%20AND%20NOT%20%D1%82%D0%B5%D1%81%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D1%89%D0%B8%D0%BA%20AND%20NOT%20%D0%B0%D0%BD%D0%B0%D0%BB%D0%B8%D1%82%D0%B8%D0%BA%20AND%20NOT%20%D0%9A%D0%BE%D0%BD%D1%81%D1%83%D0%BB%D1%8C%D1%82%D0%B0%D0%BD%D1%82%20AND%20NOT%20%D0%90%D0%B4%D0%BC%D0%B8%D0%BD%D0%B8%D1%81%D1%82%D1%80%D0%B0%D1%82%D0%BE%D1%80%20AND%20NOT%20%D0%BE%D0%BF%D0%B5%D1%80%D0%B0%D1%82%D0%BE%D1%80%20AND%20NOT%20%D0%B1%D1%83%D1%85%D0%B3%D0%B0%D0%BB%D1%82%D0%B5%D1%80)&area=1"
	m["JAVA"] = "https://api.hh.ru/vacancies?search_field=name&text=java&area=1"
	m["PHP"] = "https://api.hh.ru/vacancies?search_field=name&text=PHP&area=1"
	m["PYTHON"] = "https://api.hh.ru/vacancies?search_field=name&text=python&area=1"
	m["GOLANG"] = "https://api.hh.ru/vacancies?search_field=name&text=name:(go%20or%20golang)&area=1"
	m["FRONTEND"] = "https://api.hh.ru/vacancies?search_field=name&text=name:(frontend%20OR%20react%20OR%20node.js%20OR%20vue%20OR%20angular%20OR%20javascript)&area=1"

	for k, v := range m {
		totalPage := 1

		for currentPage := 0; currentPage < totalPage; currentPage++ {
			//resourceUrl := v + "&page="+currentPage+"&per_page=100"
			resourceUrl := v + "&page=" + strconv.Itoa(currentPage) + "&per_page=100"
			response, err := http.Get(resourceUrl)
			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			responseData, err := io.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			var hhResponse dto.HhResponse
			json.Unmarshal(responseData, &hhResponse)

			if currentPage == 0 {
				totalPage = hhResponse.Pages
			}

			vacancyService.SaveAll(mapper.ItemsDtoToVacancies(hhResponse.Items, k))

		}

		vacancyService.statisticsService.CalculateStatistics(k)
		fmt.Println(k)
	}

}

func (vacancyService *VacancyService) SaveAll(vacancies []*entity.Vacancy) {
	for _, vacancy := range vacancies {
		vacancyService.convertGrossToNetSalary(vacancy)
	}
	vacancyService.vacancyRepository.SaveAll(vacancies)
}

func (vacancyService *VacancyService) convertGrossToNetSalary(vacancy *entity.Vacancy) {

	if vacancy.SalaryGross {
		vacancy.SalaryGross = false
		vacancy.SalaryFrom = int32(float64(vacancy.SalaryFrom) * 0.87)
		vacancy.SalaryTo = int32(float64(vacancy.SalaryTo) * 0.87)
		vacancy.SalaryAvg = int32(float64(vacancy.SalaryAvg) * 0.87)
	}

}

func (vacancyService *VacancyService) DeleteAll() {
	vacancyService.vacancyRepository.DeleteAll()
}

func (vacancyService *VacancyService) CountByType(statisticsName string) int32 {
	return vacancyService.vacancyRepository.CountByType(statisticsName)
}

func (vacancyService *VacancyService) GetAverageSalary(statisticsName string) int32 {
	return vacancyService.vacancyRepository.GetAverageSalary(statisticsName)
}

func (vacancyService *VacancyService) GetCountWithSalaryLess100k(statisticsName string) int32 {
	return vacancyService.vacancyRepository.GetCountWithSalaryLess100k(statisticsName)
}

func (vacancyService *VacancyService) GetCountWithSalaryBetween100k150k(statisticsName string) int32 {
	return vacancyService.vacancyRepository.GetCountWithSalaryBetween100k150k(statisticsName)
}

func (vacancyService *VacancyService) GetCountWithSalaryBetween150k200k(statisticsName string) int32 {
	return vacancyService.vacancyRepository.GetCountWithSalaryBetween150k200k(statisticsName)
}

func (vacancyService *VacancyService) GetCountWithSalaryBetween200k250k(statisticsName string) int32 {
	return vacancyService.vacancyRepository.GetCountWithSalaryBetween200k250k(statisticsName)
}

func (vacancyService *VacancyService) GetCountWithSalaryBetween250k300k(statisticsName string) int32 {
	return vacancyService.vacancyRepository.GetCountWithSalaryBetween250k300k(statisticsName)
}

func (vacancyService *VacancyService) GetCountWithSalaryBetween300k350k(statisticsName string) int32 {
	return vacancyService.vacancyRepository.GetCountWithSalaryBetween300k350k(statisticsName)
}

func (vacancyService *VacancyService) GetCountWithSalaryBetween350k400k(statisticsName string) int32 {
	return vacancyService.vacancyRepository.GetCountWithSalaryBetween350k400k(statisticsName)
}

func (vacancyService *VacancyService) GetCountWithSalaryMore400k(statisticsName string) int32 {
	return vacancyService.vacancyRepository.GetCountWithSalaryMore400k(statisticsName)
}

func (vacancyService *VacancyService) GetCountWithoutSalary(statisticsName string) int32 {
	return vacancyService.vacancyRepository.GetCountWithoutSalary(statisticsName)
}
