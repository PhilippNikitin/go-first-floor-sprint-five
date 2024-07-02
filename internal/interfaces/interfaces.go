package interfaces

import (
	"fmt"
	"time"

	"github.com/Yandex-Practicum/go-1fl-homework-sprint5/internal/structs"
)

// CaloriesCalculator интерфейс для структур: Running, Walking и Swimming.
type CaloriesCalculator interface {
	Calories() float64                 // возвращает количество ккал, израсходованных во время тренировки
	TrainingInfo() structs.InfoMessage // возвращает переменную структуры InfoMessage{}
}

// ReadData возвращает информацию о проведенной тренировке.
func ReadData(training CaloriesCalculator) string {
	// вычисляем количество затраченных калорий
	calories := training.Calories()
	// получаем информацию о тренировке в виде структуры InfoMessage
	info := training.TrainingInfo()

	if (calories == 0) || (info.Duration == 0*time.Microsecond) { // если расход калорий равен 0, или продолжительность тренировки равна 0, то на этапе вычислений возникла ошибка, не выводим ничего
		return ""
	}

	// если ошибок не было, добавляем затраченные ккал в структуру с информацией о тренировке и возвращаем данную информацию
	info.Calories = calories

	return fmt.Sprint(info)
}
