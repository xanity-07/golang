package intermediate

// import "fmt"

// func IntroClosures() {
// 	// secuence := adder()
// 	// result := secuence()
// 	// fmt.Println(result)

// 	daysRemaining := func() func(int) int {
// 		days := 10
// 		fmt.Println("Days left:", days)
// 		return func(x int) int {
// 			days -= x
// 			if days < 0 {
// 				fmt.Println("Days left are already 0")
// 				return -1
// 			}
// 			fmt.Printf("%d Days left.\n", days)
// 			return days
// 		}
// 	}

// 	oneDayLess := daysRemaining()
// 	oneDayLess(1)
// 	oneDayLess(1)

// }

// // func adder() func() int {
// // 	i := 0
// // 	fmt.Println("Previous value of i:", i)
// // 	return func() int {
// // 		i++
// // 		fmt.Println("Added 1 to i")
// // 		return i
// // 	}
// // }
