# Jagger #
Jagger is a simple Go logging library.

## Usage ##

Use default log:

```go
import (
  logger "github.com/JamesYYang/jagger"
)

func main() {
  logger.Info("I'm about to do something!")
	if err := doSomething(); err != nil {
		logger.Errorf("Error running doSomething: %v", err)
	}
}
```

Create a log instance

```go
func main() {
	l := logger.New("NewSystem")
	l.SetFlags(log.Ldate)
	l.SetLevel(logger.ErrorLevel)
	l.SetOutput(os.Stdout)
	if err := doSomething(); err != nil {
		l.Errorf("Error running doSomething: %v", err)
	}
}

```
