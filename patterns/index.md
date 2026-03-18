

# [<<< ---](../README.md)

## Patterns of Enterprise Application Architecture

###### Основные
- [Шлюз (Gateway)](./PoEAA/basic/gateway.md)
- [Mapper](./PoEAA/basic/mapper.md)
- [Layer Supertype](./PoEAA/basic/layer-supertype.md)
- [Отделенный интерфейс (Separated Interface)](./PoEAA/basic/separated-interface.md)
- [Registry (Реестр)](./PoEAA/basic/registry.md)
- [Объект-значение (Value Object)](./PoEAA/basic/value-object.md)
- [Частный случай (Special Case)](./PoEAA/basic/special-case.md)
- [Дополнительный модуль (Plugin)](./PoEAA/basic/plugin.md)
- [Фиктивная служба (Service Stub)](./PoEAA/basic/service-stub.md)
- [Record Set](./PoEAA/basic/record-set.md)
- [Заместитель (Proxy)](./PoEAA/basic/proxy.md)
- [Functional Options](./PoEAA/basic/functional-options.md)

###### Структурные
- [Приспособленец (Flyweight)](./PoEAA/struct/flyweight.md)
- [Composite](./PoEAA/struct/composite.md)

## Golang

- [Bounded parallelism](./golang/bounded-parallelism.md)
- [For-select](./golang/for-select.md)
- [For-select-done](./golang/for-select-done.md)
- [Or-channel](./golang/or-channel.md)
- [Tee channel](./golang/tee-channel.md)
- [Bridge channel](./golang/bridge-channel.md)
- [Ring buffer](./golang/ring-buffer.md)
- [Fan-In](./golang/fan-in.md)
- [Fan-Out](./golang/fan-out.md)
- [Publish & Subscribe](./golang/pub-sub.md)
- [Generator](./golang/generator.md)
- [Select Statement with Timeout](./golang/select-statement-with-timeout.md)
- [Wait For Result](./golang/wait-for-result.md)
- [Wait For Task](./golang/wait-for-task.md)
- [Pooling](./golang/pooling.md)
- [Drop](./golang/drop.md)
- [Cancellation](./golang/cancellation.md)
- [Take First N](./golang/take-first-n.md)
- [Map & Filter](./golang/map-filter.md)
- [Filter](./golang/filter.md)
- [Pipeline](./golang/pipeline.md)
- [Worker Pool](./golang/worker-pool.md)
- [Queuing](./golang/queuing.md)
- [Context](./golang/context.md)
- [Exponential backoff](./golang/exponential-backoff.md)
- [Fault-tolerance](./golang/fault-tolerance.md)
- [Deadline](./golang/deadline.md)
- [Fail-Fast](./golang/fail-fast.md)
- [Handshaking](./golang/handshaking.md)
- [Steady-State](./golang/steady-state.md)
- [Stopping short](./golang/stopping-short.md)
- [Explicit cancellation](./golang/explicit-cancellation.md)
- [Digesting a tree](./golang/digesting-a-tree.md)
- [Parallel digestion](./golang/parallel-digestion.md)
- [Conclusion](./golang/conclusion.md)
- [Round tripper](./golang/round-tripper.md)
- [errgroup (параллельные задачи + отмена при первой ошибке)](./golang/errgroup.md)
- [singleflight (coalescing дублирующихся запросов)](./golang/singleflight.md)
- [debounce (схлопывание событий по таймеру)](./golang/debounce.md)
- [throttle / rate limiter (ограничение частоты)](./golang/rate-limiter.md)
- [token bucket / leaky bucket (варианты rate limiting)](./golang/token-bucket-leaky-bucket.md)
- [jitter для retry/backoff (рандомизация задержек, чтобы избегать “thundering herd”)](./golang/jitter.md)
- [sync.Pool / object pooling (переиспользование объектов)](./golang/sync-pool.md)
- [graceful shutdown / drain каналов (упорядоченное завершение конвейеров без утечек)](./golang/graceful-shutdown.md)

## Микросервисы

- [Shared database](./microservice/shareddatabase.md)
- [Database per Microservice](./microservice/databasepermicro.md)
- [Gateway Routing](./microservice/gatewayrouting.md)
- [Gateway Aggregation](./microservice/gatewayaggregation.md)
- [Gateway Offloading паттерн](./microservice/gatewayoffloading.md)
- [Saga и распределенные транзакции](./microservice/saga.md)
- [CQRS](./microservice/cqrs.md)
- [Event Sourcing](./microservice/eventsourcing.md)
- [Aggregator](./microservice/aggregator.md)
- [Chain](./microservice/chain.md)
- [Branch](./microservice/branch.md)
- [Ambassador](./microservice/ambassador.md)
- [Circuit Breaker](./microservice/circuitbreaker.md)
- [Choreography](./microservice/choreography.md)
- [Bulkhead](./microservice/bulkhead.md)
- [Sidecar](./microservice/sidecar.md)

## Источники данных

- [Row Data Gateway (Шлюз к данным записи)](./sourcedata/rowdatagateway.md)
- [Active Record (Активная запись) ](./sourcedata/activerecord.md)
- [Table Data Gateway (Шлюз к данным таблицы)](./sourcedata/tabledatagateway.md)
- [Data Mapper](./sourcedata/datamapper.md)

###### Источники

- [https://ornlu-is.github.io/categories/golang/](https://ornlu-is.github.io/categories/golang/)
- [https://habr.com/ru/companies/otus/articles/781964/](https://habr.com/ru/companies/otus/articles/781964/)
- [https://books.studygolang.com/go-patterns/](https://books.studygolang.com/go-patterns/)

