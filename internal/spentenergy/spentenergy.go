package spentenergy

import (
	"fmt"
	"log"
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
	if steps <= 0 {
		return 0.0, fmt.Errorf("count of steps must be greater than zero")
	}

	if weight <= 0 {
		return 0.0, fmt.Errorf("user weight must be greater than zero")
	}

	if height <= 0 {
		return 0.0, fmt.Errorf("user height must be greater than zero")
	}

	if duration <= 0 {

		return 0.0, fmt.Errorf("duration must be greater than zero")
	}

	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0.0, fmt.Errorf("count of steps must be greater than zero")
	}

	if weight <= 0 {
		return 0.0, fmt.Errorf("user weight must be greater than zero")
	}

	if height <= 0 {
		return 0.0, fmt.Errorf("user height must be greater than zero")
	}

	if duration <= 0 {

		return 0.0, fmt.Errorf("duration must be greater than zero")
	}

	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		log.Println("duration must be greater than zero")
		return 0.0
	}

	distance := Distance(steps, height)

	return distance / duration.Hours() // Calculating average speed
}

func Distance(steps int, height float64) float64 {

	return height * stepLengthCoefficient * float64(steps) / mInKm // Calculating traveled distance using user height, step length coefficient and count of steps; in Km
}
