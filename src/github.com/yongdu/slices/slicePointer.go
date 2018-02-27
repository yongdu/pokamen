
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
        for index, value := range slice {
            fmt.Printf("Index: %d Value: %d\n", index, value)
        }
        for index, value := range slice{
            // the address of value is same after each iterate 
            fmt.Printf("Value: %d Value-Addr: %X ElementAddress: %X \n",value,&value,&slice[index])
        }
        dict := map[string]string{"Red":"#da117","Orange":"#e95a22"}
        // get value based on the key
        value,exists := dict["Blue"]

        if exists {
            fmt.Println(value)
        }

        value1:= dict["Red"]
        if value1!= ""{
            fmt.Println(value1)
        }
}

