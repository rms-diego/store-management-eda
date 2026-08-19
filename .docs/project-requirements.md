# EDA Store Challenge

Desafio técnico com o objetivo de estudar e praticar **Event-Driven Architecture (EDA)** utilizando:

- Go
- RabbitMQ
- PostgreSQL
- Monorepo
- Docker Compose

A aplicação representa um fluxo simplificado de compra dividido em três serviços principais:

- **Order Service**
- **Inventory Service**
- **Billing Service**

Além deles, haverá um **API Gateway** como ponto único de comunicação entre o client e os serviços.

## Arquitetura

```text
                    ┌──────────────┐
                    │    Client    │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │ API Gateway  │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │    Order     │
                    └──────┬───────┘
                           │
                           ▼
                     ┌──────────┐
                     │ RabbitMQ │
                     └────┬─────┘
                          │
                ┌─────────┴─────────┐
                ▼                   ▼
        ┌──────────────┐     ┌──────────────┐
        │  Inventory   │     │   Billing    │
        └──────────────┘     └──────────────┘
```

O objetivo é evitar que os serviços dependam diretamente uns dos outros. A comunicação entre eles deverá acontecer principalmente através de eventos.

## Serviços

### Order

Responsável por:

- criar pedidos;
- armazenar os itens;
- acompanhar o status do pedido;
- reagir aos eventos de estoque e pagamento.

### Inventory

Responsável por:

- manter o estoque dos produtos;
- reservar estoque para pedidos;
- rejeitar pedidos sem estoque suficiente;
- liberar reservas quando necessário.

### Billing

Responsável por:

- processar pagamentos;
- simular pagamentos aprovados ou recusados;
- informar o resultado através de eventos.

## Fluxo principal

O fluxo esperado é:

```text
Order criado
    ↓
order.created
    ↓
Inventory reserva estoque
    ↓
inventory.reserved
    ↓
Billing processa pagamento
    ↓
payment.succeeded
    ↓
Order confirmado
```

Em caso de estoque insuficiente:

```text
Order criado
    ↓
inventory.reservation_failed
    ↓
Order cancelado
```

Em caso de falha no pagamento:

```text
Order criado
    ↓
Inventory reservado
    ↓
payment.failed
    ↓
Inventory libera estoque
    ↓
Order cancelado
```

Esse fluxo deve servir como exercício de uma **Saga baseada em choreography**.

## Eventos

Alguns eventos esperados:

```text
order.created

inventory.reserved
inventory.reservation_failed
inventory.released

payment.succeeded
payment.failed

order.confirmed
order.cancelled
```

Os eventos devem possuir pelo menos:

```json
{
  "event_id": "uuid",
  "event_type": "order.created",
  "occurred_at": "2026-08-19T20:00:00Z",
  "payload": {}
}
```

## Banco de dados

Cada serviço deve ser responsável pelos seus próprios dados.

```text
Order
    └── orders

Inventory
    └── products
    └── inventory
    └── reservations

Billing
    └── payments
```

Um serviço não deve acessar diretamente as tabelas pertencentes a outro serviço.

## Estrutura do monorepo

Uma possível estrutura:

```text
.
├── cmd
│   ├── gateway
│   ├── order
│   ├── inventory
│   └── billing
│
├── internal
├── pkg
├── migrations
├── go.work
└── README.md
```

A organização interna de cada serviço fica a critério da implementação.

## Requisitos

- [ ] Criar pedidos através do API Gateway.
- [ ] Processar o fluxo de compra utilizando RabbitMQ.
- [ ] Manter os dados de cada serviço isolados.
- [ ] Reservar estoque de maneira segura.
- [ ] Processar pagamentos de forma assíncrona.
- [ ] Cancelar pedidos quando estoque ou pagamento falharem.
- [ ] Liberar estoque quando um pagamento falhar.
- [ ] Garantir que mensagens duplicadas não produzam efeitos duplicados.
- [ ] Implementar Transactional Outbox.
- [ ] Implementar retry e Dead Letter Queue.
- [ ] Executar toda a infraestrutura com Docker Compose.

## Cenários que devem funcionar

### Happy path

```text
Order
 → Inventory
 → Billing
 → CONFIRMED
```

### Estoque insuficiente

```text
Order
 → Inventory Failed
 → CANCELLED
```

### Pagamento recusado

```text
Order
 → Inventory Reserved
 → Payment Failed
 → Inventory Released
 → CANCELLED
```

### Mensagem duplicada

O mesmo evento pode ser entregue mais de uma vez sem gerar efeitos colaterais duplicados.

### Falha temporária

Se um consumer ficar indisponível, as mensagens devem continuar disponíveis para processamento quando ele retornar.

### Concorrência

Dois pedidos concorrendo pelo último item disponível não podem fazer o estoque ficar negativo.

## Conceitos a serem exercitados

O principal objetivo do desafio é entender na prática:

- Event-Driven Architecture;
- RabbitMQ;
- exchanges, queues e routing keys;
- producers e consumers;
- consistência eventual;
- Saga / choreography;
- idempotência;
- at-least-once delivery;
- Transactional Outbox;
- retry;
- Dead Letter Queue;
- concorrência;
- isolamento entre serviços.

## Critério de conclusão

O desafio estará concluído quando for possível iniciar o ambiente com:

```bash
docker compose up
```

criar um pedido através do API Gateway e acompanhar todo o fluxo entre **Order → Inventory → Billing**, incluindo os principais cenários de sucesso e falha.

O foco do projeto não é construir um e-commerce completo, mas utilizar um domínio simples para **entender os problemas e trade-offs de uma arquitetura orientada a eventos**.
