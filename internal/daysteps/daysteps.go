// Package daysteps разбирает строки с данными о прогулках и формирует по ним сводку.
package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps хранит данные одной прогулки.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "678,0h50m" и заполняет поля структуры.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("в строке %q ожидалось 2 поля, получено %d", datastring, len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("некорректное количество шагов %q: %w", parts[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным: %d", steps)
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("некорректная длительность %q: %w", parts[1], err)
	}
	if duration <= 0 {
		return fmt.Errorf("длительность должна быть положительной: %v", duration)
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// ActionInfo формирует строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("расчёт данных прогулки: %w", err)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
