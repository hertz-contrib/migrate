## hertz_migrate

### How to install

```bash
go install github.com/hertz-contrib/migrate/cmd/hertz_migratel@latest
```

### How to use

#### Args

```bash
Usage of ./hertz_migrate:
  -hz-repo string
         (default "github.com/cloudwego/hertz")
  -target-dir string
        target project directory
  -v    v0.0.1
```

#### Usage
```bash
# example hertz_migration -target-dir ./haha
hertz_migrate -target-dir ${dir-name}
```

### Warn

1. You must make sure that each subproject in the `target-dir` has a `go.mod` file, but you can continue to use it even if it doesn't
