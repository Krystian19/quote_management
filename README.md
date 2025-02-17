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

### DB diagram

(Inventory Item Supply is no longer relevant)

![Alt text](images/db_schema.png?raw=true)

### External BFF Pprof

Snapshot taken when running the `getQuotes` query with all the associated fields included:

![Alt text](images/external_bff_cpu_profile.png?raw=true)
