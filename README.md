# 🏆 Game API

Uma API desenvolvida em **Go** para gerenciar partidas de um jogo de cartas com funcionalidades como controle de rodadas, apostas, cartas jogadas e definição de vencedores.

---

## 📜 Índice

- [Descrição](#descrição)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Funcionalidades](#funcionalidades)
- [Instalação e Configuração](#instalação-e-configuração)
- [Variáveis de Ambiente](#variáveis-de-ambiente)
- [Execução](#execução)
- [Estrutura de Diretórios](#estrutura-de-diretórios)
- [Endpoints da API](#endpoints-da-api)

---

## 📝 Descrição

A aplicação **Game API** gerencia partidas, jogadores e rodadas de um jogo de cartas. Foi desenvolvida para fornecer suporte robusto e escalável, com um foco em performance e manutenibilidade.

---

## 🛠️ Tecnologias Utilizadas

- **Go**: Linguagem principal para o backend.
- **PostgreSQL**: Banco de dados relacional.
- **Supabase**: Gerenciamento de banco de dados e autenticação.
- **Heroku**: Hospedagem da aplicação.
- **REST API**: Estrutura para comunicação entre cliente e servidor.

---

## ✨ Funcionalidades

- **Gerenciamento de partidas**:
  - Criação, atualização e exclusão de partidas.
  - Controle de status (início, andamento, fim).

- **Controle de rodadas**:
  - Atualização automática de rodadas.
  - Definição de turnos e identificação do vencedor.

- **Gerenciamento de jogadores**:
  - Atualização de vidas dos jogadores.
  - Controle de apostas e cartas jogadas.

- **Integração com Supabase**:
  - Utilização de funções RPC para manipulação de dados.

---

## 🚀 Instalação e Configuração

### Pré-requisitos

- **Go** instalado (>= 1.19).
- Banco de dados **PostgreSQL**.
- Conta no **Heroku** (opcional para hospedagem).

### Passo a passo

1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/game-api.git
   cd game-api
   ```

2.	Instale as dependências:
  ```bash
  go mod tidy
  ```

3.	Configure as variáveis de ambiente no Heroku ou em um arquivo .env:
  ```bash
  SUPABASE_URL=<sua-url-supabase>
  SUPABASE_KEY=<sua-chave-api>
  SUPABASE_SERVICE_ROLE=<sua-chave-service-role>
  ```

4.  Execute as migrações do banco de dados (caso necessário).

## 🌍 Variáveis de Ambiente
	•	SUPABASE_URL: URL do seu Supabase.
	•	SUPABASE_KEY: Chave pública de acesso.
	•	SUPABASE_SERVICE_ROLE: Chave com permissões de Service Role.
	•	PORT: Porta para rodar a aplicação (padrão: 3333).

## 🖥️ Execução
Para rodar a aplicação localmente:
  ```bash
  go run main.go
  ```

#### Execução no Heroku
1.	Faça login no Heroku:
  ```bash
  heroku login
  ```

2.	Configure o repositório remoto:
  ```bash
  heroku git:remote -a nome-da-aplicacao
  ```

3.	Faça o deploy:
  ```
  git push heroku main
  ```

## 📁 Estrutura de Diretórios
  ```plaintext
  game-api/
  ├── main.go               # Arquivo principal
  ├── config/               # Configurações gerais
  │   ├── env.go            # Carregamento de variáveis de ambiente
  ├── controllers/          # Lógica para manipulação de dados
  ├── models/               # Estruturas de dados (types)
  ├── repositories/         # Comunicação com o banco de dados
  ├── services/             # Regras de negócio
  └── routes/               # Definição de rotas e handlers
  ```

## 📡 Endpoints da API
	•	GET /update: Atualiza todos os dados do jogo.
	•	GET /trumps: Lista os trunfos da rodada.
	•	POST /deal: Distribui as cartas se o jogador for o dealer.
	•	POST /play: Joga uma carta.
	•	PUT /finish-round: Finaliza uma rodada.
