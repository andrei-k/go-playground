# Eliza

This code is a port of the original  "Eliza" program to Go. The program appears to have a conversation with the user, using string replacement and basic parsing, to give the appearance of understanding what the user types.  

The responses are based on a Python program by Joe Strout, Jeff Epler and Jez Higgins. The Python code is available at https://github.com/jezhiggins/eliza.py. This code is licensed under the terms of the MIT License.

## Usage

Run in Terminal

```
go run .     
```

Generate an executable binary on Mac OS

```
go build -o eliza main.go
```

Generate an executable binary on Windows

```
go build -o eliza.exe main.go
```