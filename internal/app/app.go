package app

import (
	"fmt"
	"hexhoc/go-hh/config"
	"hexhoc/go-hh/internal/repository"
	"hexhoc/go-hh/internal/usecase"
	postgres "hexhoc/go-hh/pkg/datasource"
	"log"
)

func Run(cfg *config.Config) {

	// Repository
	pg, err := postgres.New(cfg.Datasource.Url, postgres.MaxPoolSize(cfg.Datasource.PoolMax))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()
	defer End()

	statisticsRepository := repository.NewStatisticsPostgres(pg)
	vacancyRepository := repository.NewVacancyPostgres(pg)
	statisticsService := usecase.NewStatisticsService(statisticsRepository, vacancyRepository)
	vacancyService := usecase.NewVacancyService(vacancyRepository, statisticsService)

	vacancyService.Load()

}

func End() {
	fmt.Println("END PROGRAM")
}