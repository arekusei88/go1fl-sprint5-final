// Package spentenergy считает дистанцию, среднюю скорость и калории.
package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// checkParams проверяет входные данные расчёта калорий.
func checkParams(steps int, weight, height float64, duration time.Duration) error {
	switch {
	case steps <= 0:
		return fmt.Errorf("количество шагов должно быть положительным: %d", steps)
	case weight <= 0:
		return fmt.Errorf("вес должен быть положительным: %.2f", weight)
	case height <= 0:
		return fmt.Errorf("рост должен быть положительным: %.2f", height)
	case duration <= 0:
		return fmt.Errorf("длительность должна быть положительной: %v", duration)
	}
	return nil
}

// Distance рассчитывает дистанцию в километрах; рост задаётся в метрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

// MeanSpeed рассчитывает среднюю скорость в км/ч; при неположительных шагах
// или длительности возвращает 0.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// RunningSpentCalories рассчитывает калории, потраченные при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := checkParams(steps, weight, height, duration); err != nil {
		return 0, fmt.Errorf("расчёт калорий при беге: %w", err)
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

// WalkingSpentCalories рассчитывает калории, потраченные при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := checkParams(steps, weight, height, duration); err != nil {
		return 0, fmt.Errorf("расчёт калорий при ходьбе: %w", err)
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH * walkingCaloriesCoefficient, nil
}
