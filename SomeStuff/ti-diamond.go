/*
Функция рисует ромб из #
*/

func printDiamond(num int) {
	fmt.Println("Мой бриллиант:")
	for i := 0; i < num*2-1; i++ {
		for j := 0; j < num*2-1; j++ {
			switch {
			case i < num && j < num && i+j+1 == num:
				fmt.Printf("#")
			case i < num && j >= num && i == j-num+1:
				fmt.Printf("#")
			case i >= num && j < num && i-num+1 == j:
				fmt.Printf("#")
			case i >= num && j >= num && i+j == 3*(num-1):
				fmt.Printf("#")
			default:
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
