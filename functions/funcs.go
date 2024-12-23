package functions

import "fmt"

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}
