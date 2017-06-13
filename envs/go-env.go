package environment

import (
  "os"
  "fmt"
)

func Environment() {
  for _, env := range os.Environ() {
  fmt.Println(env)
}

}
