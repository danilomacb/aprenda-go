package main

import "log"

type erroEspecial struct {
	Message string
}

func (e erroEspecial) Error() string {
	return e.Message
}

func errorHandler(err error) {
	log.Println("Erro:", err)
}

func main() {
	err := erroEspecial{"Mensagem confusa e pouco descritiva"}

	errorHandler(err)
}
