package main
import (
    "fmt"
    /*"strconv"*/   
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
    var num int;
    //fmt.Scanf("%d", &num)
    if (x > 0) {
        fmt.Scanf("%d", &num);
        return cal_square(num) + single_case_sum(x-1);
    } 
    return 0;
}

/*display the result of each case*/
func show_cases(x int){
    var amount int;
    if(x>0) {
        fmt.Scanf("%d", &amount);
        /*fmt.Printf(strconv.Itoa(single_case_sum(amount)));*/
        show_cases(x-1);
    }    
}

/*execution*/
func main(){
    var num int;
    fmt.Scanf("%d", &num);
    show_cases(num);
}
