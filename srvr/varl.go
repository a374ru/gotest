package srvr

import (
	"fmt"
	gor "github.com/gorilla/mux" // Новая зависимость
	"math/rand"
	"time"
)

func Varr(per string) string {
	var Ttt = time.Now().UTC()
	const message2 = "Hello, Go!"
	const message3 string = "Hello, Programming!"

	var message string
	message = "Hello, World!"

	fmt.Println(message)

	mass := []int{12, 12} // инициализация массива из 10 элементов типа int д

	for i := range mass {
		mass[i] = -(rand.Intn(100)) // заполняем массив случайными числами от 0 до 1000
	}
	mass3 := []int{14, 24, 43, 44, 45}
	fmt.Println(mass, Ttt)

	// mass2 := []int{} // создаем новый массив для объединения
	mass = append(mass3, mass3...)
	asd := gor.NewRouter()
	fmt.Println(mass, per, asd)

	return per
}
