# Round tripper

## [<<< ---](../index.md)

Паттерн **Round tripper** описывает полный цикл запроса–ответа, часто поверх сети или другого медленного канала.

В Go он формализован интерфейсом `http.RoundTripper`, но общая идея применима и к пользовательским протоколам:

- одна горутина формирует запрос и отправляет его;
- другая горутина получает ответ и связывает его с исходным запросом.

```go
type Request struct {
	ID      int
	Payload string
	RespCh  chan Response
}

type Response struct {
	ID      int
	Payload string
	Err     error
}

// client отправляет запрос и ждёт ответа по отдельному каналу.
func client(reqs chan<- Request, payload string) (Response, error) {
	respCh := make(chan Response, 1)
	req := Request{
		ID:      rand.Int(),
		Payload: payload,
		RespCh:  respCh,
	}

	reqs <- req
	resp := <-respCh
	return resp, resp.Err
}

// server обрабатывает запросы и отправляет ответы обратно в указанный канал.
func server(reqs <-chan Request) {
	for req := range reqs {
		// ... обработка ...
		req.RespCh <- Response{ID: req.ID, Payload: "OK: " + req.Payload}
	}
}
```

Такой паттерн удобен для построения простых RPC‑механизмов поверх каналов.

## Типовые ошибки / антипаттерны
- Не закрывать/не освобождать каналы `RespCh`: ожидание ответа может зависнуть.
- Ошибаться с соответствием `ID`/response: клиент получает не тот ответ.
- Отсутствие таймаута: при недоставке/краше server клиент ждёт бесконечно.

## Практический чеклист
- У каждого запроса уникальный `ID` и ответ маппится строго по этому `ID`.
- `RespCh` создаётся на клиенте и имеет понятный life-cycle.
- Есть таймаут/отмена (через `context`) поверх ожидания ответа.
- Сервер корректно завершает обработку при shutdown.

## Как адаптировать под кейс
- Для реальной сети используйте стандартный `http.RoundTripper`, а для внутренних процессов — этот паттерн поверх каналов.
- Если запросов много, добавьте ограничение concurrency (semaphore/worker pool).

## Связанные паттерны
- `./wait-for-result.md`
- `./context.md`
- `./graceful-shutdown.md`

