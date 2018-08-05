# toml2env

```
$ cat example.toml | go run cmd/toml2env/main.go servers.beta
export ip=10.0.0.2
export dc=eqdc10
```

## Usage

### Direct

```bash
eval "$(cat example.toml | go run cmd/toml2env/main.go servers.beta)"
```

### Wrapped

```bash
# wrapped.sh
eval "$(cat example.toml | go run cmd/toml2env/main.go servers.beta)"
```

```bash
source wrapped.sh
```

## See also

- [environment variables - docker-compose .env vs direnv .envrc - Stack Overflow](https://stackoverflow.com/questions/47879565/docker-compose-env-vs-direnv-envrc)
- [bcoe/toml-to-env](https://github.com/bcoe/toml-to-env)

