package structs

import (
	"fmt"
	"math"
	"time"

	"github.com/Yandex-Practicum/go-1fl-homework-sprint5/internal/constants"
	"github.com/Yandex-Practicum/go-1fl-homework-sprint5/internal/errs"
)

// Training общая структура для всех тренировок
type Training struct {
	TrainingType    string        // тип тренировки (бег, ходьба или плавание)
	Action          int           // количество повторов(шаги, гребки при плавании)
	LenStep, Weight float64       // длина одного шага или гребка в м; вес пользователя в кг
	Duration        time.Duration // продолжительность тренировки
}

// Метод distance структуры Training возвращает дистанцию в км, которую преодолел пользователь.
// Формула расчета:
// количество_повторов * длина_шага / м_в_км
func (t Training) distance() float64 {
	return float64(t.Action) * t.LenStep / constants.MInKm // MInKm - константа, которая не равна 0 по умолчанию, поэтому проверку на равенство 0 не добавлял
}

// Метод meanSpeed структуры Training возвращает среднюю скорость бега или ходьбы и ошибку, если такая возникла при вычислениях
// Формула расчета средней скорости:
// преодолённая_дистанция_за_тренировку_в_км / время_тренировки_в_часах
func (t Training) meanSpeed() (float64, error) {
	var result float64
	if t.Duration.Hours() == 0. { // проверяем равенство делителя (времени тренировки) нулю
		return result, errs.ErrZeroDuration // если делитель равен 0, возвращаем 0. и ошибку
	}
	result = t.distance() / t.Duration.Hours()
	return result, nil
}

// Метод Calories структуры Training возвращает количество потраченных ккал на тренировке.
// Возвращает 0., данный метод будет переопределен для каждого типа тренировки.
func (t Training) Calories() float64 {
	return 0.
}

// Метод TrainingInfo структуры Training возвращает структуру InfoMessage, в которой хранится информация о проведенной тренировке.
func (t Training) TrainingInfo() InfoMessage {
	mnSpd, err := t.meanSpeed()
	if err != nil { // если при расчете средней скорости возникла ошибка, возвращаем пустую структуру InfoMessage{}
		return InfoMessage{}
	}
	return InfoMessage{
		TrainingType: t.TrainingType,
		Duration:     t.Duration,
		Distance:     t.distance(),
		Speed:        mnSpd,
		Calories:     t.Calories(),
	}
}

// Структура InfoMessage содержит информацию о проведенной тренировке.
type InfoMessage struct {
	TrainingType              string        // тип тренировки
	Duration                  time.Duration // длительность тренировки
	Distance, Speed, Calories float64       // расстояние, преодоленное пользователем; сред.скорость, с которой двигался пользователь; количество потраченных ккал на тренировке
}

// Метод String структуры InfoMessage возвращает строку с информацией о проведенной тренировке.
func (i InfoMessage) String() string {
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %v мин\nДистанция: %.2f км.\nСр. скорость: %.2f км/ч\nПотрачено ккал: %.2f\n",
		i.TrainingType,
		i.Duration.Minutes(),
		i.Distance,
		i.Speed,
		i.Calories,
	)
}

// Running структура, описывающая тренировку Бег.
type Running struct {
	Training
}

// Метод Calories структуры Running возвращает количество потраченных ккал при беге.
// Формула расчета:
// ((18 * средняя_скорость_в_км/ч + 1.79) * вес_спортсмена_в_кг / м_в_км * время_тренировки_в_часах * мин_в_часе)
// Это переопределенный метод Calories() из Training.
func (r Running) Calories() float64 {
	mnSpd, err := r.Training.meanSpeed() // определяем среднюю скорость, пользуясь методом встроенной структуры
	if err != nil {                      // если при определении средней скорости возникла ошибка, выводим ее в stdout
		fmt.Println(err)
		return 0.
	}
	// MInKm в формуле ниже - ненулевая константа, поэтому не проверяем ее на равенство 0
	return ((constants.CaloriesMeanSpeedMultiplier*mnSpd + constants.CaloriesMeanSpeedShift) * r.Training.Weight / constants.MInKm * r.Training.Duration.Hours() * constants.MinInHours)
}

// Метод TrainingInfo структуры Running возвращает структуру InfoMessage с информацией о проведенной тренировке.
// Это переопределенный метод TrainingInfo() из Training.
func (r Running) TrainingInfo() InfoMessage {
	return r.Training.TrainingInfo()
}

// Walking структура описывающая тренировку Ходьба
type Walking struct {
	Training
	Height float64 // рост пользователя в см
}

// Метод Calories структуры Walking возвращает количество потраченных килокалорий при ходьбе.
// Формула расчета:
// ((0.035 * вес_спортсмена_в_кг + (средняя_скорость_в_метрах_в_секунду**2 / рост_в_метрах)
// * 0.029 * вес_спортсмена_в_кг) * время_тренировки_в_часах * мин_в_ч)
// Это переопределенный метод Calories() из Training.
func (w Walking) Calories() float64 {
	mnSpd, err := w.Training.meanSpeed() // определяем среднюю скорость, пользуясь методом встроенной структуры
	if err != nil {                      // если при определении средней скорости возникла ошибка, выводим ее в stdout
		fmt.Println(err)
		return 0.
	}
	if w.Height == 0 { // проверяем входные данные по росту пользователя на равенство 0
		fmt.Println(errs.ErrZeroHeight)
		return 0.
	}

	return ((constants.CaloriesWeightMultiplier*w.Training.Weight + (math.Pow(constants.KmHInMsec*mnSpd, 2)/(w.Height/constants.CmInM))*constants.CaloriesSpeedHeightMultiplier*w.Training.Weight) * w.Training.Duration.Hours() * constants.MinInHours)
}

// Метод TrainingInfo структуры Walking возвращает структуру InfoMessage с информацией о проведенной тренировке.
// Это переопределенный метод TrainingInfo() из Training.
func (w Walking) TrainingInfo() InfoMessage {
	return w.Training.TrainingInfo()
}

// Swimming структура, описывающая тренировку Плавание
type Swimming struct {
	Training
	LengthPool, CountPool int // длина бассейна в метрах; количество пересечений бассейна в метрах
}

// Метод meanSpeed структуры Swimming возвращает среднюю скорость при плавании.
// Формула расчета:
// длина_бассейна в м * количество_пересечений в м / м_в_км / продолжительность_тренировки
func (s Swimming) meanSpeed() (float64, error) {
	var result float64
	if s.Training.Duration.Hours() == 0. { // проверяем равенство делителя (время тренировки) нулю
		return result, errs.ErrZeroDuration // если делитель равен 0, возвращаем 0. и ошибку ErrZeroDuration
	}
	result = float64(s.LengthPool*s.CountPool) / constants.MInKm / s.Training.Duration.Hours()
	return result, nil
}

// Метод Calories структуры Swimming возвращает количество калорий, потраченных при плавании.
// Формула расчета:
// (средняя_скорость_в_км/ч + SwimmingCaloriesMeanSpeedShift) * SwimmingCaloriesWeightMultiplier * вес_спортсмена_в_кг * время_тренировки_в_часах
// Это переопределенный метод Calories() из Training.
func (s Swimming) Calories() float64 {
	mnSpd, err := s.meanSpeed()
	if err != nil {
		fmt.Println(err)
		return 0.
	}

	return (mnSpd + constants.SwimmingCaloriesMeanSpeedShift) * constants.SwimmingCaloriesWeightMultiplier * s.Training.Weight * s.Training.Duration.Hours()
}

// Метод TrainingInfo структуры Swimming возвращает информацию о тренировке "плавание".
// Это переопределенный метод TrainingInfo() из Training.
// для расчета преодоленной дистанции при плавании использована следующая формула:
// длина бассейна * количество пересечений бассейна / количество метров в одном километре
// также для расчета дистанции при плавании может использоваться формула:
// длина одного гребка * количество гребков / количество метров в одном километре
func (s Swimming) TrainingInfo() InfoMessage {
	mnSpd, err := s.meanSpeed()
	if err != nil { // если при расчете средней скорости произошла ошибка, возвращаем пустую структуру InfoMessage{}
		return InfoMessage{}
	}

	return InfoMessage{
		TrainingType: s.Training.TrainingType,
		Duration:     s.Training.Duration,
		Distance:     float64(s.LengthPool*s.CountPool) / constants.MInKm,
		Speed:        mnSpd,
		Calories:     s.Calories(),
	}
}
