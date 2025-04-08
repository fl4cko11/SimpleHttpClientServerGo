![Task](https://github.com/fl4cko11/SimpleHttpClientServerGo/blob/main/task.png)

[Конфигурация Клиента и Сервера](https://github.com/fl4cko11/SimpleHttpClientServerGo/tree/main/config)

[Клиент](https://github.com/fl4cko11/SimpleHttpClientServerGo/blob/main/internal/client/activation.go) использует [генерации JSON](https://github.com/fl4cko11/SimpleHttpClientServerGo/blob/main/internal/jsonGenerator/funcs.go) событий, причём на каждое 100е отправляется дубликат, число событий генрерируется каждый период случайно, период отправки генрируется 1 раз

[Сервер](https://github.com/fl4cko11/SimpleHttpClientServerGo/blob/main/internal/server/api.go) настроен на приянтие только POST запроса по пути /endpoint, хэширует данные запроса и записывает в очередь запросов, дубликаты ищет с помощью хэш-таблицы и собирает указанные метрики.
