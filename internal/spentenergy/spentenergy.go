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

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	RunningSpentCalories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return RunningSpentCalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("the number of steps must be greater than 0")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("weight must be greater than 0")
	}

	if height <= 0 {
		return 0, fmt.Errorf("height must be greater than 0")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("duration must be greater than 0")
	}

	MeanSpeed := MeanSpeed(steps, height, duration)

	return weight * MeanSpeed * duration.Minutes() / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return float64(steps) * height * stepLengthCoefficient / mInKm
}
