# 🎫 Sistema de Venda de Ingressos em Larga Escala

Este projeto é uma aplicação escalável e eficiente para gerenciar a venda de ingressos em larga escala, utilizando uma arquitetura de microserviços e várias tecnologias, como Golang, Java, e serviços da AWS.

## 🚀 Tecnologias Utilizadas

- **Golang** com [Gin Gonic](https://github.com/gin-gonic/gin): Framework web para criar APIs de forma rápida e eficiente.
- **Java** com [Spring Boot](https://spring.io/projects/spring-boot): Framework para facilitar o desenvolvimento de microserviços robustos.
- **AWS Lambda**: Serviço serverless para executar o código em resposta a eventos.
- **AWS SQS (Simple Queue Service)**: Serviço de filas de mensagens, utilizado para gerenciar as tarefas de forma assíncrona.
- **AWS EventBridge**: Serviço de roteamento de eventos que conecta as diferentes partes do sistema.
- **Docker**: Para containerização dos microserviços e garantir que o ambiente seja consistente.

## 🎯 Objetivo

O objetivo desta aplicação é resolver os desafios de vender ingressos em larga escala, garantindo alta escalabilidade, processamento eficiente, e uma arquitetura modular, baseada em microserviços. A integração com a AWS proporciona flexibilidade e escalabilidade, permitindo o processamento de grandes volumes de eventos e mensagens.

## ⚙️ Arquitetura

A aplicação segue a arquitetura de microserviços, dividindo as funcionalidades em pequenos serviços independentes, que se comunicam de forma assíncrona usando AWS SQS e EventBridge. Cada microserviço desempenha uma função específica no sistema, como a gestão de usuários, processamento de pagamentos e envio de confirmações de ingressos.

### Componentes:

1. **Gateway (Java + Spring Boot)**
   - Gerencia o processamento de autenticação.
   - Faz a integração com os outros microserviços.

2. **Serviço de Gerenciamento de Ingressos (Golang + Gin Gonic)**
   - Responsável por lidar com as solicitações de compra de ingressos.
   - Gerencia o inventário de ingressos em tempo real.
   - Usa AWS SQS para processar ordens de compra de forma assíncrona.

3. **Serviço de Gerenciamento de Eventos (Golang + Gin Gonic)**
   - Responsável por Gerenciar os eventos.
   - Faz o debito do ingresso para um determinado evento.
   - Usa AWS SQS para processar o desconto do ingresso forma assíncrona.

4. **AWS Lambda**
   - Executa funções serverless para tarefas críticas e baseadas em eventos.
   - Exemplo: Notificação ao usuário após a confirmação do pagamento.

5. **AWS EventBridge**
   - Gerencia eventos disparados em todo o sistema, garantindo que as mudanças de estado (ex: compra confirmada) sejam propagadas corretamente para os serviços relevantes.

6. **AWS SQS**
   - Assegura que todas as ordens e pagamentos sejam processados de maneira assíncrona, com alta resiliência e escalabilidade.

## 🏗️ Instalação e Configuração

Siga os passos abaixo para rodar o projeto localmente:

### Pré-requisitos

- **Golang**
- **Java**
- **Docker** e **Docker Compose**
- Conta **AWS** para configurar serviços como SQS, Lambda, e EventBridge.
- **LocalStack** (opcional, para simular serviços AWS localmente)

### 1. Clonar o Repositório

### 2. Rodar o docker compose

### 3. Startar as aplicações

## 📜 Documentação
Acesse http://localhost:8080/swagger-ui/index.html#/