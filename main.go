package main

import (
//        "context"
        "fmt"
        "bufio"
        "os"
        "strings"
        "math/rand"
        "time"
        "strconv"
)

func gameRules() {
  fmt.Println("Guess the correct Number Game")
  fmt.Println("---------------------")
  fmt.Println("Game Rules: ")
  fmt.Println("  Guess the correct number by input a number.")
  fmt.Println("  The number must between 0 and 100.")
  fmt.Println("  The machine answer only with two indicators.")
  fmt.Println("  Indicators ::== [ above | below ] ")
  fmt.Println("  If the number is lower than your current guess,")
  fmt.Println("  than the 'below' indicator is shown.")
  fmt.Println("  Is the number above your current guess,")
  fmt.Println("  than the 'above' indicator is shown.")
  fmt.Println("")
  fmt.Println("Control Commands: ")
  fmt.Println("  exit")
  fmt.Println("  help")

}

func isInt(s string) bool {
    _, err := strconv.ParseInt(s, 10, 64)
    return err == nil // err == nil means there is no error, so True means is valid Integer
}

func main() {
  var limit int = 100

  // use always a new Seed
  rand.Seed( time.Now().UnixNano() )
  var my_secret_number int = rand.Intn( limit )

  reader := bufio.NewReader(os.Stdin)

  gameRules()

  // endless loop
  for {
    fmt.Println("guess a number between 0 and ", limit)
    fmt.Print("-> ")
    text, _ := reader.ReadString('\n')
    // convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)

    foo := isInt(text)
    if foo == true {
      guess, _ := strconv.Atoi(text)

      if guess < my_secret_number {
          fmt.Println("above!")
      } else if guess > my_secret_number {
          fmt.Println("below!")
      } else if guess == my_secret_number {
          fmt.Println("================")
          fmt.Println("=== success! ===")
          fmt.Println("================")
          os.Exit(0)
      }
    } else {
      // This is a text/string
      switch text  {
        case "foobar":
          fmt.Println("foobar yes")
        case "barfoo":
          fmt.Println("yes barfoo!")
        //case "debug":
        //  fmt.Println("secret", my_secret_number)
        case "exit":
          os.Exit(0)
        case "help":
          gameRules()
        default:
         fmt.Println("something completly different...")
      }
    }

  } // for

}
