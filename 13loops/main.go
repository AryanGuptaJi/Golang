package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// Join the elements with a comma and a space
	//joinedDays := strings.Join(days, ", ")
	//fmt.Println(joinedDays) // Output: Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday


	//1.
	for index, day := range days {
		fmt.Printf("index is %v, day is %v\n", index, day)
	}

	
	// fmt.Println(days)
	//[Sunday Monday Tuesday Wednesday Thursday Friday Saturday]


	//2.
	for i := range days {
		fmt.Println(days[i])
	}
	
	//3.
	// for d := 0; d < len(days); d++ {
	// 	fmt.Printf("%s", days[d])

	// 	// fmt.Printf(days[d])

	// 	// Add a comma after each element except the last one
	// 	if d < len(days)-1 {
	// 		fmt.Print(", ") // Note the space after the comma
	// 	}
	// }



	//Break 
	rogueValue := 1
	for rogueValue < 10 {

		if rogueValue == 5 {
			break
		}

		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

	//Continue 
	rogueValu := 1
	for rogueValu < 10 {

		if rogueValu == 2 {
			goto loc
		}
		if rogueValu == 5 {
			rogueValu++
			continue
		}

		fmt.Println("Value is: ", rogueValu)
		rogueValu++
	}

	loc:{
		fmt.Println("Jumping")
	}
}
