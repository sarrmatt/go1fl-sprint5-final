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

// WalkingSpentCalories вычисляет количество калорий, потраченных при ходьбе,
// по количеству шагов, весу и росту пользователя и продолжительности ходьбы.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным, получено %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть положительным, получен %v", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть положительным, получен %v", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("длительность должна быть положительной, получена %v", duration)
	}

	speed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * speed * durationInMinutes / minInH) * walkingCaloriesCoefficient

	return calories, nil
}

// RunningSpentCalories вычисляет количество калорий, потраченных при беге,
// по количеству шагов, весу и росту пользователя и продолжительности бега.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным, получено %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть положительным, получен %v", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть положительным, получен %v", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("длительность должна быть положительной, получена %v", duration)
	}

	speed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := weight * speed * durationInMinutes / minInH

	return calories, nil
}

// MeanSpeed вычисляет среднюю скорость в км/ч по количеству шагов,
// росту пользователя в метрах и продолжительности активности.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

// Distance вычисляет пройденную дистанцию в километрах по количеству шагов
// и росту пользователя в метрах.
func Distance(steps int, height float64) float64 {
	return (float64(steps) * (height * stepLengthCoefficient)) / mInKm
}
