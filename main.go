package main

import "fmt"

var objective_arr [2]int

type Node struct{
  // a1,a2 are the coefficients and b1 is a constant.
  a1 int
  a2 int
  b int
  next *Node
}

var head *Node = nil

var number_of_constrains = 0


func add_objective(){
  fmt.Printf("c1x1 + c2X2 = Profit \n")  
  fmt.Printf("where c1,c2,…,cn are the coefficients (profits) for each decision variable.\n")
  fmt.Printf("enter c1 \n")  
  fmt.Scanln(&objective_arr[0])
  fmt.Printf("enter c2 \n")  
  fmt.Scanln(&objective_arr[1])

}

func add_constrain(){

  var new_constrain Node

  fmt.Printf("a1x1 + a2x2 ≤ b")
  fmt.Printf("where a1,a2 are the coefficients and b is a constant.\n")
  fmt.Printf("enter a1 \n")  
  fmt.Scanln(&new_constrain.a1)
  fmt.Printf("enter a2 \n")  
  fmt.Scanln(&new_constrain.a2)
  fmt.Printf("enter b \n")  
  fmt.Scanln(&new_constrain.b)

  if head == nil {
    head = &new_constrain
  }else{
    current := head
    for current.next != nil{
      current = current.next
    }
    current.next = &new_constrain
  }

  number_of_constrains++

}

func simplex_method(){

} 

func graphical_method(){

}

func debug_out(){
  fmt.Println("debug")
  
  fmt.Printf("objective: %d, %d \n", objective_arr[0], objective_arr[1])

  current := head
  fmt.Printf(" a1 from head: %d \n", current.a1)

  for current.next != nil{
    current = current.next
    fmt.Printf(" a1 from head: %d \n", current.a1)
  }

  fmt.Printf("number_of_constrains %d \n", number_of_constrains)


}

func main(){

  fmt.Printf("Max profit, simplivied version 0.1 \n\n")  
  fmt.Printf("<><><><><><><><><><><><><><><><><><><><>\n")  
  fmt.Printf("Example of Problem format: \nMaximize: c1*x1 + c2*x2 \nSubject to:\n  a1*x1 + a2*x2 ≤ b \n  ...\n  (possibly more constraints))\n")
  fmt.Printf("<><><><><><><><><><><><><><><><><><><><>\n")  
  fmt.Printf("1. What is an objective? \n")  
  add_objective()
  fmt.Printf("2. What are the constants? \n")  
  var loop = true
  for loop {
    add_constrain()
    var choose string
    for choose != "y" && choose != "n"{
      fmt.Printf("add another constrain? (y/n) \n")  
      fmt.Scanf("%s", &choose)
    }
    if choose == "n"{
      loop = false 
    }
  }

  debug_out()

  fmt.Printf("3. Right now there is only graphical method implemented: \n")

  // TODO: 
  // 1. Consider that constrains may be "<" or ">"
  // 2. Implement graphical method

}
