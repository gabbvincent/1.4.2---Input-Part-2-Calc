package main

import "fmt"

func main() {
  //Declare a variable for two numbers

  var numberOne int 
  var numberTwo int

  //Asker user to enter two numbers

  fmt.Println("Enter a whole number")

  //scan the line above for numberOne

  fmt.Scanln(&numberOne)

  fmt.Println("Enter another whole number")

  //scan the line above for numberTwo

  fmt.Scanln(&numberTwo)

  //declare a variable each math operator using numberOne and numberTwo.

  var plus int = (numberOne + numberTwo)

  var minus int = (numberOne - numberTwo)

  var multiply int = (numberOne * numberTwo)

  var devide int = (numberOne / numberTwo)

  //display a message with the whole math operator for each one defined above.

  fmt.Println(numberOne ,"+" , numberTwo , "=" , plus)

  fmt.Println(numberOne ,"-" , numberTwo , "=" , minus)

  fmt.Println(numberOne ,"*" , numberTwo , "=" , multiply)

  fmt.Println(numberOne ,"/" , numberTwo , "=" , devide)

//end.
}