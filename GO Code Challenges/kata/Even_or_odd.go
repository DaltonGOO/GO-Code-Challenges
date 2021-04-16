package kata
import ("fmt")
func EvenOrOdd(number int) string {
  // your code here
      fmt.Scanln(&number)
    if(number%2==0){
        return("Even")
    }else{
        return("Odd")
    } 
}