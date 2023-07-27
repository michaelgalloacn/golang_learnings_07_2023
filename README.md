# Running Commands in Golang

Golang files that are meant to be run directly must always start with
```go
package main
```

To run a go file directly from the terminal run
```bash
go run [FILENAME].go
```


To compile a go program into a binary run
```go build [FILENAME].go```

Which should then give you a binary you can run directy from your terminal, IE: `./[FILENAME]` in Linux/OSX
