# Логгер

Пакет `Logger` реализует более сложный вариант журналирования, чем пакет `log`, фактически являясь его надстройкой.

Помимо стандартного вывода сообщений, данный пакет позволяет:
- Выводить данные в формате JSON.
- Связываться с файлами вывода по их имени (а не по открытому соединению).
- Ограничиваться исключительно журналированием сообщений, исключая действия вроде `os.Exit(1)`.
- Предоставлять больше типов сообщений по сравнению с пакетом `log`.

---

## 1. Спецификация

### 1.1 Основы

Предоставляет в Interface восемь методов для записи журналов в соответствии с уровнями из [RFC 5424](https://datatracker.ietf.org/doc/html/rfc5424#section-6.2.1):

1. **Emergency (0)** — система не может использоваться.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Emergency(ctx, "Message about error")
    ```

2. **Alert (1)** — действие должно быть предпринято немедленно.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Alert(ctx, "Message about error")
    ```

3. **Critical (2)** — критические условия работы программы.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Critical(ctx, "Message about error")
    ```

4. **Error (3)** — программа работает с ошибками.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Error(ctx, "Message about error")
    ```

5. **Warning (4)** — сообщение требует внимания, но работа в рамках приемлемого.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Warning(ctx, "Message")
    ```

6. **Notice (5)** — важное сообщение в рамках нормальной работы.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Notice(ctx, "Message")
    ```

7. **Info (6)** — информационные сообщения.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Info(ctx, "Message")
    ```

8. **Debug (7)** — сообщения уровня отладки.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Debug(ctx, "Message")
    ```

---

### 1.2 Сообщение

Сообщение состоит из трех параметров:
- **ctx** — контекст запроса или работающего процесса (содержит такие переменные, как `x-request-id` — идентификатор лога, и `user-agent` — имя источника выполнения процесса).
- **message** — сообщение в данном логе.
- **[]context** — дополнительные элементы, описывающие события, приведшие к созданию лога.

---

### 1.3 Настройка

Для каждого модуля программы используется свой канал записи логов, который настраивается индивидуально.

При создании канала в функцию `New` передаются параметры:
- **Имя канала**: уникальное название канала.
- **InterfaceOption**: структура, содержащая настройки канала, такие как:
    - **Output**: имя стандартного канала вывода или файла.
    - **Format**: формат вывода (`FormatPlain` для форматированных строк или `FormatJSON` для JSON).
    - **MinimalLevel**: минимальный уровень логирования.
    - **Flags**: параметры форматирования (`OptionFlags`), включая:
        - Date, Time, Microseconds — вывод даты, времени и наносекунд.
        - LongFile, ShortFile — вывод полного пути или только имени файла.
        - Utc — приведение времени к часовому поясу UTC.
        - MsgPrefix — отображение префикса или имени канала.
        - StdFlags — включение параметров Date и Time.

---

## 2. Пример использования

```go
ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
log, _ := logger.New("testChannel", &logger.Option{
        Format: &logger.FormatPlain,
        MinimalLevel: logger.Notice,
        Output: logger.STDOUT,
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
log.Debug(ctx, "это сообщение не будет записано")
log.Notice(ctx, "а это сообщение будет записано")
log.Notice(ctx, "сообщение с контекстом", struct {
    Name string
}{
    Name: "John Doe",
})
```