## Garbage

### How to install

```bash
go install github.com/hertz-contrib/migrate/cmd/garbage@latest
```

### How to use

```bash
go run main.go -targer-dir ./haha
```

### Warn

1. You must make sure that each subproject in the `target-dir` has a `go.mod` file, but you can continue to use it even if it doesn't
