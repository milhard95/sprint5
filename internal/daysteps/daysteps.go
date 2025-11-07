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
	values := strings.Split(datastring, ",")
	if len(values) != 2 {
		return fmt.Errorf("incorrect value received")
	}

	steps, err := strconv.Atoi(values[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("the number of steps must be greater than 0")
	}

	duration, err := time.ParseDuration(values[1])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}

	if duration <= 0 {
		return fmt.Errorf("duration must be greater than 0")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps) +
		fmt.Sprintf("Дистанция составила %.2f км.\n", distance) +
		fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories)

	return result, nil
}
