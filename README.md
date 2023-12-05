# POC (Proof of Concept) - Kafka com Go

 
Este repositório contém uma prova de conceito (POC) que ilustra a integração do Apache Kafka com a linguagem de programação Go. A seguir, são apresentados os componentes e o que esperar ao utilizar este repositório.

## Consumidor (Consumer)

O consumidor (`./cmd/consumer/main.go`) é responsável por ler mensagens de um tópico Kafka específico. Ele está configurado para consumir mensagens do tópico "teste". Ao executar o consumidor, espere ver as mensagens recebidas do Kafka sendo exibidas no console.

## Produtor (Producer)

O produtor (`./cmd/producer/main.go`) é responsável por enviar mensagens para um tópico Kafka específico. Ao executar o produtor, ele enviará uma mensagem de teste para o tópico "teste" no Kafka. O relatório de entrega indicará se a mensagem foi enviada com sucesso ou se ocorreu algum erro.

## Dockerfile

O Dockerfile (`./Dockerfile`) é usado para construir a imagem do ambiente Go. Ele inclui as dependências necessárias para o Kafka (`librdkafka-dev`), assim devera ser usado tanto pro consumer quanto pro producer.

## Docker Compose

O arquivo docker-compose (`docker-compose.yml`) define os serviços necessários para criar um ambiente Kafka local. Ele inclui serviços para o aplicativo Go, o [Zookeeper](https://zookeeper.apache.org/), o [Kafka](https://kafka.apache.org/documentation/) e o [Control Center](https://docs.confluent.io/platform/current/control-center/index.html). Ao usar o Docker Compose, espere que todos os serviços do Kafka estejam em execução e prontos para interação.

## Links úteis :globe_with_meridians:
- ### [Github librdkafka](https://github.com/confluentinc/librdkafka)
	
## 
<p align="center">
  <img src="https://miro.medium.com/v2/resize:fit:660/1*6srDxEshHgrk58AJQIAYiQ.png" height="300" alt="Imagem 2">
</p>
