# Go Generate

## 用法

``` sh
# -run if non-empty, specifies a regular expression to select
#        directives whose full original source text (excluding
#        any trailing spaces and final newline) matches the
#        expression.
# -n   The -n flag prints commands that would be executed.
# -x   The -x flag prints commands as they are executed.
#
$ go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]

# run ls
$ go generate -run "^.*ls.*$"
ls.go   pwd.go

# run pwd
$ go generate -run "^.*pwd.*$"
/Users/fangsihao/Documents/code/go/src/examples/stdlib/generate/code
```

``` go
//go:generate ls
package code
```

``` go
//go:generate pwd
package code
```

## 内置环境变量

``` text
$GOARCH
    The execution architecture (arm, amd64, etc.)

$GOOS
    The execution operating system (linux, windows, etc.)

$GOFILE
    The base name of the file.

$GOLINE
    The line number of the directive in the source file.

$GOPACKAGE
    The name of the package of the file containing the directive.

$DOLLAR
    A dollar sign.
```

``` sh
$ go generate -run "^.*echo.*$"
GOARCH=amd64
GOOS=darwin
GOFILE=echo.go
GOLINE=10
GOPACKAGE=code
DOLLAR=$
```

``` go
package code

//go:generate echo GOARCH=$GOARCH
//go:generate echo GOOS=$GOOS
//go:generate echo GOFILE=$GOFILE
//go:generate echo GOLINE=$GOLINE
//go:generate echo GOPACKAGE=$GOPACKAGE
//go:generate echo DOLLAR=$DOLLAR
```

## 编码

符合正则表达式 ^//go:generate .*$, '.*' 表示要执行的命令 