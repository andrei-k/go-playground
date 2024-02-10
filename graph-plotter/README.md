# Graph plotter

## Set up this Go project
Verify that you've installed Go and it's up to date by running this command in the terminal:
```
go version
```

If Go is not installed or isn't up to date, follow the instructions on the [official Go website](https://golang.org/doc/install).

Close the terminal and open a new one to ensure that the changes take effect. However, if `go version` returns an `command not found: go`, you need to add the PATH to the `~/.zshrc` file.

Add these lines to the `~/.zshrc` file:
```
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$GOPATH/bin
```

Close the terminal and open a new one to ensure that the changes take effect.

Navigate to project folder.

Initialize the project by running this command in the terminal.
```
go mod init graph-plotter
```

Add module requirements necessary to build the packages and dependencies.
```
go mod tidy
```

## Usage
Compile and run the package.
```
go run .
```
