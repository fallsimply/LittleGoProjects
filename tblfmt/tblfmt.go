package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	promptScanner = bufio.NewScanner(os.Stdin)
)

func main() {
	fmt.Print(getData())
}

func getData() (colNames []string, data [][]string) {
	done := false
	cols := promptNum("Columns")
	for i := 1; i <= cols; i++ {
		col := prompt("Column " + strconv.Itoa(i))
		colNames = append(colNames, col)
	}
	for rowI := 1; !done; rowI++ {
		var row []string
		fmt.Printf("Row %d\n", rowI)

		for col := 1; col <= cols; col++ {
			fmt.Printf("\t%s? ", colNames[col - 1])
			rowS := ""
			_, _ = fmt.Scanln(&rowS)
			row = append(row, rowS)
		}
		str := ""
		_, _ = fmt.Scanln(&str)
		if (str == "exit") {
			done = true
		}
		data = append(data, row)
	}
	
	return
}

func fmtData(data [][]string, columns bool) {

}

func promptNum(title string) (num int) {
	fmt.Printf("%s? ", title)
	_, _ = fmt.Scanf("%d", &num)
	return
}

func prompt(title string) (str string) {
	fmt.Printf("%s? ", title)

	promptScanner.Scan()
	str = promptScanner.Text()

	return
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
	}
}