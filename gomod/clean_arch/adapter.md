# Адаптеры / Инфраструктура

## [<<< ---](../clean_arch.md)

**Контроллеры, сетевые шлюзы, презентеры (Controllers, Gateways(репозитории), Presenters)** — набор адаптеров, которые наиболее удобным способом преобразуют данные из сценариев использования и формата объектов для передачи в верхний слой (обычно пользовательский интерфейс).

Адаптеры могут импортировать внутренние слои. Обычно они будут работать с типами, найденными в приложении и домене, например, извлекая их из базы данных. 

Порты могут импортировать внутренние слои. Порты — это точки входа в приложение, поэтому они часто исполняют службы или команды приложения. Однако они не могут напрямую обращаться к адаптерам.

На оригинальной схеме есть слово Controllers. Оно появилось на схеме из-за frontend’a, в частности из Ruby On Rails. Там зачастую разделяют Controller, который обрабатывает запрос и отдает результат, и Presenter, который выводит этот результат на View.