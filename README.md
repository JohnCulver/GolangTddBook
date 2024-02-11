# GolangTddBook
Working through https://quii.gitbook.io/learn-go-with-tests/

# Helpful tools

## GoDoc
```sh
go get golang.org/x/tools/cmd/godoc
go install golang.org/x/tools/cmd/godoc
~/go/bin/godoc -http=:6060
```

## GoPls
```sh
brew install gopls
```

## errcheck
```sh
go install github.com/kisielk/errcheck@latest
~/go/bin/errcheck .
```

## VScode settings
```json
{
    "go.formatTool": "gofmt",
    "go.useLanguageServer": true,
    "go.lintTool": "golangci-lint",
    "go.lintFlags": [
        "--fast"
    ],
    "[go]": {
        "editor.insertSpaces": true,
        "editor.formatOnSave": true,
        "editor.defaultFormatter": "golang.go",
    }
}
```

# Additonal Commands

## Test Benchmarks
```sh
go test -v ./... -bench=. -run=xxx -benchmem
```


# Completed
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors

# Next Step
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps
