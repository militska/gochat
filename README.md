На данный момент работает  (недо)апи метод вставки сообщения в рэдис.
и вытягивание из него.

В планах: привести в норм вид, +добавить  вебсокеты, мб добавить чат 1-1, что бы это стало похоже на чат,
а не на расшаренный todo лист.

Так же неплохо бы мб отдельным проектом написать ~~вебморду~~  интерфейс для работы с апи

Особенности запуска: с портами всё пока криво. Поэтому что бы узнать какой ипшник выдался
смотрим `docker inspect gochat_app_run_3` (имя контейнера) и там ищем `IPAddress`

