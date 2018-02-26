
package main
import "fmt"
func main(){
        slice := []int{1,2,3,4,5}
        newslice := slice[1:3]// slice[i:j], the length is j-i
        newslice2 := slice[2:4]
        fmt.Println(slice)
        fmt.Println(newslice,newslice2)
        newslice[1]= 10

        fmt.Println(slice)
        fmt.Println(newslice,newslice2)

}

