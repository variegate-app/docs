# sync.Pool / object pooling

## [<<< ---](../index.md)

`sync.Pool` — это механизм переиспользования объектов, уменьшающий давление на GC.

Идея object pooling:

- вместо создания нового объекта на каждую операцию;
- мы берем объект из пула, используем, приводим к “чистому” состоянию и возвращаем обратно.

```go
package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func main() {
	buf := bufPool.Get().(*bytes.Buffer)
	// обязательно сбрасываем состояние объекта перед повторным использованием
	buf.Reset()

	buf.WriteString("hello")
	fmt.Println(buf.String())

	// возвращаем объект в пул
	bufPool.Put(buf)
}
```

Важно: `sync.Pool` — не гарантия “строгого” контроля объектов. Планировщик/GC может очищать пул по своему усмотрению, поэтому объект после `Get()` всегда должен быть корректно инициализирован/сброшен.

