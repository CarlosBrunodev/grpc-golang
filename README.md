![image grpc](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*c9Nw9AKUfuV_f0Wve3zbew.png)

## Introdution gRPC 

### Diference between gRPC and RPC 

- gRPC: é um framework do google de alto desempenho desenvolvido pelo google para atender chamadas RPC (Remote Procedure Call) utiliza a *protobuf* como mecanismo de serializacao de dados que tem papel significativo na reducao de trafego de rede 

- RPC: é a definicao de um protocolo que executa procedimentos em outros computadores atravez de rede.

### Protocol Buffer

Protocol buffer ou protobuf é um protocolo de mensagens agnostico de linguagem criado pela google para serializar dados estruturados.

A transferencia utilizando protobuf chega a ser 6x mais rapida se comparada com json, isso acontece pois as mensagens serializadas com o protobuf sao enviadas em binario deixando o arquivo de transferencia bem menor .


O diagrama a seguir mostra como o protobuf trabalha:

![esquema](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*S0nYiTkdNglJ_zADT4Inww.png)

## Arquitetura gRPC

Como em muitos servicos RPC, o gRPC utiliza a ideia de definir um servidor e um cliente, no qual o servidor especifica os metodos para manipular as chamadas dos clientes. No lado do cliente é definido um stub que contem os mesmo metodos que o servidor.

![Diagram de um servico gRPC](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*FPso9QBErw6SX1ynuo8qHQ.png)

### Principais vantagens 

- Leve e rapida pois utiliza codificacao binaria http/2
- disponivel em varias linguagens 
- multiplataform 
- dar suporte a chamadas streamings do cliente para o servidor 


### Tipos de api com grpc 

    1.Unary
    2.Server streaming
    3.client streaming
    4.Bi direction streaming

[Link de referencia](https://marcoslimadev.medium.com/introdu%C3%A7%C3%A3o-ao-grpc-d0935f5f70cf)