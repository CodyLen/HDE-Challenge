package main
import (
    "fmt"
    "strconv"   
)

/*function for calculating sqaure*/
func cal_square(x int) int{
    var sqr int;
    if (x >= 0) {
        sqr = x*x;
    } else {
         sqr=0;
    }
    return sqr;
}

/*sum up the single case. Due to the requirements of the test, the recursive method is applied*/
func single_case_sum(x int) int{
    var num,sum int;
    fmt.Scanf("%d", &num)
    if (x > 0) {
        sum = cal_square(num) + single_case_sum(x-1);
    } else {   
        sum = 0;      
    }
    return sum;
}

/*display the result of each case*/
func cases(x int){
    var amount int;
    fmt.Scanf("%d", &amount);
    fmt.Printf(strconv.Itoa(single_case_sum(amount)));
    if (x > 0) {
        cases(x-1);
    } 
}

/*execution*/
func main(){
    var num int;
    fmt.Scanf("%d", &num);
    cases(num);
}
