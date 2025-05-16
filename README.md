# translate

Golang translation library that supports several free providers.

Current supported provider list:

- [Google translate free API](./providers/google/README.md)
- [MyMemory](./providers/mymemory/README.md)
- [Linguee](./providers/linguee/README.md) (old and unaccurate)

## Installation

```sh
go get github.com/maxwolf8852/translate
```

## Usage

You can use this library with one or several providers (e.g. if the first one fails, the library will request translation from from the next one, and so on).

### One-provider

```go
client, err := translate.New(translate.WithProvider(google.New()))
if err != nil {
    return err
}
out, err := client.Translate(t.Context(), translate.EN, translate.FR, "Hello world!")
if err != nil {
    return err
}
fmt.Println(out)
```

### Multiple-providers (with error skip)

This example skips all errors until the last one and prints the last (successful or not) output.

```go
client, err := translate.New(translate.WithSkipErrors(),
			translate.WithProvider(google.New()),
			translate.WithProvider(mymemory.New()), ...)
if err != nil {
    return err
}
out, err := client.Translate(context.TODO(), translate.EN, translate.FR, "Hello world!")
if err != nil {
    return err
}
fmt.Println(out)
```
