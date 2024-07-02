package main

import (
	"fmt"
	"time"

	"github.com/Yandex-Practicum/go-1fl-homework-sprint5/internal/constants"
	"github.com/Yandex-Practicum/go-1fl-homework-sprint5/internal/interfaces"
	"github.com/Yandex-Practicum/go-1fl-homework-sprint5/internal/structs"
)

func main() {

	swimming := structs.Swimming{
		Training: structs.Training{
			TrainingType: "Плавание",
			Action:       2000,
			LenStep:      constants.SwimmingLenStep,
			Duration:     90 * time.Minute,
			Weight:       85,
		},
		LengthPool: 50,
		CountPool:  5,
	}

	fmt.Println(interfaces.ReadData(swimming))

	walking := structs.Walking{
		Training: structs.Training{
			TrainingType: "Ходьба",
			Action:       20000,
			LenStep:      constants.LenStep,
			Duration:     3*time.Hour + 45*time.Minute,
			Weight:       85,
		},
		Height: 185,
	}

	fmt.Println(interfaces.ReadData(walking))

	running := structs.Running{
		Training: structs.Training{
			TrainingType: "Бег",
			Action:       5000,
			LenStep:      constants.LenStep,
			Duration:     30 * time.Minute,
			Weight:       85,
		},
	}

	fmt.Println(interfaces.ReadData(running))

}
