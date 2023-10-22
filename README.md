# go-i18n-sample

## How to run the API server

```zsh
cd go-i18n-sample
go run .
```

## How to send a request with language specified

### Japanese

```zsh
curl --header "Accept-Language:ja" localhost:8080/greet
```

### English

```zsh
curl --header "Accept-Language:en" localhost:8080/greet
```
