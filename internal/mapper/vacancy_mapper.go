package mapper

import (
	"hexhoc/go-hh/internal/dto"
	"hexhoc/go-hh/internal/entity"
	"strconv"
)

func ItemDtoToVacancy(item *dto.Item, vacancyType string) *entity.Vacancy {
	id, _ := strconv.ParseInt(item.Id, 10, 64)
	var avgSalary int32 = 0
	if item.Salary.From != 0 && item.Salary.To != 0 {
		avgSalary = (item.Salary.From + item.Salary.To) / 2
	} else if item.Salary.From != 0 && item.Salary.To == 0 {
		avgSalary = item.Salary.From
	} else if item.Salary.From == 0 && item.Salary.To != 0 {
		avgSalary = item.Salary.To
	}
	employerId, _ := strconv.ParseInt(item.Employer.Id, 10, 64)

	//vacancy.setEmployer(employerMapper.toEntity(item.getEmployer()));
	//if (Objects.nonNull(item.getSchedule())) {
	//	vacancy.setSchedule(item.getSchedule().getName());
	//}
	vacancy := entity.Vacancy{
		Id:                     id,
		Premium:                item.Premium,
		Name:                   item.Name,
		HasTest:                item.HasTest,
		ResponseLetterRequired: item.ResponseLetterRequired,
		Area:                   item.Area.Name,
		SalaryFrom:             item.Salary.From,
		SalaryTo:               item.Salary.To,
		SalaryAvg:              avgSalary,
		SalaryCurrency:         item.Salary.Currency,
		SalaryGross:            item.Salary.Gross,
		PublishedAt:            item.PublishedAt,
		CreatedAt:              item.CreatedAt,
		Archived:               item.Archived,
		Url:                    item.Url,
		AlternateUrl:           item.AlternateUrl,
		Schedule:               item.Schedule.Name,
		AcceptTemporary:        item.AcceptTemporary,
		EmployerId:             employerId,
		DataType:               vacancyType,
	}

	return &vacancy

}

func ItemsDtoToVacancies(items []dto.Item, vacancyType string) []*entity.Vacancy {
	var vacancies = make([]*entity.Vacancy, len(items))
	for i, item := range items {
		vacancies[i] = ItemDtoToVacancy(&item, vacancyType)
	}

	return vacancies
}
