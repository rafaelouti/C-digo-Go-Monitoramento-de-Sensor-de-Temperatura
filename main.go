package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Função que simula a leitura de um sensor de temperatura
func lerTemperatura() float64 {
	// Simula uma leitura aleatória entre 20 e 120 graus Celsius
	return 20 + rand.Float64()*(120-20)
}

// Função que verifica se a temperatura está dentro do limite e gera alerta se necessário
func verificarTemperatura(temp float64, limite float64) {
	if temp > limite {
		fmt.Printf("ALERTA: Temperatura alta! Leitura: %.2f°C às %s\n", temp, time.Now().Format("15:04:05"))
	} else {
		fmt.Printf("Temperatura estável: %.2f°C às %s\n", temp, time.Now().Format("15:04:05"))
	}
}

func main() {
	// Definindo o limite máximo de temperatura aceitável
	limiteTemperatura := 80.0

	// Simula leituras do sensor a cada 2 segundos
	for {
		// Leitura simulada do sensor
		temperatura := lerTemperatura()

		// Verificação da leitura e alerta, se necessário
		verificarTemperatura(temperatura, limiteTemperatura)

		// Espera 2 segundos antes de fazer a próxima leitura
		time.Sleep(2 * time.Second)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
