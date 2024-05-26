package imports

import "fmt"


type Ticket struct{

  ID int
  event string
}

func (t Ticket) printEvent(){
  fmt.Println(t.event)
}
