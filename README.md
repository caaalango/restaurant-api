# Projeto de Automação de Processos de Restaurante

## Índice

- [Visão Geral](#visão-geral)
- [Funcionalidades](#funcionalidades)
- [Arquitetura](#arquitetura)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Pré-requisitos](#pré-requisitos)
- [Instalação](#instalação)
- [Como Usar](#como-usar)
- [Próximas Etapas](#próximas-etapas)
- [Licença](#licença)
- [Contato](#contato)

## Visão Geral

Este projeto, desenvolvido em **Golang**, tem como objetivo automatizar os processos de um restaurante, oferecendo uma solução completa que engloba frontend, backend e infraestrutura. O projeto será desenvolvido em etapas e servirá como portfólio, acompanhado de artigos e vídeos explicativos sobre cada funcionalidade implementada.

## Funcionalidades

### Etapa 1: Cardápio Interativo (Concluída)

- **Sistema de Usuários**: Cadastro, login tradicional e social, com gerenciamento de permissões para acesso a endpoints específicos.
- **Recuperação de Senha**: Envio de e-mail para redefinição de senha.
- **Catálogo Virtual Personalizado**: Cada restaurante possui um catálogo exclusivo e customizável.
- **Filtros Avançados**: Auxílio a clientes com restrições alimentares, permitindo substituições e personalizações nos pedidos.
- **Sistema de Comentários**: Usuários autenticados podem avaliar e comentar sobre os pedidos realizados.

## Arquitetura

- **Padrão Arquitetural**: Hexagonal (Ports and Adapters), promovendo uma separação clara entre o domínio e a infraestrutura.
- **Monólito em Monorepo**: Organização centralizada do código, facilitando o gerenciamento e a manutenção.

## Tecnologias Utilizadas

- **Linguagem de Programação**: Golang
- **Banco de Dados**: PostgreSQL
- **Cache**: Redis
- **Filas de Mensagens**: RabbitMQ
- **Autenticação**: Login tradicional e OAuth para login social
- **Comunicação**: APIs RESTful
- **Documentação**: Geração automática via Swagguer
- **Testes**: Testes unitários e de integração
- **Deploy**: Pipeline via GitHub Actions para a AWS

## Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas:

- **Docker** e **Docker Compose**
- **Make**

## Instalação

1. **Clone o Repositório**

   ```bash
   git clone git@github.com:caaalango/restaurant-app.git
   cd restaurant-app
   ```

2. **Configuração Inicial**

   Na primeira vez, execute:

   ```bash
   make setup
   ```

   Este comando irá:

   - Construir as imagens Docker necessárias.
   - Configurar o banco de dados PostgreSQL.
   - Iniciar os serviços de cache e filas de mensagens.

## Como Usar

Para iniciar a aplicação após a configuração inicial, execute:

```bash
make run
```

A aplicação estará disponível em `http://localhost:8080`.

## Próximas Etapas

### Etapa 2: Gerenciador em Tempo Real de Pedidos

- **Tecnologia**: Implementação utilizando GraphQL.
- **Objetivo**: Permitir o gerenciamento de pedidos em tempo real, melhorando a interação e a eficiência do serviço.

### Etapa 3: Isolamento em Microserviços

- **Tecnologia**: Utilização de gRPC para comunicação entre serviços.
- **Objetivo**: Refatorar a aplicação monolítica em microserviços para aumentar a escalabilidade e a modularidade.

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## Contato

Para dúvidas ou sugestões:

- **Email**: eliabebastosdias01@gmail.com
- **LinkedIn**: [Eliabe Bastos](https://www.linkedin.com/in/eliabebastos/)

- **Email**: pinheiro.nomegabriel@gmail.com
- **Linkedin**: [Gabriel Palitot](https://www.linkedin.com/in/gabriel-palitot-3a4b87186/)

- **Website**: Em breve
