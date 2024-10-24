# Monitoramento de Sensor de Temperatura

Este projeto é uma aplicação simples em Go que simula o monitoramento de um sensor de temperatura. O programa lê a temperatura a cada intervalo de tempo e gera alertas se a temperatura ultrapassar um limite definido.

## Estrutura do Projeto

- **main.go**: Contém a lógica do monitoramento de temperatura e simulação da leitura do sensor.

## Requisitos

Para executar o projeto, você precisará:

- [Golang](https://golang.org/dl/) instalado na sua máquina.
- Editor de texto ou IDE, como [VSCode](https://code.visualstudio.com/) ou [GoLand](https://www.jetbrains.com/go/).

## Como Executar o Projeto

Siga os passos abaixo para rodar o projeto:

1. **Clonar o Repositório**:
   ```bash
   git clone https://github.com/seu-usuario/monitoramento-temperatura.git
   cd monitoramento-temperatura
go mod init monitoramento-temperatura
go run main.go


Iniciando monitoramento de temperatura...
Temperatura atual: 24.75°C
Temperatura atual: 27.43°C
Temperatura atual: 29.12°C
Alerta: Temperatura muito alta!
Temperatura atual: 23.58°C
Temperatura atual: 28.75°C
Alerta: Temperatura muito alta!
