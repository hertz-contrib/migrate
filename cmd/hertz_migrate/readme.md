## hertz_migrate

### Introduction

Migrate go web code by analyzing ast, currently `hertz_migrate` supports migrating the following frameworks chi (only `func (http.ResponseWriter, *http.Request)` net/http migrations are supported for now)

### How to install

```bash
go install github.com/hertz-contrib/migrate/cmd/hertz_migrate@latest
```

### How to use

#### Args

```bash
NAME:
   hertz_migrate - A tool for migrating to hertz from other go web frameworks

USAGE:
   hertz_migrate [global options] command [command options] 

VERSION:
   v0.0.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --hz-repo value, -r value                                        Specify the url of the hertz repository you want to bring in. (default: github.com/cloudwego/hertz)
   --target-dir value, -d value                                     project directory you wants to migrate
   --ignore-dirs value, -D value [ --ignore-dirs value, -D value ]  Fill in the folders to be ignored, separating the folders with ",".
      Example:
          hertz_migrate -target-dir ./project -ignore-dirs=kitex_gen,hz_gen
   --help, -h     show help
   --version, -v  print the version

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
