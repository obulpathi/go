// The standard library's `strings` package provides many
// useful string-related functions. Here are some examples
// to give you a sense of the package.

package main

import s "strings"
import "fmt"

func main() {

    // Here's a sample of the functions available in
    // `strings`. Note that these are all functions from
    // package, not methods on the string object itself.
    // This means that we need pass the string in question
    // as the first argument to the function.
    fmt.Println("Contains:  ", s.Contains("test", "es"))
    fmt.Println("Count:     ", s.Count("test", "t"))
    fmt.Println("HasPrefix: ", s.HasPrefix("test", "te"))
    fmt.Println("HasSuffix: ", s.HasSuffix("test", "st"))
    fmt.Println("Index:     ", s.Index("test", "e"))
    fmt.Println("Join:      ", s.Join([]string{"a", "b"}, "-"))
    fmt.Println("Repeat:    ", s.Repeat("a", 5))
    fmt.Println("Replace:   ", s.Replace("foo", "o", "0", -1))
    fmt.Println("Replace:   ", s.Replace("foo", "o", "0", 1))
    fmt.Println("Split:     ", s.Split("a-b-c-d-e", "-"))
    fmt.Println("ToLower:   ", s.ToLower("TEST"))
    fmt.Println("ToUpper:   ", s.ToUpper("test"))
    fmt.Println()

    // You can find more functions in the [`strings`](http://golang.org/pkg/strings/)
    // package docs.

    // Not part of `strings` but worth mentioning here are
    // the mechanisms for getting the length of a string
    // and getting a character by index.
    fmt.Println("Len: ", len("hello"))
    fmt.Println("Char:", "hello"[1])
}
