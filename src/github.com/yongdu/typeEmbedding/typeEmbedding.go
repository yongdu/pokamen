package main
import (
        "fmt"
    )

    // define a user struct

    type user struct{
            name string
                email string
            }

   func (u *user) notify(){
        fmt.Printf("Sending user email to %s<%s>\n", u.name,u.email)

                }

   type admin struct{
               user //type embedding
               level string
                }

   func main(){
               ad := admin{
                   user : user{
                   name : "yong",
                   email: "yong@yahoo.com",
                              },
                   level:"super",
                }
         ad.user.notify()
         ad.notify()
   }
