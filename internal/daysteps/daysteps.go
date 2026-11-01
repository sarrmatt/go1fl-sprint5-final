package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sarrmatt/tracker/internal/personaldata"
	"github.com/sarrmatt/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку формата "<шаги>,<длительность>",
// например "792,30m", и заполняет поля Steps и Duration.
// Возвращает ошибку при неверном формате или неположительных значениях.
func (ds *DaySteps) Parse(datastring string) (err error) {
	dataSplit := strings.Split(datastring, ",")

	if len(dataSplit) != 2 {
		return fmt.Errorf("неверный формат %q: ожидалось 2 значения, получено %d",
			datastring, len(dataSplit))
	}

	steps, err := strconv.Atoi(dataSplit[0])
	if err != nil {
		return fmt.Errorf("неверный формат количества шагов %q: %w", dataSplit[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным, получено %d", steps)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(dataSplit[1])
	if err != nil {
		return fmt.Errorf("неверный формат времени выполнения %q: %w", dataSplit[1], err)
	}

	if duration <= 0 {
		return fmt.Errorf("длительность должна быть положительной, получена %v", duration)
	}
	ds.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о прогулке:
// количество шагов, дистанция и потраченные калории.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("расчёт калорий: %w", err)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	), nil
}
