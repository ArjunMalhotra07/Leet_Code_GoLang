package arrays

import "fmt"

func FloodFillHelper() {
	fmt.Println("PROGRAM 2 : Change Color in the Matrix ")
	arr := [][]int{{0, 0, 0}, {0, 1, 1}}
	arr2 := [][]int{{1, 1, 1, 1}, {0, 1, 0, 0}, {0, 1, 0, 1}}
	loopExamplesFunc(arr)
	loopExamplesFunc(arr2)
	fmt.Println(" ")
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	presentColor := image[sr][sc]
	if presentColor == newColor {
		return image
	}
	columns := len(image[0])
	rows := len(image)
	return changeColor(image, sr, sc, newColor, columns, rows, presentColor)
}

func changeColor(arr [][]int, sr, sc, newColor int, columns, rows, presentColor int) [][]int {
	if arr[sr][sc] == presentColor {
		arr[sr][sc] = newColor
	}
	if sc-1 != -1 {
		if arr[sr][sc-1] == presentColor {
			changeColor(arr, sr, sc-1, newColor, columns, rows, presentColor)
		}
	}
	if sc+1 < columns {
		if arr[sr][sc+1] == presentColor {
			changeColor(arr, sr, sc+1, newColor, columns, rows, presentColor)
		}
	}
	if sr-1 != -1 {
		if arr[sr-1][sc] == presentColor {
			changeColor(arr, sr-1, sc, newColor, columns, rows, presentColor)
		}
	}
	if sr+1 < rows {
		if arr[sr+1][sc] == presentColor {
			changeColor(arr, sr+1, sc, newColor, columns, rows, presentColor)
		}
	}
	return (arr)
}

func printLoopFunc(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Println()
	}
}

func loopExamplesFunc(arr [][]int) {
	f := fmt.Println
	f("Input 2D Array")
	printLoopFunc(arr)
	f()
	ansArray := floodFill(arr, 1, 1, 2)
	f("Altered 2D Array")
	printLoopFunc(ansArray)
	f()

}
