# GoCommerce 🛒

> API RESTful de um e-commerce didático desenvolvido em Go, com foco em boas práticas de arquitetura e clean code.

---

## Sobre o Projeto

Este projeto tem o objetivo de demonstrar minhas habilidades como desenvolvedor backend, criando um sistema de e-commerce completo com:

- CRUD de produtos e categorias
- Carrinho de compras
- Sistema de autenticação (JWT)
- Checkout fake (simulando pagamento)
- Histórico de pedidos
- Admin com painel simples
- API documentada com Swagger
- Deploy online

---

## Tecnologias Utilizadas

- Go + Gin
- PostgreSQL
- Redis (cache)
- TailwindCSS (frontend estático)
- Swagger (docs)
- Docker

---

## Como executar

```bash
# Clonar o repositório
git clone https://github.com/seuusuario/gocommerce.git

# Entrar na pasta
cd gocommerce

# Subir os serviços
docker-compose up --build
```

# Acesse
- API: http://localhost:8080/api
- Swagger: http://localhost:8080/swagger
- Frontend: http://localhost:8080
