// Вопросы:
// 1) Что вообще за `defer`? Почему порядок выполнения именно такой? -- defer откладывает выполнение до конца функции, возвращает по принципу LIFO.
// 2) Что общего у `defer` и стека? -- defer помещает отложенные функции в стек, откуда они потом выполняются по LIFO.
package main

import "fmt"

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}
