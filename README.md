# Gopher

Gopher is a command-line tool written in Go for managing Go projects. It provides a set of features to 
help you build, test, and run your Go applications efficiently. 

Read about the project [here](https://theghostmac.github.io/posts/gopher).

![Demo of gopher](gopher-demo.png)

## Inspirations
1. The inspiration for this project is the `cargo` tool owned by Rust. As a Rust developer, I have come to love cargo, and this project is born out of that mutual love for Rust and Go.
2. Another reason I am building this is to master the use of Vim instead of smart IDEs. This idea was enforced on me by [Timothy Brymes](https://github.com/Brymes) after his first attempt at the Certified Kubernetes Administrator Exam. Hence, this project was built from the terminal alone, with Vim. 
3. The third inspiration for this project is the need for Golang to do something new. The most novel additions to the Golang toolchain is the support for Generics and the Go workspace idea, both of which are not new anymore. 
4. The final inspiration for this project is to try out the Microkernel architecture rumored to be suitable for systems programming. Normally, I would have written this tool in Rust,out of convenience and suitability of tools, but my conscience would not allow me.

## Installation

To install Gopher, you need to have Go installed on your machine. Next, run the installation:
```shell
go install github.com/theghostmac/gopher@latest
```

## Usage

Gopher provides a set of commands to manage your Go projects efficiently. Here's how you can use each command:

### `new`

Create a new Go project with the specified name and project type.

```bash
gopher new my_project --web
```

Creates a new Go web project named `my_project` with a basic web server boilerplate.

```bash
gopher new cli_tool --app
```

Creates a new command-line tool project named `cli_tool` with a basic application boilerplate.

### `add`

Add external Go dependencies to your project using `go get`.

```bash
gopher add github.com/gin-gonic/gin
```

Adds the `gin` package from GitHub to your project.

### `build`

Build your Go project.

```bash
gopher build
```

Compiles and builds your Go project, creating an executable binary.

### `init`

Initialize Go modules for your project.

```bash
gopher init
```

Initializes Go modules for your project, allowing you to manage external dependencies easily.

### `test`

Run tests for your Go project.

```bash
gopher test
```

Compiles and runs tests in your project, showing the test results.

### `run`

Run your Go project.

```bash
gopher run
```

Builds and executes your Go project.

### `check`

Check the syntax of your Go codebase or individual files for any issues.

```bash
gopher check
```

Checks the entire codebase for any "TODO" comments and reports them as issues.

```bash
gopher check ./cmd
```

Checks the `cmd` directory for any "TODO" comments and reports them as issues.

```bash
gopher check main.go
```

Checks the `main.go` file for any "TODO" comments and reports them as issues.

---

In this updated `#Usage` section, we've provided a detailed explanation of each command and included example usage for better clarity. Users can now follow the examples to understand how to use the `gopher` app effectively for their Go projects.

# Shout out
Shout out to the `os.WriteFile()` ([here](https://cs.opensource.google/go/go/+/go1.20.5:src/os/file.go;l=720)) function for the extra abilities. Well engineered.

## Contributing

Contributions to Gopher are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the GitHub repository.

## License

This project is licensed under the [MIT License](LICENSE).