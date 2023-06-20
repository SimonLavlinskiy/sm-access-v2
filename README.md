
## Git Hooks (НЕОБХОДИМО)

```shell
$ git config core.hooksPath .githooks
```

## Swagger

Для генерации документации использовать команду

* Выполнить команду для linux
    ```shell
    $ ./bin/swag init
    ```
*  Выполнить команду для macOS
    ```shell
    $ ./bin/swag-mac init
    ```
* [Документация по Swag v1.7.8](https://github.com/swaggo/swag/tree/v1.7.8)


## Тестирование 

### Запуск тестов

```shell
$ go test --cover ./tests/...
```

    
