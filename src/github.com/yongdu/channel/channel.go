package main 

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// wg is used to wait for the program to finish
var wg sync.WaitGroup

func init(){
    rand.Seed(time.Now().UnixNano())
}

// main is the entry point for all Go program.

func main(){
    // create an unbuffered channel
    court := make(chan int)
    
    //add a count of two, one for each goroutine
    wg.Add(2)

    go player("Yong",court)
    go player("Bo",court)

    // start the set
    court <- 1
    // wait for the game to finish
    wg.Wait()
}
    // player simulates a person playing the game of tennis

    func player (name string, court chan int) {
        defer wg.Done()

        for {
        //wait for the ball to be hit back to us
        ball,ok:=<-court
        if !ok {
        // if the channel was closed, we won
        fmt.Printf("Player %s Won\n",name)
        return
        }
        // pick a random number and see if we miss the ball

        n := rand.Intn(100)
        if n%13 ==0 {
        fmt.Printf("Player %s Missed\n",name)

        //close the channel to signal we lost
        close(court)
        return
        }

        fmt.Printf("Player %s hit %d\n",name,ball)
        ball ++

        // hit the ball´back to the opposing player
        court <- ball

        
    }

}
