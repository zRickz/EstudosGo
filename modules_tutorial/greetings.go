package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Gerar um Olá para um nome
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Por favor, especifique um nome.")
	}
	message := fmt.Sprintf("Olá, %v, boas-vindas!", name)
	return message, nil
}

// Gerar um Olá para vários nomes (Utilizando a função Hello)
func Hellos(names []string) (map[string]string, error) {
	if len(names) == 0 {
		return nil, errors.New("Por favor, especifique pelo menos um nome.")
	}

	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

// Gerar mensagem aleatória
func RandomMessage() (string, error) {
	rand.Seed(time.Now().UnixNano())

	messages := []string{
		"Olá marinheiro!",
		"Boas-vindas senhor(a)!",
		"Como a vida é bela!",
		"Temos que caminhar!",
	}
	randomIndex := rand.Intn(len(messages))
	return messages[randomIndex], nil
}
