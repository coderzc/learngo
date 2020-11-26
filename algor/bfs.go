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

func (p point) move(r point) point {
	return point{p.i + r.i, p.j + r.j}
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
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}

func (p *point) at(grid [][]int) (val int, ok bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

var dirs = []point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
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

func appendIndex(slice []point, value point, i int) []point {
	slice = append(slice, point{})
	copy(slice[i+1:], slice[i:])
	slice[i] = value
	return slice
}

func bfs(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range maze {
		steps[i] = make([]int, len(maze[i]))
	}

	var queue = new(Queue)
	queue.offer(start)
	//steps[start.i][start.j] = 0
	for queue.length > 0 {
		cur := queue.poll().(point)

		// 到达终点
		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.move(dir)

			// 碰墙
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			// 已经走过了
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			// 回到原点
			if next == start {
				continue
			}
			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1

			queue.offer(next)
		}
	}

	return steps
}

func genPath(steps [][]int, start point, end point) []point {
	if endVal, _ := end.at(steps); endVal == 0 {
		return nil
	}
	var path []point
	cur := end
	for cur != start {
		path = appendIndex(path, cur, 0)
		curVal, _ := cur.at(steps)
		for _, dir := range dirs {
			next := cur.move(dir)
			val, _ := next.at(steps)
			if val == curVal-1 {
				cur = next
				break
			}
		}
	}
	path = appendIndex(path, start, 0)
	return path
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("终于捕获到了panic产生的异常：", err) // 这里的err其实就是panic传入的内容
			fmt.Println("我是defer里的匿名函数，我捕获到panic的异常了，我要recover，恢复过来了。")
		}
	}()
	maze := readMaze("./maze.txt")
	//printMaze(maze)

	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}
	steps := bfs(maze, start, end)

	printMaze(steps)

	// 走到终点需要多少步
	step, _ := end.at(steps)
	fmt.Printf("move step:%d\n", step)

	fmt.Println("path: ", genPath(steps, start, end))
}
