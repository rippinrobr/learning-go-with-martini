package main

// loading in the Martini package
import "github.com/codegangsta/martini"

func main() {
  // if you are new to Go the := is a short variable declaration
  m := martini.Classic()

  // the func() call is creating an anonymous function that retuns a stringa
  m.Get("/", func() string {
    return "Where are the attributes?!?!"
  })

  m.Run()
}
