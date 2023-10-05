package usecase

import (
	"encoding/json"
	"fmt"
	"hexhoc/go-hh/internal/dto"
	"hexhoc/go-hh/internal/entity"
	"hexhoc/go-hh/internal/interfaces"
	"hexhoc/go-hh/internal/mapper"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type VacancyUseCase struct {
	vacancyRepository interfaces.VacancyRepository
	statisticsUseCase interfaces.StatisticsUseCase
}

func NewVacancyUseCase(
	vacancyRepository interfaces.VacancyRepository,
	statisticsUseCase interfaces.StatisticsUseCase) *VacancyUseCase {

	return &VacancyUseCase{
		vacancyRepository,
		statisticsUseCase,
	}
}

func (v *VacancyUseCase) Load() {

	v.DeleteAll()

	m := make(map[string]string)
	m["JAVA"] = "https://api.hh.ru/vacancies?search_field=name&text=NAME:(java%20AND%20NOT%20script)%20AND%20DESCRIPTION:(NOT%20Android)&area=1"
	m["SCALA"] = "https://api.hh.ru/vacancies?search_field=name&text=scala&area=1"
	m["RUST"] = "https://api.hh.ru/vacancies?search_field=name&text=rust&area=1"
	m["KOTLIN"] = "https://api.hh.ru/vacancies?search_field=name&text=kotlin&area=1"
	m["PHP"] = "https://api.hh.ru/vacancies?search_field=name&text=PHP&area=1"
	m["PYTHON"] = "https://api.hh.ru/vacancies?search_field=name&text=python&area=1"
	m["GOLANG"] = "https://api.hh.ru/vacancies?search_field=name&text=name:(go%20or%20golang)&area=1"
	m["1C"] = "https://api.hh.ru/vacancies?search_field=name&text=name:(1%D0%A1%20AND%20NOT%20bitrix%20AND%20NOT%20%D1%82%D0%B5%D1%81%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D1%89%D0%B8%D0%BA%20AND%20NOT%20%D0%B0%D0%BD%D0%B0%D0%BB%D0%B8%D1%82%D0%B8%D0%BA%20AND%20NOT%20%D0%9A%D0%BE%D0%BD%D1%81%D1%83%D0%BB%D1%8C%D1%82%D0%B0%D0%BD%D1%82%20AND%20NOT%20%D0%90%D0%B4%D0%BC%D0%B8%D0%BD%D0%B8%D1%81%D1%82%D1%80%D0%B0%D1%82%D0%BE%D1%80%20AND%20NOT%20%D0%BE%D0%BF%D0%B5%D1%80%D0%B0%D1%82%D0%BE%D1%80%20AND%20NOT%20%D0%B1%D1%83%D1%85%D0%B3%D0%B0%D0%BB%D1%82%D0%B5%D1%80)&area=1"
	m["FRONTEND"] = "https://api.hh.ru/vacancies?search_field=name&text=name:(frontend%20OR%20react%20OR%20node.js%20OR%20vue%20OR%20angular%20OR%20javascript)&area=1"
	m["DEVOPS"] = "https://api.hh.ru/vacancies?search_field=name&text=devops&area=1"

	// var wg sync.WaitGroup

	for vacancyType, url := range m {
		// Increment the WaitGroup counter.
		// wg.Add(1)
		// go func(vacancyType string, url string) {
		// Decrement the counter when the goroutine completes.
		// defer wg.Done()
		v.downloadVacanciesByType(vacancyType, url)
		v.statisticsUseCase.CalculateStatistics(vacancyType)
		fmt.Println(vacancyType)
		// }(vacancyType, url)
	}

	// Wait for all HTTP fetches to complete.
	// wg.Wait()
}

func (v *VacancyUseCase) downloadVacanciesByType(vacancyType string, url string) {
	totalPage := 1
	for currentPage := 0; currentPage < totalPage; currentPage++ {
		//resourceUrl := value + "&page="+currentPage+"&per_page=100"
		resourceUrl := url + "&page=" + strconv.Itoa(currentPage) + "&per_page=100"
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

		v.SaveAll(mapper.ItemsDtoToVacancies(hhResponse.Items, vacancyType))
	}

}

func (v *VacancyUseCase) SaveAll(vacancies []*entity.Vacancy) {
	for _, vacancy := range vacancies {
		v.convertGrossToNetSalary(vacancy)
	}
	v.vacancyRepository.SaveAll(vacancies)
}

func (v *VacancyUseCase) convertGrossToNetSalary(vacancy *entity.Vacancy) {

	if vacancy.SalaryGross {
		vacancy.SalaryGross = false
		vacancy.SalaryFrom = int32(float64(vacancy.SalaryFrom) * 0.87)
		vacancy.SalaryTo = int32(float64(vacancy.SalaryTo) * 0.87)
		vacancy.SalaryAvg = int32(float64(vacancy.SalaryAvg) * 0.87)
	}

}

func (v *VacancyUseCase) DeleteAll() {
	v.vacancyRepository.DeleteAll()
}

func (v *VacancyUseCase) CountByType(statisticsName string) int32 {
	return v.vacancyRepository.CountByType(statisticsName)
}

func (v *VacancyUseCase) GetAverageSalary(statisticsName string) int32 {
	return v.vacancyRepository.GetAverageSalary(statisticsName)
}

func (v *VacancyUseCase) GetCountWithSalaryLess100k(statisticsName string) int32 {
	return v.vacancyRepository.GetCountWithSalaryLess100k(statisticsName)
}

func (v *VacancyUseCase) GetCountWithSalaryBetween100k150k(statisticsName string) int32 {
	return v.vacancyRepository.GetCountWithSalaryBetween100k150k(statisticsName)
}

func (v *VacancyUseCase) GetCountWithSalaryBetween150k200k(statisticsName string) int32 {
	return v.vacancyRepository.GetCountWithSalaryBetween150k200k(statisticsName)
}

func (v *VacancyUseCase) GetCountWithSalaryBetween200k250k(statisticsName string) int32 {
	return v.vacancyRepository.GetCountWithSalaryBetween200k250k(statisticsName)
}

func (v *VacancyUseCase) GetCountWithSalaryBetween250k300k(statisticsName string) int32 {
	return v.vacancyRepository.GetCountWithSalaryBetween250k300k(statisticsName)
}

func (v *VacancyUseCase) GetCountWithSalaryBetween300k350k(statisticsName string) int32 {
	return v.vacancyRepository.GetCountWithSalaryBetween300k350k(statisticsName)
}

func (v *VacancyUseCase) GetCountWithSalaryBetween350k400k(statisticsName string) int32 {
	return v.vacancyRepository.GetCountWithSalaryBetween350k400k(statisticsName)
}

func (v *VacancyUseCase) GetCountWithSalaryMore400k(statisticsName string) int32 {
	return v.vacancyRepository.GetCountWithSalaryMore400k(statisticsName)
}

func (v *VacancyUseCase) GetCountWithoutSalary(statisticsName string) int32 {
	return v.vacancyRepository.GetCountWithoutSalary(statisticsName)
}
