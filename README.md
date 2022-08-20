# Digital Bank API

## 1. Installation

Please, install it if you don't have it yet.

Docker: https://docs.docker.com/get-docker/

Docker Compose: https://docs.docker.com/compose/install/

Go to project root and install the dependencies:
```shell
go mod tidy
```

Before starting make sure to stop running containers:
```shell
docker stop $(docker ps -a -q)
```
Now build the application:
```shell
docker-compose up -d --build
```

Once it's done, you can check logs by running:
```shell
docker-compose logs -f
```
Wait until server is up and running. You will see this message:
```
database_1         | LOG:  listening on IPv4 address "0.0.0.0", port 5432
database_1         | LOG:  listening on IPv6 address "::", port 5432
database_1         | LOG:  database system is ready to accept connections
pgadmin-compose_1  | [INFO] Starting gunicorn 20.1.0
pgadmin-compose_1  | [INFO] Listening at: http://[::]:80
```

Open your Postman/Insomnia and paste the url below:
```
http://127.0.0.1:8080/{your-name}
```

=====================
#### Regras gerais

* Usar formato JSON para leitura e escrita. (ex: `GET /accounts/` retorna json, `POST /accounts/ {name: 'james bond'}`)

#### Rotas esperadas

##### `/accounts`

A entidade `Account` possui os seguintes atributos:

* `id`
* `name` 
* `cpf`
* `secret`
* `balance` 
* `created_at` 

Espera-se as seguintes ações:

- `GET /accounts` - obtém a lista de contas
- `GET /accounts/{account_id}/balance` - obtém o saldo da conta
- `POST /accounts` - cria uma `Account`

*Regras para esta rota*

- `balance` pode iniciar com 0 ou algum valor para simplificar
- `secret` deve ser armazenado como hash

* * *
##### `/login`

A entidade `Login` possui os seguintes atributos:

* `cpf`
* `secret`

Espera-se as seguintes ações:

- `POST /login` - autentica a usuaria

*Regras para esta rota*

- Deve retornar token para ser usado nas rotas autenticadas

* * * 

##### `/transfers`

A entidade `Transfer` possui os seguintes atributos:

* `id`
* `account_origin_id`
* `account_destination_id`
* `amount`
* `created_at`

Espera-se as seguintes ações:

- `GET /transfers` - obtém a lista de transferencias da usuaria autenticada.
- `POST /transfers` - faz transferencia de uma `Account` para outra.

*Regras para esta rota*

- Quem fizer a transferência precisa estar autenticada.
- O `account_origin_id` deve ser obtido no Token enviado.
- Caso `Account` de origem não tenha saldo, retornar um código de erro apropriado
- Atualizar o `balance` das contas
