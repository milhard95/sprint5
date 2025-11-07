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
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	values := strings.Split(datastring, ",")
	if len(values) != 3 {
		return fmt.Errorf("incorrect value received")
	}

	steps, err := strconv.Atoi(values[0])
	if err != nil {
		return err
	}

	if steps <= 0 {
		return fmt.Errorf("the number of steps must be greater than 0")
	}

	duration, err := time.ParseDuration(values[2])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}

	if duration <= 0 {
		return fmt.Errorf("duration must be greater than 0")
	}

	t.Steps = steps
	t.TrainingType = values[1]
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {

	distance := spentenergy.Distance(t.Steps, t.Height)
	MeanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var (
		calories float64
		err      error
	)

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("unknown type of training")
	}

	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Тип тренировки: %s\n", t.TrainingType) +
		fmt.Sprintf("Длительность: %.2f ч.\n", t.Duration.Hours()) +
		fmt.Sprintf("Дистанция: %.2f км.\n", distance) +
		fmt.Sprintf("Скорость: %.2f км/ч\n", MeanSpeed) +
		fmt.Sprintf("Сожгли калорий: %.2f\n", calories)

	return result, nil
}
