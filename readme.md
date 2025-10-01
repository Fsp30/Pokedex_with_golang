# 📖 Pokedex with Golang

## ✨ Introdução

Este projeto é a minha primeira aplicação publicada utilizando Golang. Embora eu já tivesse tido contato prévio com a linguagem, este é o primeiro projeto que realmente coloco em prática e disponibilizo como estudo e aprendizado.

A ideia principal é construir uma Pokedex interativa, onde usuários podem:

- Buscar informações de Pokémons a partir da PokeAPI
- Dar likes em seus Pokémons favoritos
- Registrar opiniões sobre eles

Toda a informação consultada é persistida em um MongoDB, permitindo manter um histórico de interações dos usuários.

O projeto tem como objetivo principal ser um laboratório prático de Golang, integrando boas práticas com ferramentas modernas de desenvolvimento, além de servir como base para evolução futura de aplicações maiores.

## 🛠️ Tecnologias Utilizadas

- **Golang 1.25** → Linguagem principal do projeto
- **Gin** → Framework web mais popular no ecossistema Go
- **MongoDB** → Banco de dados NoSQL utilizado para persistência
- **Docker** → Containerização da aplicação
- **Docker Compose** → Orquestração de containers (aplicação + MongoDB)
- **PokeAPI** → Fonte de dados de Pokémons, acessada pela aplicação

## 📂 Estrutura de Pastas

A aplicação segue uma estrutura modular simples:

```
├── config/ # Configurações gerais (ex: conexão com o banco)
├── handlers/ # Camada que trata as requisições HTTP
├── models/ # Estruturas de dados (ex: User, Pokemon, Opinion)
├── services/ # Camada de acesso ao banco de dados
├── main.go # Ponto de entrada da aplicação
├── go.mod # Configuração de dependências do Go
├── go.sum # Checksum das dependências
└── .gitignore # Documentações auxiliares
```


Essa separação facilita a manutenção e deixa claro o papel de cada parte da aplicação.

## 🚀 Setup do Projeto

### Pré-requisitos

- Go 1.25+
- Docker e Docker Compose

### Clonar o repositório

```bash
  git clone https://github.com/Fsp30/Pokedex_with_golang.git
  cd pokedex_with_golang
```

### Rodar com Docker
```bash
  docker-compose up --build
```

1.Subir MongoDB:
```bash
  docker run -d -p 27017:27017 mongo
```
2.Rodar a aplicação:

```bash
go run main.go
A API estará disponível em: http://localhost:8080
```

---

## 📌 Funcionalidades Atuais

**Usuários**
- Criar usuários no sistema

**Pokémon**
- Buscar Pokémon por nome (consultando a PokeAPI)
- Dar like em um Pokémon
- Registrar opinião de usuário sobre um Pokémon
- Listar todas as opiniões de um Pokémon

---

## 🔍 Exemplos de Uso

`Criar usuário`
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"username": "ash", "email": "ash@pokedex.com"}'
```

`Buscar Pokémon`
```bash
curl -X GET http://localhost:8080/pokemon/pikachu
```

`Dar like`
```bash
curl -X POST http://localhost:8080/pokemon/pikachu/like
```

`Opinar sobre um Pokémon`
```bash
curl -X POST http://localhost:8080/pokemon/pikachu/opinion \
  -H "Content-Type: application/json" \
  -d '{"user_id": "634e0f1d9a5f123456789012", "text": "Pikachu é incrível!"}'
```

`Listar opiniões`
```bash
curl -X GET http://localhost:8080/pokemon/pikachu/opinions
```

--- 

## 🎯 Objetivo do Projeto

O principal objetivo não é apenas construir uma aplicação funcional, mas também:
- Praticar boas práticas em Golang
- Aprender a organizar um projeto real com múltiplas camadas
- Integrar Docker e Docker Compose ao fluxo de desenvolvimento
- Publicar e manter um projeto open source no GitHub

Esse é um primeiro passo para consolidar meu aprendizado em Go e, futuramente, expandir com mais funcionalidades e boas práticas (testes unitários, autenticação, CI/CD, etc.).

---

## 📜 Licença
Este projeto é de uso livre para estudo e aprendizado.


