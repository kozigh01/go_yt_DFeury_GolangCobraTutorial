# Donald Feury: Golang Cobra Tutorial

From youtube: https://www.youtube.com/watch?v=7U12a-TTtfo
Cobra Project: [github](https://github.com/spf13/cobra) | [generator](https://github.com/spf13/cobra/blob/master/cobra/README.md)

## Examples
* Root: For boolean flags,set default to false, then pass empty flag to set to true

  ```bash
  go run ./cobra-example/main.go
  go run ./cobra-example/main.go --help

  go run ./cobra-example/main.go --version

  go run ./cobra-example/main.go -p -l
  go run ./cobra-example/main.go -pl
  go run ./cobra-example/main.go --localFlag --persistFlag

  ```

* Echo

  ```bash
  go run ./cobra-example/  echo -h
  go run ./cobra-example/  echo times -h

  go run ./cobra-example/ echo abcd efg
  go run ./cobra-example/ echo abcd efg -p

  go run ./cobra-example/ echo times abcd efg -t 3
  go run ./cobra-example/ echo times abcd efg -p -t 3

  go run ./cobra-example/ -p -s "a string" echo something
  go run ./cobra-example/ echo something -p -s "a string"
  go run ./cobra-example/ echo something -ps "a string"

  # required flag
  go run ./cobra-example/  echo times abcd efg

  # to see returned custom error
  go run ./cobra-example/  echo times abcd efg --persistFlag=true -t 4

  # custom suggestion
  go run ./cobra-example/main.go repeat
  ```

* Hello

  ```bash
  go run ./cobra-example/ hello -h
  ```

* Validargs 

  ```bash
  # valid call
  go run ./cobra-example/ validargs three two

  # throw error for invalid argument
  go run ./cobra-example/ validargs five
  ```