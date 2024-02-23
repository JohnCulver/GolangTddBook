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
Provides IDE features

## errcheck
```sh
go install github.com/kisielk/errcheck@latest
# cd into a package
~/go/bin/errcheck .
```
Checks for unchecked errors.


## go vet
```sh
go vet
```
Detects suspscious code and sublte errors

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

## VS Code Exstentions

### Go test explorer
    Great for visually running tests
    Also allows for running individual benchmarks

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

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/select

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/sync

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/context

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/roman-numerals

# Next Step
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reading-files

# TODO later
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reflection

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/math
