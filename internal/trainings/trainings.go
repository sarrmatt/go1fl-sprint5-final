package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sarrmatt/tracker/internal/personaldata"
	"github.com/sarrmatt/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "<шаги>,<активность>,<длительность>",
// например "3456,Ходьба,3h00m", на количество шагов, вид активности
// и time.Duration.
func (t *Training) Parse(datastring string) (err error) {
	dataSplit := strings.Split(datastring, ",")

	if len(dataSplit) != 3 {
		return fmt.Errorf("неверный формат %q: ожидалось 3 значения, получено %d",
			datastring, len(dataSplit))
	}

	steps, err := strconv.Atoi(dataSplit[0])
	if err != nil {
		return fmt.Errorf("неверный формат количества шагов %q: %w", dataSplit[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным, получено %d", steps)
	}
	t.Steps = steps

	t.TrainingType = dataSplit[1]

	duration, err := time.ParseDuration(dataSplit[2])
	if err != nil {
		return fmt.Errorf("неверный формат времени выполнения %q: %w", dataSplit[2], err)
	}
	if duration <= 0 {
		return fmt.Errorf("длительность должна быть положительной, получена %v", duration)
	}
	t.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о тренировке:
// тип, длительность, дистанция, скорость и потраченные калории.
func (t Training) ActionInfo() (string, error) {
	var calories float64
	var err error

	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %q", t.TrainingType)
	}
	if err != nil {
		return "", fmt.Errorf("расчёт калорий: %w", err)
	}

	return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, speed, calories,
		),
		nil
}
