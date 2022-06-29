package main

import (
	"fmt"

	"github.com/mymmrac/fascii"
)

func main() {
  fmt.Println("fASCII")
  fmt.Println()

  r, err := fascii.NewRendere(fascii.Standard)

  text := r.Render("fASCII")
  fmt.Print(text)

  fmt.Println(err, len(text))
}
