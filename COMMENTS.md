# BALLOT BOX

Serviço para apuração de votos do paredão do Big Brother

Para iniciar o serviço execute: ***docker compose up***

## Estrutura

Serviço desenvolvido em Golang acessando os bancos de dados Redis e MongoDB

O MongoDB foi escolhido por ser um banco NoSQL e pensando principalmente na facilidade
para criar um shard em tabelas, já que a inserção no banco pode ser um gargalo

A escolha do Redis é porque ele é extremamente performático e se encaixa muito bem no contexto do
serviço

## Arquitetura

A idéia principal é manter o serviço saudável mesmo sob grande quantidade de requisições e evitar votos de robôs.

Para isso implementei um sequência de middlewares que vou detalhar abaixo:

### BackPressure Middleware

O excesso de requisições pode causar uma sobrecarga em todo o sistema, para evitar essa sobrecarga o BackPressure 
Middleware vai retornar um erro para as requisições que ultrapassarem o limite definido dentro do período de
tempo definido.

Ele é o primeiro middleware a ser executado.

### BloomFilter Middleware

O algoritmo de bloom filter é a forma mais eficaz para verificar se um usuário está em uma lista de usuários 
bloqueados (são inseridos nessa lista no middleware de rating).

### BasicValidation Middleware

Esse middleware faz a validação se a requisição tem os cabeçalhos esperados para dessa forma tentar diminuir
votos automatizados.

### Rating Middleware

O Rating Middleware confere de forma individual se cada usuário está enviando uma quantidade máxima de requisições 
dentro do período estabelecido, por exemplo, um usuários pode enviar uma requisição por segundo.

Se o usuário ultrapassar esse limite, ele serpe inserido no BloomFilter e sua próxima requisição retornará um erro.

## Envio da requisição para o stream

Após a requisição passar por todos os middlewares ela será colocada em duas streams do Redis,
uma stream é para serviços OLTP e outra stream é para serviços OLAP.

### Consumidores

Existe um consumidor para a stream de OLTP e um consumidor para a strem de OLAP

Os consumidores fazem parte de um groupreader, dessa forma, mais de um consumidor pode receber as mensagens do 
mesmo stream

#### OLTP

Aqui os dados são enviados para uam collection no MongoDB

#### OLAP

Aqui os dados são enviados para uam collection no MongoDB

Mas em um cenário real poderia ser enviado para um ClickHouse ou GCP BigQuery por exemplo


## Melhorias

Aqui vou elencar quais melhorias eu faria se tivesse mais tempo.

Adicionar um dead letter queue para tentativas que deram erro no consumer

Log com opentelemetry

Adicionar os registros no MongoDB utilizando insertMany, implicaria em mudar um pouco a lógica
atual de ack

Adicionar o retorna da média e outras informações necessárias


# Observações

Vou deixar dois vídeos no diretório ***docs*** com o teste de performance e
também os testes com o BloomFilter e RatingLimit

A performance foi testada utilizando JMeter, vou deixar o arquivo de teste na 
pasta ***docs*** também


# CURLS

Endpoint de health para ser utilizado com K8s

curl --location 'http://localhost:5000/health'

Endpoint e votação

curl --location 'http://localhost:5000/voting' \
--header 'user: 11' \
--header 'accept;' \
--header 'accept-encoding;' \
--header 'accept-language;' \
--header 'cookie;' \
--header 'referer;' \
--header 'user-agent;' \
--header 'sec-ch-ua;' \
--header 'sec-ch-ua-mobile;' \
--header 'sec-ch-ua-platform;' \
--header 'Content-Type: application/json' \
--data '{
"wall": 1,
"candidate": 10
}'



