package main
import(
"fmt"
"github.com/yongdu/exportingAndUnexporting/counters"

)

func main(){
   
    
    counter := counters.alertCounter(10)
    fmt.Printf("Counter:%d\n", counter)

}
