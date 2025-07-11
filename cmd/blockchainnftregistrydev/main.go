// cmd/blockchainnftregistrydev/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftregistrydev/internal/blockchainnftregistrydev"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftregistrydev.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
