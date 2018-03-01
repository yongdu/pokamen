package main

import (
    "fmt"
)

// define a struct 

type user struct{
    name  string
    email string
}

// notify -- a method 
func (u user)notify(){
    fmt.Printf("Sending User Email To %s<%s>\n",
            u.name, 
            u.email)
}

func (u *user) changeEmail(email string){
    u.email = email
}

func main(){
    bill := user{"Bill","bill@email.com"}
    bill.notify()
    lisa := &user{"Lisa","lisa@email.com"}
    lisa.notify()

    bill.changeEmail("bill@newdomain.com")
    bill.notify()
    lisa.changeEmail("lisa@newdomain.com")
    lisa.notify()
}


