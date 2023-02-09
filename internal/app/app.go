package app

import (
	"fmt"
	"hexhoc/go-hh/config"
	"hexhoc/go-hh/internal/interfaces"
	"hexhoc/go-hh/internal/repository"
	"hexhoc/go-hh/internal/usecase"
	postgres "hexhoc/go-hh/pkg/datasource"
	"log"
	"time"
)

func Run(cfg *config.Config) {
	start := time.Now()

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()
	defer End()

	var statisticsRepository interfaces.StatisticsRepository = repository.NewStatisticsPostgres(pg)
	var vacancyRepository interfaces.VacancyRepository = repository.NewVacancyPostgres(pg)
	var statisticsUseCase interfaces.StatisticsUseCase = usecase.NewStatisticsUseCase(statisticsRepository, vacancyRepository)
	var vacancyUseCase interfaces.VacancyUseCase = usecase.NewVacancyUseCase(vacancyRepository, statisticsUseCase)

	vacancyUseCase.Load()

	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println(fmt.Sprintf("Total time: %s", elapsed))
	fmt.Println(elapsed)
}

func End() {
	fmt.Println("END PROGRAM")
}
