package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10, 20))                                                        // => 30
	fmt.Println(sum(10, 20, 30, 40))                                                // => 100
	fmt.Println(sum(10, 20, 30, "40"))                                              // => 100
	fmt.Println(sum(10, 20, 30, "40", "abc"))                                       // => 100
	fmt.Println(sum(10, 20, []int{30, 40}))                                         // => 100
	fmt.Println(sum(10, 20, []interface{}{30, 40, []int{50, 60}}))                  // => 210
	fmt.Println(sum(10, 20, []interface{}{30, 40, []interface{}{50, "60", "abc"}})) // => 210
}

//My solution
func sum(nos ...interface{}) int {
	//var x interface{}
	result := 0
	for _, no := range nos {
		fmt.Println("no : ", no)
		switch val := no.(type) {

		case int:
			result += val
		case string:
			if i, err := strconv.Atoi(val); err == nil {
				result += i
			}
		case []int:
			// for _, num := range val {
			// 	result += num
			// }

			//If not want to repeat code .. create a interface slice with values
			intValueList := make([]interface{}, len(val))
			for i, v := range val {
				intValueList[i] = v
			}
			fmt.Println("inside []int ", intValueList)
			result += sum(intValueList...)
		case []interface{}:
			fmt.Println("inside []interface")
			// for _, num := range val {
			// 	result += sum(num)
			result += sum(val...)
		default:
			continue
		}
	}
	return result
}
