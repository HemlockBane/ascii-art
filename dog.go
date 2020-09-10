package main

import f "fmt"
import "time"

func main(){
  f.Println("  __      _")
  f.Println("o'')}____//")
  f.Println(" `_/      )")
  f.Println(" (_(_/-(_/")


  f.Printf("Printed at %s", time.Now().String());

  // Interesting fact: PrintLn doesn't support format specifiers
  // Here's a link I found: https://stackoverflow.com/questions/45645900/fmt-println-prints-out-format-verbs-like-s

}