// Package trainings разбирает строки с данными о тренировках и формирует по ним сводку.
package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training хранит данные одной тренировки.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "3456,Ходьба,3h00m" и заполняет поля структуры.
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("в строке %q ожидалось 3 поля, получено %d", datastring, len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("некорректное количество шагов %q: %w", parts[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным: %d", steps)
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("некорректная длительность %q: %w", parts[2], err)
	}
	if duration <= 0 {
		return fmt.Errorf("длительность должна быть положительной: %v", duration)
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration
	return nil
}

// ActionInfo формирует строку с информацией о проведённой тренировке.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}
	if err != nil {
		return "", fmt.Errorf("расчёт данных тренировки %q: %w", t.TrainingType, err)
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}
