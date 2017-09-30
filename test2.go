
package main

import "fmt"
import "strings"
import "math"

func main() {
   char := [17]string{" ", " ", "᠎", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "​", " ", " "}
   //charlist
   var length int = 2;
   //var pot = math.Pow(float64((len(char)+1)), length)
   //var list []string
   list := make([]string, 0)
   //output array
  // fmt.Println(cap(list))
   //fmt.Println(len(list))
   z := []string{"potato", "dont change", ""}
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
            if a == 1 {
               z[2] = "lol"
            }
            list = append(list, strings.Join(z, ""))
         }
      }
   }
   fmt.Println(list)
}
