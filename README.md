# Integração do Gin com ngrok


Instalação do ngrok no MacOS:

```sh
brew install ngrok/ngrok/ngrok
```

Configuração do ngrok:

```sh
ngrok config add-authtoken x...y...z
```

Abrindo um tunel ainda sem a integração com Go e Gin:

```sh
hgrok http 3000
```

Fazendo uma integração entre ngrok e gin no Go:

```sh
go get golang.ngrok.com/ngrok
go get github.com/gin-gonic/gin
```

Para execução:

```sh
 NGROK_AUTHTOKEN=x...y...z go run main.go
```
Onde NGROK_AUTHTOKEN é o seu token de autenticação