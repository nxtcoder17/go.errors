# go.errors

A lightweight Go error handling library that enhances standard errors with contextual key-value pairs and visual error chain tracking.

## Features

- **Contextual Errors**: Attach key-value pairs to errors for better debugging
- **Error Wrapping**: Build error chains with clear visualization using `|>` separator
- **Clean API**: Fluent interface for easy error construction
- **Standard Compatible**: Implements the standard `error` interface

## Installation

```bash
go get github.com/nxtcoder17/go.errors
```

## Usage

### Basic Error Creation

```go
import "github.com/nxtcoder17/go.errors"

func processUser() error {
    return errors.New("failed to process user").KV("userID", 123)
}
// Output: failed to process user <userID=123>
```

### Error Wrapping

Build error chains to track the flow through your application:

```go
func A() error {
    return errors.New("failed to process A").KV("func", "A")
}

func B() error {
    if err := A(); err != nil {
        return errors.New("failed to process B").Wrap(err).KV("func", "B")
    }
    return nil
}

// Output: failed to process B |> failed to process A <func=B, func=A>
```

The `|>` separator visually shows the error propagation path, making it easy to trace errors through your call stack.

### Multiple Key-Value Pairs

Add multiple contextual fields to your errors:

```go
err := errors.New("database query failed").
    KV("table", "users", "operation", "SELECT", "timeout", "5s")
// Output: database query failed <table=users, operation=SELECT, timeout=5s>
```

## API Reference

### `New(msg string) *Error`
Creates a new error with the given message.

### `KV(kvPairs ...any) *Error`
Adds key-value pairs to the error for context. Keys should be strings, followed by their values.

### `Wrap(err error) *Error`
Wraps another error, creating an error chain. The wrapped error is appended with the `|>` separator.

### `Error() string`
Returns the formatted error message with all context and wrapped errors.

## License

See LICENSE file for details.

## Author

[nxtcoder17](https://github.com/nxtcoder17)
