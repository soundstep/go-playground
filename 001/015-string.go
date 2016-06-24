package main

import (
    "fmt"
    "strings"
    "sort"
)

func main() {

    str := "Hello World"

    fmt.Println(strings.Contains(str, "lo"))
    fmt.Println(strings.Index(str, "lo"))
    fmt.Println(strings.Count(str, "l"))
    fmt.Println(strings.Replace(str, "l", "x", 3))

    csvString := "1,2,3,4,5,6"
    fmt.Println(strings.Split(csvString, ","))

    listOfLetters := []string{"c", "a", "b"}
    fmt.Println(listOfLetters)

    listOfNums := strings.Join([]string{"3", "2", "1"}, " - ")
    fmt.Println(listOfNums)

    sort.Strings(listOfLetters)
    fmt.Println("Letters:", listOfLetters)

}
