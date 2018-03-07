package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main(){
    runtime.GOMAXPROCS(1)
    var wg sync.WaitGroup
    wg.Add(2)
    fmt.Println("Start Goroutines")

    // decleare an anonymous func and create a goroutine
    go func(){
        defer wg.Done()

        // Display the alphbet three times
        for count :=0; count<3;count++ {
            for char :='a'; char <'a'+26; char++ {
                fmt.Printf("%c ", char)
            }
        }

    }()

    // declare an anoymous function and create a goroutine

    go func(){
        defer wg.Done()
        //display the alphabet three times
        for count := 0; count < 3; count ++ {
            for char  :='A'; char <'A'+26; char ++{
                fmt.Printf("%c ",char )
            }
        }
    }()

    fmt.Println("Waiting to Finish")
    wg.Wait()
    fmt.Println("\nTermianting Program")


}
