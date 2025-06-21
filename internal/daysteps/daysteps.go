package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataParts := strings.Split(datastring, ",") // Getting count of steps and duration from string. Format "3456,3h00m"

	if len(dataParts) != 2 {
		return fmt.Errorf("invalid data format")
	}

	steps, err := strconv.Atoi(dataParts[0])
	if err != nil {
		return fmt.Errorf("conversion error, %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("count of steps must be greater than zero")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(dataParts[1])
	if err != nil {
		return fmt.Errorf("conversion error, %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("duration must be greater than zero")
	}

	ds.Duration = duration

	return
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", fmt.Errorf("something went wrong : %w", err)
	}

	outputString := fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`, ds.Steps, distance, calories)

	return outputString, nil
}
