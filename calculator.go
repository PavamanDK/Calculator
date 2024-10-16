package main
import "fmt"

func main() {
    var operator string
    fmt.Println("Enter operator: +, -, *, /")
    fmt.Scan(&operator)

    var num1, num2 float64
    fmt.Println("Enter two numbers:")
    fmt.Scan(&num1, &num2)
    fmt.Println(calc(num1,num2,operator))

    /*switch operator {
    case "+":
        fmt.Println(num1 + num2)
    case "-":
        fmt.Println(num1 - num2)
    case "*":
        fmt.Println(num1 * num2)
    case "/":
        fmt.Println(num1 / num2)
    default:
        fmt.Println("Invalid operator")
    }*/
}

func calc(num1 float64, num2 float64, operator string ) float64 {
   var num3 = 0.0
   switch operator {
    case "+":
        num3 = num1 + num2
    case "-":
        num3 = num1 - num2
    case "*":
        num3 = num1 * num2
    case "/":
        num3 = num1 / num2
    default:
        fmt.Println("Invalid operator")
    }

  return num3
}
