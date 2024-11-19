package main

import (
	"fmt"
	"math"
	"strconv"
)

type Number struct {
	Num    int    //our number
	path   string //how did we get it
	quant  int    //quantity of symbols
	GetWay byte   //way of getting 'm'(multyplying) or 'p' (stands for plus)
}

func main() {

	fmt.Println("Number Presentation")
	n, p := NumberPresentation(115, 20, true)
	fmt.Println("your Number is presented", n, "\nnumber of symbols is ", p)

}

func (n Number) String() string {
	return fmt.Sprintf("%v=%v\ntotal length is %v, get through %c\n", n.Num, n.path, n.quant, n.GetWay)
}

/*
NumberPresentation returns symbolic presentation of given number using addition, multyplication and digits up to limit

goal - number, that we want to get
limit - we can use only numbers up to limit
show - if true, it shows the whole dp table

returns
string representation
number of symbols
*/
func NumberPresentation(goal int, limit int, show bool) (string, int) {
	goal += 1 //to make an array containing n=goal numbers
	dp := make([]Number, goal)

	if limit > goal {
		fmt.Println("limit is bigger than goal!")
		return strconv.Itoa(goal), len(strconv.Itoa(goal))
	}

	dp[0].path = ""
	dp[0].Num = 0
	dp[0].quant = 0 // zero value

	for i := 1; i < goal; i++ {
		dp[i].Num = i
		dp[i].path = ""
		dp[i].quant = math.MaxInt
		dp[i].GetWay = 'm' //m stands for multiply, how did we get our value. now it's like i*1, because multiplying takes less symbols("()")
	}

	for i := 1; i <= limit; i++ {

		dp[i].path = strconv.Itoa(i)
		dp[i].quant = len(strconv.Itoa(i))
	}

	for i := limit; i < goal; i++ {

		for j := 1; j <= i; j++ { //adding section
			if dp[i-j].quant+len(dp[j].path)+1 < dp[i].quant {
				dp[i].path = dp[j].path + "+" + dp[i-j].path  //updating the path
				dp[i].quant = dp[i-j].quant + 1 + dp[j].quant //updating the quantity of symbols
				dp[i].GetWay = 'p'                            //p stands for plus
			}
		}

		for j := 1; j <= int(math.Sqrt(float64(i))); j++ { //multyplying section
			if i%j == 0 { //without any float numbers

				switch {
				case dp[i/j].GetWay == 'p' && dp[j].GetWay == 'p' && dp[i/j].quant+len(dp[j].path)+5 <= dp[i].quant: // "<=" here is, again, because myltiplying takes less symbols
					{
						dp[i].path = "(" + dp[j].path + ")" + "*" + "(" + dp[i/j].path + ")"
						dp[i].quant = dp[i/j].quant + 5 + dp[j].quant
						dp[i].GetWay = 'm' //m stands for multiply
					}
				case dp[i/j].GetWay == 'p' && dp[j].GetWay == 'm' && dp[i/j].quant+len(dp[j].path)+3 <= dp[i].quant:
					{
						dp[i].path = dp[j].path + "*" + "(" + dp[i/j].path + ")"
						dp[i].quant = dp[i/j].quant + 3 + dp[j].quant
						dp[i].GetWay = 'm' //m stands for multiply
					}
				case dp[i/j].GetWay == 'm' && dp[j].GetWay == 'p' && dp[i/j].quant+len(dp[j].path)+3 <= dp[i].quant:
					{
						dp[i].path = "(" + dp[j].path + ")" + "*" + dp[i/j].path
						dp[i].quant = dp[i/j].quant + 3 + dp[j].quant
						dp[i].GetWay = 'm' //m stands for multiply
					}

				case dp[i/j].GetWay == 'm' && dp[j].GetWay == 'm' && dp[i/j].quant+len(dp[j].path)+1 <= dp[i].quant:
					{
						dp[i].path = dp[j].path + "*" + dp[i/j].path
						dp[i].quant = dp[i/j].quant + 1 + dp[j].quant
						dp[i].GetWay = 'm' //m stands for multiply
					}
				}

			}
		}

	}
	if show { //is show is true, we show the whole thing
		fmt.Println(dp)
	}

	return dp[goal-1].path, dp[goal-1].quant
}
