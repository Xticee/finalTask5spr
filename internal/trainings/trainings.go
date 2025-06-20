package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	dataParts := strings.Split(datastring, ",") // Getting count of steps, training type and duration from string. Format "3456,Ходьба,3h00m"

	if len(dataParts) != 3 {
		return fmt.Errorf("invalid data format")
	}

	steps, err := strconv.Atoi(dataParts[0])
	if err != nil {
		return fmt.Errorf("conversion error, %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("count of steps must be greater than zero")
	}
	t.Steps = steps

	t.TrainingType = dataParts[1]

	duration, err := time.ParseDuration(dataParts[2])
	if err != nil {
		return fmt.Errorf("conversion error, %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("duration must be greater than zero")
	}

	t.Duration = duration

	return
}

func (t Training) ActionInfo() (string, error) {

	distance := spentenergy.Distance(t.Steps, t.Height)
	averageSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	training := t.TrainingType
	var (
		calories float64
		err      error
	)

	switch training {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", fmt.Errorf("something went wrong : %w", err)
	}

	outputString := fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, training, t.Duration.Hours(), distance, averageSpeed, calories)

	return outputString, nil
}
