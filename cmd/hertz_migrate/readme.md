## hertz_migrate

### How to install

```bash
go install github.com/hertz-contrib/migrate/cmd/hertz_migrate@latest
```

### How to use

#### Args

```bash
Usage of hertz_migrate:
  -hz-repo string
         (default "github.com/cloudwego/hertz")
  -hz-version string
        add hertz version when tool exec go get ...
  -ignore-dirs value
        Fill in the folders to be ignored, separating the folders with ",".
        Example:
            hertz_migrate -target-dir ./project -ignore-dirs=kitex_gen,hz_gen
                
  -target-dir string
        target project directory
  -v    v0.0.1
```

#### Usage
```bash
# example hertz_migration -target-dir ./haha
hertz_migrate -target-dir ${dir-name}
```

### Adapt progress
[readme](./adapt.md)

### Warn

1. You must make sure that each subproject in the `target-dir` has a `go.mod` file, but you can continue to use it even if it doesn't
