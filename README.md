# SWrap

## Slice Wrapper

adding GC firendly, Well optimised utility method to slice of Go.


## Version

ver 0.1 (2014/01/28)


## API list

- New()
- Make()
- Bytes()
- Len()
- Add()
- Merge()
- Delete()
- Compate()
- Push()
- Pop()
- Unshift()
- Shift()
- Replace()


## Code Generator

default swrap is based on this type.

```go
type SWrap []byte
```

you can also generate your own swrap
based on your original type using render.go

if you want type like blow.

```go
package myapp

type MyType []string
```

you can generate like this

```go
$ go run main/generate.go -p myapp -n MyType -t string -f myapp.go
```

default value for all param makes swap.go

```go
$ go run main/generate.go -p swrap -n SWrap -t byte -f swrap.go
```

## License

The MIT License (MIT)
Copyright (c) 2013 Jxck
