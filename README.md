# Gopher

Gopher is a command-line tool written in Go for managing Go projects. It provides a set of features to 
help you build, test, and run your Go applications efficiently.

## Inspirations
1. The inspiration for this project is the `cargo` tool owned by Rust. As a Rust developer, I have come to love cargo, and this project is born out of that mutual love for Rust and Go.
2. Another reason I am building this is to master the use of Vim instead of smart IDEs. This idea was enforced on me by [Timothy Brymes](https://github.com/Brymes) after his first attempt at the Certified Kubernetes Administrator Exam. Hence, this project was built from the terminal alone, with Vim. 
3. The third inspiration for this project is the need for Golang to do something new. The most novel additions to the Golang toolchain is the support for Generics and the Go workspace idea, both of which are not new anymore. 
4. The final inspiration for this project is to try out the Microkernel architecture rumored to be suitable for systems programming. Normally, I would have written this tool in Rust,out of convenience and suitability of tools, but my conscience would not allow me.

# Shout out
Shout out to the `os.WriteFile()` ([here](https://cs.opensource.google/go/go/+/go1.20.5:src/os/file.go;l=720)) function for the extra abilities. Well engineered.

## Installation

To install Gopher, you need to have Go installed on your machine. 
You can download the binary for your operating system from the [Releases](https://github.com/your-username/gopher/releases) section.

### Adding Gopher to the PATH

To use `gopher` as a command from anywhere in your terminal, you need to add the directory containing the `gopher` binary to your 
system's `PATH` environment variable. 
Follow these steps based on your shell:

#### Bash (`.bashrc` or `.bash_profile`)

1. Open the `.bashrc` or `.bash_profile` file in your preferred text editor:

```bash
nano ~/.bashrc
```

or

```bash
nano ~/.bash_profile
```

2. Add the following line to the file, replacing `/path/to/gopher` with the actual path to the `gopher` binary:

```bash
export PATH="/path/to/gopher:$PATH"
```

3. Save the file and exit the text editor.

4. Close and reopen your terminal, or run the following command to apply the changes:

```bash
source ~/.bashrc
```

or

```bash
source ~/.bash_profile
```

#### Zsh (`.zshrc`)

1. Open the `.zshrc` file in your preferred text editor:

```bash
nano ~/.zshrc
```

2. Add the following line to the file, replacing `/path/to/gopher` with the actual path to the `gopher` binary:

```bash
export PATH="/path/to/gopher:$PATH"
```

3. Save the file and exit the text editor.

4. Close and reopen your terminal, or run the following command to apply the changes:

```bash
source ~/.zshrc
```

### Usage

Gopher provides various commands to assist you in your Go project development. Here is an example of the commonly used command:

- `gopher new <project_name>`: Creates a new Go project with the specified name.

  Example: `gopher new myproject`

Please note that the other features are still under development and will be available soon.

## Contributing

Contributions to Gopher are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the GitHub repository.

## License

This project is licensed under the [MIT License](LICENSE).
```
