# Logger

The `Logger` package implements a more advanced logging system compared to the `log` package, essentially serving as its extension.

In addition to standard message output, this package allows:
- Outputting data in JSON format.
- Associating with output files by their names rather than open connections.
- Restricting its operation to logging messages only, without additional actions like `os.Exit(1)`.
- Providing more types of messages compared to the `log` package.

---

## 1. Specification

### 1.1 Basics

The package provides an Interface with eight methods for logging messages at eight levels, as defined in [RFC 5424](https://datatracker.ietf.org/doc/html/rfc5424#section-6.2.1):

1. **Emergency (0)** — the system is unusable.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Emergency(ctx, "Message about error")
    ```

2. **Alert (1)** — immediate action is required.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Alert(ctx, "Message about error")
    ```

3. **Critical (2)** — critical conditions in program operation.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Critical(ctx, "Message about error")
    ```

4. **Error (3)** — the program is operating with errors.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Error(ctx, "Message about error")
    ```

5. **Warning (4)** — attention is required, but the operation is within acceptable bounds.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Warning(ctx, "Message")
    ```

6. **Notice (5)** — important messages within normal operation.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Notice(ctx, "Message")
    ```

7. **Info (6)** — informational messages.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Info(ctx, "Message")
    ```

8. **Debug (7)** — debug-level messages.
    ```go
    log, _ := logger.New("channelName", nil)
    ctx := context.WithValue(context.Background(), "x-request-id", uuid.New())
    log.Debug(ctx, "Message")
    ```

---

### 1.2 Message

A log message consists of three parameters:
- **ctx** — the context of a request or running process (contains variables like `x-request-id` — log identifier, and `user-agent` — the name of the source that initiated the process).
- **message** — the actual log message.
- **[]context** — additional elements describing the event context that led to the log creation.

---

### 1.3 Configuration

Each program module uses its own logging channel, which is configured individually.

When creating a channel, the following parameters must be set:
- **Channel Name**: a unique name for the channel.
- **InterfaceOption**: a structure containing channel settings, such as:
    - **Output**: the name of the standard output channel or file.
    - **Format**: output format (`FormatPlain` for formatted strings or `FormatJSON` for JSON).
    - **MinimalLevel**: the minimum logging level.
    - **Flags**: formatting options (`OptionFlags`), including:
        - Date, Time, Microseconds — output of date, time, and nanoseconds.
        - LongFile, ShortFile — output of the full file path or just the file name.
        - Utc — converting the time to UTC time zone.
        - MsgPrefix — displaying the prefix or channel name.
        - StdFlags — enabling Date and Time flags.

---

## 2. Example Usage

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
log.Debug(ctx, "this message will not be logged")
log.Notice(ctx, "this message will be logged")
log.Notice(ctx, "message with context", struct {
    Name string
}{
    Name: "John Doe",
})
```