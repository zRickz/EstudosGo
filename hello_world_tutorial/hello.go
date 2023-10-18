// Código feito baseado nos códigos dos tutoriais de GoLang (Disponível nas documentação Golang)

// Funcionamento: Hello world, Boas vindas com utilização de demonstração de construção de módulo e
// demonstração de uso de bibliotecas públicas (neste caso o quote).

package main

import (
	"fmt"
	"log"

	"example/greetings"

	"rsc.io/quote"
)

func testQuote() {
	fmt.Println("Mensagem Quote: " + quote.Go())
}

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Henrique")

	if err != nil {
		log.Fatal("Ocorreu um erro!\n" + err.Error())
	}

	r_msg, r_msg_err := greetings.RandomMessage()

	if r_msg_err != nil {
		log.Fatal("Ocorreu um erro!\n" + r_msg_err.Error())
	}

	names := []string{"Henrique", "João", "Maria"}
	messages, msgs_err := greetings.Hellos(names)

	if msgs_err != nil {
		log.Fatal("Ocorreu um erro!\n" + msgs_err.Error())
	}

	fmt.Println("Mensagem boas-vindas: " + message)
	for _, name := range messages {
		fmt.Println("Mensagens boas-vindas: " + name)
	}
	fmt.Println("Mensagem aleatória: " + r_msg)
	fmt.Println("Hello World: Olá mundo! Este é um programa escrito em Golang.")
	testQuote()
}
