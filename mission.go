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
    //fmt.Scanf("%d", &num)
    if (x > 0) {
        fmt.Scanf("%d", &num);
        sum = cal_square(num) + single_case_sum(x-1);
    } else {   
        sum = 0;      
    }
    return sum;
}

/*display the result of each case*/
func show_cases(x int){
    var amount,result int;
    if(x>0) {
        fmt.Scanf("%d", &amount);
        result = single_case_sum(amount);
        fmt.Printf(strconv.Itoa(result));
        show_cases(x-1);
    }    
}

/*execution*/
func main(){
    var num int;
    fmt.Scanf("%d", &num);
    show_cases(num);
}
