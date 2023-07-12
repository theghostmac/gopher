# Future Features

## The `new` feature:
- [ ] The `gopher new` keyword should automatically initialize a new `git` repository with the project.
- [ ] The `gopher new` keyword should accept a flag `--vcs none` to exclude initializing a new `git` repository.

## The `build` feature:
- [ ] The build feature should accept a `--release` flag, that builds the executable and pushes it to a `releases/` directory



## Tests support
There will be support for running tests with `gopher`, including the following specific features:
- [ ] The `gopher test` keyword will run tests in the entire code base, for every `*_test.go` files.
- [ ] The `gopher test` keyword will accept a directory and/or filename.
- [ ] Other


## go-check
I will remove the [go-check ](https://github.com/theghostmac/go-check) tool and add a `gopher check` feature that:
- [ ] compiles the codebase and executes it
- [ ] if it executes normally, deletes the executable.


## Versioning info
- [ ] The `gopher --version` command should work by giving the downloaded version of the tool.

## Listing available commands
- [ ] Running `gopher --list` should release a catalog of supported commands and what they do, in a tabular manner.
- [ ] Adding the `gopher --help` command.
- [ ] Making `--help` accept other commands and provide help information on them.


## Explain code
- [ ] Adding a `gopher --explain` code. Cargo's feature just runs rustc --explain under the hood. I could go a little further. Or just ignore this feature.

## Update dependencies
- [ ] Running `gopher update` in a codebase should locate `go.mod` and update dependencies.

