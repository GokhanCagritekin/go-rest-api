# In memory key-value store REST-API.

Simple in memory key-value store.

## Usage

Clone the repo

```bash
git clone https://github.com/GokhanCagritekin/go-rest-api.git
```

```bash
go run .\cmd\server\main.go 
```

## Endpoints

- localhost:8080/set?key=key&value=value
- localhost:8080/get?key=key
- localhost:8080/flush