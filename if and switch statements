package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())

  isHeistOn:=true
  
  eludedGuards:=rand.Intn(100)
  fmt.Println(eludedGuards)
  if eludedGuards >= 50{
    fmt.Println("Looks good ,But be careful")
  } else  {
    fmt.Println("disguise and move on")
  }

  openedVault:= rand.Intn(100)
  fmt.Println(openedVault)

  if  isHeistOn && openedVault >=70{
    fmt.Println("Grab and Go")
  } else if isHeistOn {
    
    fmt.Println("valut can't be opened")
  } 
  leftSafely:= rand.Intn(5)
  fmt.Println(leftSafely)
  if isHeistOn{
    switch leftSafely{
      case 0:
      isHeistOn =false 
      fmt.Println("Looks like you tripped an alarm...run run..")

      case 1:
      isHeistOn=false
      fmt.Println("looks like vault need to be opened from inside....")

      case 2:
      isHeistOn=false
      fmt.Println("attack the guards...")

      case 3:
      isHeistOn=false
      fmt.Println("hide somewhere....")
      
      default:
      fmt.Println("Start the getaway car !")
    }
  }

  if isHeistOn == true {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Println("amount looted",amtStolen)
  }
  fmt.Println("Heist status :",isHeistOn)

}
