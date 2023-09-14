# FreeCatalogo
 Api gratuita para catalogar, livros, filmes, mangás, quadrinhos etc.

Em desenvolvimento.


Categorias
As rotas abaixo permitem que você realize operações CRUD (Create, Read, Update, Delete) em categorias.

Claro! Aqui está a documentação formatada em Markdown:

# Documentação da API de Categorias

Esta é a documentação da API de Categorias do nosso aplicativo. As rotas abaixo permitem que você realize operações CRUD (Create, Read, Update, Delete) em categorias.

## Rotas

### Criação de Categoria

#### `POST /categories`

Cria uma nova categoria.

**Parâmetros de Solicitação**

Nenhum parâmetro é necessário no corpo da solicitação.

**Exemplo de Solicitação**

```http
POST /categories
{
    "name": "Nova Categoria"
}
```

**Resposta de Sucesso**

- Código de Status: 201 Created
- Corpo da Resposta:

```json
{
    "id": 1,
    "name": "Nova Categoria"
}
```

### Recuperação de Categoria por ID

#### `GET /categories/:id`

Recupera uma categoria pelo seu ID.

**Parâmetros de Solicitação**

- `id` (int): O ID da categoria que você deseja recuperar.

**Exemplo de Solicitação**

```http
GET /categories/1
```

**Resposta de Sucesso**

- Código de Status: 200 OK
- Corpo da Resposta:

```json
{
    "id": 1,
    "name": "Nova Categoria"
}
```

### Atualização de Categoria por ID

#### `PUT /categories/:id`

Atualiza uma categoria existente pelo seu ID.

**Parâmetros de Solicitação**

- `id` (int): O ID da categoria que você deseja atualizar.

**Exemplo de Solicitação**

```http
PUT /categories/1
{
    "name": "Categoria Atualizada"
}
```

**Resposta de Sucesso**

- Código de Status: 200 OK
- Corpo da Resposta:

```json
{
    "id": 1,
    "name": "Categoria Atualizada"
}
```

### Exclusão de Categoria por ID

#### `DELETE /categories/:id`

Exclui uma categoria pelo seu ID.

**Parâmetros de Solicitação**

- `id` (int): O ID da categoria que você deseja excluir.

**Exemplo de Solicitação**

```http
DELETE /categories/1
```

**Resposta de Sucesso**

- Código de Status: 204 No Content

### Recuperação de Categoria por Nome

#### `GET /categories/name/:name`

Recupera uma categoria pelo seu nome.

**Parâmetros de Solicitação**

- `name` (string): O nome da categoria que você deseja recuperar.

**Exemplo de Solicitação**

```http
GET /categories/name/Nova%20Categoria
```

**Resposta de Sucesso**

- Código de Status: 200 OK
- Corpo da Resposta:

```json
{
    "id": 1,
    "name": "Nova Categoria"
}
```

### Atualização de Categoria por Nome

#### `PUT /categories/name/:name`

Atualiza uma categoria existente pelo seu nome.

**Parâmetros de Solicitação**

- `name` (string): O nome da categoria que você deseja atualizar.

**Exemplo de Solicitação**

```http
PUT /categories/name/Nova%20Categoria
{
    "name": "Categoria Atualizada"
}
```

**Resposta de Sucesso**

- Código de Status: 200 OK
- Corpo da Resposta:

```json
{
    "id": 1,
    "name": "Categoria Atualizada"
}
```

### Exclusão de Categoria por Nome

#### `DELETE /categories/name/:name`

Exclui uma categoria pelo seu nome.

**Parâmetros de Solicitação**

- `name` (string): O nome da categoria que você deseja excluir.

**Exemplo de Solicitação**

```http
DELETE /categories/name/Nova%20Categoria
```

**Resposta de Sucesso**

- Código de Status: 204 No Content

## Erros

Em caso de erro, a API retornará uma resposta com um código de status apropriado e uma mensagem de erro no corpo da resposta.

Por favor, consulte a documentação específica do aplicativo para obter informações adicionais sobre os modelos de dados, autenticação e outras características específicas do seu sistema.