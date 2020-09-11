package main
import "fmt"

// Update marsYear so that it takes earthYears
// As a parameter
func computeMarsYears(earthYears,x,y int) int {
  // Remove earthYears definition within marsYear
 
  
  earthDays := earthYears * x
  marsYears := earthDays / y
  return marsYears
}

func main() {
  myAge := 25
  x:=365
  y:=687
  // Call `marsYear` with `myAge`
  myMartianAge := computeMarsYears(myAge,x,y)
  fmt.Println(myMartianAge)
}
