# Maioridade App (Go)

Este projeto foi desenvolvido para o desafio de consolidar conhecimentos sobre criação de pipelines de CI utilizando GitHub Actions e SonarCloud.

---

## 📝 Descrição

O **Maioridade App** é uma aplicação simples desenvolvida em Go que verifica se um usuário é maior de idade com base em sua idade informada. O projeto também foi configurado para executar testes unitários, analisar o código com SonarCloud e criar uma pipeline CI/CD utilizando GitHub Actions.

---

## 📋 Requisitos do Desafio

1. Criar uma aplicação simples.
2. Desenvolver testes unitários para a aplicação.
3. Configurar uma pipeline CI com os seguintes passos:
   - Instalar a aplicação.
   - Executar os testes.
   - Rodar a análise de qualidade com o SonarCloud.
4. Configurar o repositório no GitHub para:
   - Exigir aprovação do status check antes de permitir o merge.
   - Bloquear merges se a aplicação não passar na pipeline de CI ou no quality gate do SonarCloud.
5. Criar uma Pull Request demonstrando todo o processo.

---

## 🚀 Tecnologias Utilizadas

- **Linguagem:** Go
- **Pipeline CI/CD:** GitHub Actions
- **Análise de Qualidade:** SonarCloud
- **Containerização:** Docker

---

## ⚙️ Arquitetura do Projeto

### Estrutura de Arquivos

```plaintext
.
├── main.go             # Arquivo principal da aplicação
├── main_test.go        # Testes unitários
├── Dockerfile          # Configuração para criar a imagem Docker
├── coverage.out        # Arquivo gerado com cobertura de testes
├── sonar-project.properties # Configuração do SonarCloud
└── .github/workflows/maioridade-app-go-ci.yml # Configuração do GitHub Actions
```

---

## 🛠️ Configuração e Execução

### Pré-requisitos

- Go instalado na versão **1.23.2**
- Conta no [SonarCloud](https://sonarcloud.io/)
- Docker instalado para criar e rodar o container
- Repositório no GitHub configurado com:
  - `SONAR_TOKEN` nos **Secrets**
  - `DOCKERHUB_USERNAME` e `DOCKERHUB_TOKEN` para o Docker Hub

### Como Rodar Localmente

1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/maioridade-app-go-ci.git
   cd maioridade-app-go-ci
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Execute os testes:
   ```bash
   go test --coverprofile=coverage.out
   ```

4. Rode a aplicação:
   ```bash
   go run main.go
   ```

5. (Opcional) Gere a imagem Docker:
   ```bash
   docker build -t maioridade-app-go-ci .
   docker run -it maioridade-app-go-ci
   ```

---

## 🌟 Pipeline CI/CD

### Passos Automatizados

1. **Instalar dependências:** Configura o ambiente Go e baixa as dependências.
2. **Rodar os testes:** Executa os testes unitários e gera o relatório de cobertura.
3. **SonarCloud:** Realiza análise de qualidade do código.
4. **Build Docker:** Cria a imagem do aplicativo e faz o push para o Docker Hub.

### GitHub Actions

O arquivo [maioridade-app-go-ci.yml](.github/workflows/maioridade-app-go-ci.yml) define o pipeline completo.

---

## ⚠️ Observação sobre o SonarCloud

- Por estar utilizando a versão gratuita do SonarCloud, as análises de Pull Requests em branches diferentes da `master` não são realizadas.
- No entanto, em uma versão auto-hospedada do SonarQube, seria possível verificar o status da análise diretamente em todas as branches antes do merge.

---

## 🧪 Testes Unitários

Os testes foram implementados para verificar:

- Função de validação de maioridade (`ehMaiorDeIdade`).
- Entrada do usuário (`obterIdade`).
- Carregamento de variáveis de ambiente.
- Execução completa da aplicação com entradas válidas e inválidas.

Para rodar os testes:
```bash
go test ./... -v
```

---

## 📦 Deploy Docker

A imagem da aplicação é enviada automaticamente para o Docker Hub ao final da pipeline CI.

Docker Hub: [marceloscientist1983/maioridade-app-go-ci:latest](https://hub.docker.com/r/marceloscientist1983/maioridade-app-go-ci)

Para rodar a aplicação via Docker:
```bash
docker pull marceloscientist1983/maioridade-app-go-ci:latest
docker run -it marceloscientist1983/maioridade-app-go-ci
```

---

## 📌 Autor

Desenvolvido por **Marcelo Scientist** 👨‍💻  
[GitHub](https://github.com/marceloscientist) | [LinkedIn](https://linkedin.com/in/marceloscientist)
