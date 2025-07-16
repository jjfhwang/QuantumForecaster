// cmd/quantumforecaster/main.go
package main

import (
"flag"
"log"
"os"

"quantumforecaster/internal/quantumforecaster"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := quantumforecaster.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
