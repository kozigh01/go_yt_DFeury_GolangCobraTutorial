# Donald Feury: Golang Cobra Tutorial
From youtube: https://www.youtube.com/watch?v=7U12a-TTtfo
Cobra Project: [github](https://github.com/spf13/cobra)

## Examples
```
$ go run ./cobra-example/  echo times -h
$ go run ./cobra-example/ echo times abcd efg --persistFlag=false -t 3

# required flag
go run ./cobra-example/  echo times abcd efg

# to see returned custom error
$ go run ./cobra-example/  echo times abcd efg --persistFlag=true -t 4
```