# On Log

We need to log wisely.

## On/Off Log Easily

```go
package main

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", 0)

func main() {
	if logger != nil {
		logger.Println("something")
	}

	logger = nil

	if logger != nil {
		logger.Println("soemthing else")
	}
}
```

Output:

```
something
```

Using custom logger is more easy for hacking.

The code above is pretty simple, but it's very useful in production, from my eyes.

You could enable the logging easily, also disable, according to different
environments. When in developing, you want to log all the informations, and in
tests, you don't want to see them except some statistical results offered by
your testing tools. In production, you need to category your logs into
different files with different time stamp or something. All in all, logging
varies with environments.

## Logging With More Information

Contexts of logging is also important: what time, what file and which line.

```go
package main

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

func main() {
	if logger != nil {
		logger.Printf("something")
	}

	logger = nil

	if logger != nil {
		logger.Printf("soemthing else")
	}
}
```

Output:

```
2014/01/04 17:04:48 log2.go:12: something
```

With time stamp, file names, and line numbers, suddenly the log files just
become more specific and more useful.

## Saving Logs

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func init() {
	var now = fmt.Sprintf("log_%s.log", time.Now().Format("2006_01_02_15_04_05"))
	output, err := os.OpenFile(now, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
		return
	}
	logger = log.New(output, "", log.LstdFlags|log.Lshortfile)
}

func main() {
	if logger != nil {
		logger.Printf("something\n")
	}

	logger = nil

	if logger != nil {
		logger.Printf("should not be logged")
	}
}
```

Output:

```
$ cat log_2014_01_04_19_15_43.log
2014/01/04 19:15:43 log3.go:24: something
```

Now all logs will be saved in a different file with time stamp in its name.

Actually, we could also have different log files for different kind of logs,
separating them could be useful some times, like need to gather some
particularly verbose output for debugging some problems, then you don't have
to pollute the main logs.
