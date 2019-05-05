# Holzhacker

Holzhacker is a logger for go. It's almost just a configuration for
[lumberjack](https://github.com/natefinch/lumberjack).

## Example

This example creates a log file at `~/logs/PROGRAM_NAME/filename.log`.

```go
package main

import 	"github.com/cody/holzhacker"

func main() {
	myLog := holzhacker.Create("filename")

	myLog.Print("Hello")
	myLog.Printf("1 + 2 = %d", 1+2)
	myLog.Fatal("This is the end and we logged it.")
}

```

## Command line options for developing

With the command line argument `-log` all log messages are duplicated to stdout.

## Licence

MIT.
