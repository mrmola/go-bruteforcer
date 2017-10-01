
package main

import "fmt"
import "strings"
import "math"

func main() {
   char := [3]string{"a", "b", "c"}
   //charlist
   var length int = 4;
   var count int = 6;
   //var pot = math.Pow(float64((len(char)+1)), length)
   //var list []string
   list := make([]string, 0)
   //output array
  // fmt.Println(cap(list))
   //fmt.Println(len(list))
   z := []string{"", "dont change", "", ""}
   for i := 0; i < len(char); i++ {
      z[0] = char[i]
      //fmt.Println(list)
      list = append(list, z[0])
   }
   //fmt.Println(cap(list))
   var lengt float64 = float64(len(char))
   for a := 0; a < length; a++ {
      for i := 0; i < int(math.Pow(float64(lengt), float64(a))); i++ {
         for j := 0; j < len(char); j++ {
            z[0] = char[j]
            z[1] = list[i]
//          fmt.Println(list)
            if a == length-1 {
               z[2] = "lol"
            }
            if len(list)+1 > count {
               break
            }
            list = append(list, strings.Join(z, ""))
         }
      if len(list)+1 > count {
        break
      }
      }
      if len(list)+1 > count {
         break
      }
   }
   fmt.Println(list)
}
