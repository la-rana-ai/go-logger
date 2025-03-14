# Логгер

Пакет `Logger` реализует чуть более сложный вариант журналирования, чем пакет `log`, по сути являясь его надстройкой.
Помимо стандартного вывода сообщения в поток, данный пакет позволяет выводить данные в json, а так же связывается с
файлами вывода по их имени, а не по открытому соединению.
Помимо этого работа данного пакета ограничивается журналированием сообщения, без дополнительных действий вроде
`os.Exit(1)`, а так же данный логгер имеет больше видов сообщений, нежели пакет `log`

### 1. Спецификация

#### 1.1 Основы

Предоставляет в Interface восемь методов записи журналов на восьми
уровнях [RFC 5424](https://datatracker.ietf.org/doc/html/rfc5424#section-6.2.1)

* Emergency (0) - система не может использоваться

```Go
log, _ := logger.New("channelName", nil)
ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
log.Emergency(ctx, "Message about error")
```

* Alert (1) - действие должно быть предпринято немедленно

```Go
log, _ := logger.New("channelName", nil)
ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
log.Alert(ctx, "Message about error")
```

* Critical (2) - критические условия работы программы

```Go
log, _ := logger.New("channelName", nil)
ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
log.Critical(ctx, "Message about error")
```

* Error (3) - программа работает с ошибками

```Go
log, _ := logger.New("channelName", nil)
ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
log.Error(ctx, "Message about error")
```

* Warning (4) - сообщение требует внимание, но работа в рамках приемлемого

```Go
log, _ := logger.New("channelName", nil)
ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
log.Warning(ctx, "Message")
```

* Notice (5) - нормальное, но важное состояние

```Go
log, _ := logger.New("channelName", nil)
ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
log.Notice(ctx, "Message")
```

* Info (6) - информационные сообщения

```Go
log, _ := logger.New("channelName", nil)
ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
log.Warning(ctx, "Message")
```

* Debug (7) - сообщения уровня отладки

```Go
log, _ := logger.New("channelName", nil)
ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
log.Warning(ctx, "Message")
```

---

#### 1.2 Сообщение

Сообщение состоит из 3 параметров

* ctx - контекст запроса или же работающего процесса. В контексте передаётся переменная x-request-id (идентификатор
  лога) и user-agent (имя того, кто породил выполнение данного процесса).
* message - непосредственно сообщение в данном логе
* []context - элементы для описания того контекста непосредственного события, которые способствовали созданию лог
  сообщения

#### 1.2 Настройка

Подразумевается, что для каждого модуля программы будет использоваться свой канал записи логов. Так же каждый канал
настраивается отдельно, и может иметь свои параметры.
При создании канала мы должны указать

* channel name - имя канала логирования
* структуру интерфейса InterfaceOption (Для простоты предлагается стандартная структура Option). Данный интерфейс
  содержит:
*
    * Output (string) - имя стандартного канала вывода, или имя файла для записи сообщений данного канала
*
    * Format (InterfaceFormat) - формат вывода. Есть 2 варианта, FormatPlain для вывода данных через форматированную
      строку, и FormatJSON ждя вывода данных через формат JSON
*
    * MinimalLevel (InterfaceLevel) - минимальный уровень логирования. В канале идёт запись всех сообщений этого уровня
      логирования, а так же сообщения с меньшим кодом.
*
    * Flags (InterfaceFlags) - структура набора флагов. Для удобства существует структура OptionFlags.
*
    *
        * Date - выводить ли дату в логе
*
    *
        * Time - выводить ли в логе время
*
    *
        * Microseconds - выводить ли в логе наносекунды. Работает, если стоит флаг Time
*
    *
        * LongFile - выводить ли адрес файла и строку, где произошел вызов лога
*
    *
        * ShortFile - выводить ли имя файла и строку, где произошел вызов лога. Перекрывает параметр LongFile
*
    *
        * Utc - Выводить ли время, приводя его к часовому поясу utc
*
    *
        * MsgPrefix - Выводить ли префикс/имя канала
*
    *
        * StdFlags - выставляет флаги Date и Time

### 2 Пример

```Go
package main

import (
    "context"
    "github.com/google/uuid"
    "la-rana-ai/go-logger"
)

func main() {
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log, _ := logger.New("testChannel", &logger.Option{
        Format:       logger.FormatPlain,
        MinimalLevel: logger.Notice,
        Output:       logger.STDOUT,
        Flags: &logger.OptionFlags{
            Date:         true,
            Time:         true,
            Microseconds: true,
            LongFile:     true,
            ShortFile:    false,
            Utc:          true,
            MsgPrefix:    true,
            StdFlags:     true,
        },
    })
    log.Debug(ctx, "этот сообщение не будет записано")
    log.Notice(ctx, "а это сообщение будет записано")
    log.Notice(ctx, "сообщение с контекстом", struct {
        Name string
    }{
        Name: "John Doe",
    })
}

```



