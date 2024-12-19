## Quote Management

[![build](https://github.com/Krystian19/quote_management/actions/workflows/build.yml/badge.svg)](https://github.com/Krystian19/quote_management/actions/workflows/build.yml)

### How to run locally ?

```sh
docker compose --profile app up -d
```

And then run your migrations

```sh
./scripts/migrations.sh
```

GraphQl playground should be available @ http://localhost:4000/

### Project diagram

![Alt text](images/pic.png?raw=true)
