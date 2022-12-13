# api-rest-golang

Este projeto contém uma api REST em Golang, onde o principal serviço é um CRUD de usuários. Para salvar os dados, o banco de Dados escolhido foi o MongoDB. Como forma de facilitar a manutenibilidade e manter as boas práticas, utiliza-se clean architeture e o design pattern Adapter.

## **Modelos das entidades**

### Amosta

### Usuário

```
nome: <string>
email: <string>
senha: <string>
```

## **Funcionalidade de cada entidade**

### Usuário

- Inserção;
- Busca;
- Busca por id;
- Atualização;
- Exclusão.

## **Rotas**

## Usuário

- Inserção

```
url: /user

method: POST

header: {
    "Content-Type": "application/json"
}

body: {
    "nome": <string>,
    "email": <string>,
    "senha": <string>
}
```

- Busca

```
url: /user

method: GET

header: {
    "Content-Type": "application/json"
}
```

- Busca por Id

```
url: /cidade/:id

params: {
    id: <id_do_usuario>
}

method: GET

header: {
    "Content-Type": "application/json"
}
```
- Atualização

```
url: /user/:id

params: {
    id: <id_do_usuario>
}

method: PUT

header: {
    "Content-Type": "application/json"
}

body: {
    "nome": <string>,
    "email": <string>,
    "senha": <string>
}
```

- Exclusão

```
url: /user/:id

params: {
    id: <id_do_usuario>
}

method: DELETE

header: {
    "Content-Type": "application/json"
}

```

## **Requisitos para execução do projeto**

Deve ter:
- ***Docker*** e ***docker-compose*** instalado na máquina.

### Para executar

- ```git clone <url_repositorio>``` : clonar o repositório;
- ```docker-compose up```: rodar a aplicação

Para acessar a API diretamente é preciso acessar ```http://localhost:9999``` + o endPoint.

## TO DO
- Testes unitários
- Middleware de log
- Add serviço de autenticação

