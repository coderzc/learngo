package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	i, j int
}

type Queue struct {
	length int
	data   []interface{}
}

func (q *Queue) offer(i interface{}) {
	q.data = append(q.data, i)
	q.length++
}

func (q *Queue) poll() interface{} {
	pollVal := q.data[0]
	q.length--
	q.data = q.data[1:]
	return pollVal
}

func printMaze(maze [][]int) {
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func readMaze(filename string) [][]int {
	var row, col int
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fscanf(file, "%d %d", &row, &col)

	buf := bufio.NewReader(file)
	var maze = make([][]int, row)
	for i := 0; i < row; i++ {
		line, _ := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		tmpSlice := strings.Split(line, " ")
		for j := 0; j < col; j++ {
			val, _ := strconv.Atoi(tmpSlice[j])
			maze[i] = append(maze[i], val)
		}
	}
	return maze
}

func bfs(maze [][]int, start point, end point) {
	steps := make([][]int, len(maze))
	for i := range maze {
		steps[i] = make([]int, len(maze[i]))
	}

	var queue = new(Queue)
	for queue.length > 0 {
		val := queue.poll().(point)
		steps[val.i][val.j] = 1

	}

}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("终于捕获到了panic产生的异常：", err) // 这里的err其实就是panic传入的内容
			fmt.Println("我是defer里的匿名函数，我捕获到panic的异常了，我要recover，恢复过来了。")
		}
	}()
	maze := readMaze("./maze.txt")
	printMaze(maze)

	bfs(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
}
