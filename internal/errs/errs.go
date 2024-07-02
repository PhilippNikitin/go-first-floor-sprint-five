package errs

import "errors"

// статически создаем ошибки для обработки ошибок во входных данных
var (
	// если во входных данных указано, что длительность тренировки равна 0
	ErrZeroDuration = errors.New("ошибка во входных данных: длительность тренировки указана, как равная 0")
	// если во входных данных указано, что рост пользователя равен 0
	ErrZeroHeight = errors.New("ошибка во входных данных: рост пользователя не может быть равен 0")
)
