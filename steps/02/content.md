## Exported names

[Exported names](https://go.dev/tour/basics/3)

> In Go, a name is exported if it begins with a capital letter.

Now let’s take a look on [Go Standard Library](https://pkg.go.dev/std), spesifically math package [here](https://pkg.go.dev/math). We can easily see that math package exports mathematical functions like `Abs`, `Cos`, `Floor` and `Log`. It also exports mathematical constants like `E`, `Pi` and `Phi`.

When you run the example as it is, you’ll see a familiar error message

```text
./prog.go:9:14: cannot refer to unexported name math.pi
```

That’s the exact message we got when we tried to call `math.rand.Intn(10))` function by importing `math` package instead of importing `math/rand` and use `rand.Intn(10)`.

Go compiler would not let you have unused imports in your code. When you try to run the code below in A Tour of Go’s runner,

```go
~{
  title: Example 1
  executable: true
  executeScript: |
    go run ./step/example1.go
  file: example1.go
}~
```

you’ll see this error message.

```text
./prog.go:6:2: imported and not used: "net/http"
```

There are formatting tools for go like [gofmt](https://pkg.go.dev/cmd/gofmt), [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) and so on. These tools usually delete unused imports before the compilation takes place. That’s what happened before when we tried calling `math.rand.Intn(10))` function by importing `math` package in Go Playground before. The formatter deleted the `math` import because it caanot find any usage of any exported identifiers of `math` package.

You can try the code with the unused `net/http` reference on Go Playground now by clicking [here](https://go.dev/play/p/BvbZ84f5N64). When you click Run, you can see that `net/http` line is deleted from imports, and the program compiles and runs successfuly.

Another important thing that needs mentioning here is, “What does unexported mean?” Up to now, we used methods and constants from other packages albeit they are from standard library. What’s our package than? You can clearly see that every .go file starts with a package declaration, and our package is main. You can access unexported identifiers within the same package. In the example below you can see that we define a constant named `pi` and a function named `multiply` and successfully call them even if they are not exported because they belong the `main` package.

```go
~{
  title: Example 2
  executable: true
  executeScript: |
    go run ./step/example2.go
  file: example2.go
}~
```

