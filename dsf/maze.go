package dsf

import "fmt"

type MazeSolver struct {
	maze               [][]int
	visited            [][]bool
	startRow, startCol int
}

func NewMazeSolver(maze [][]int, startRow, startCol int) *MazeSolver {

	var visited = make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[i]))
	}
	return &MazeSolver{
		maze:     maze,
		visited:  visited,
		startRow: startRow,
		startCol: startCol,
	}
}

func (m *MazeSolver) FindDestination() {
	if m.dfs(m.startRow, m.startCol) {
		fmt.Println("Found destination")
	} else {
		fmt.Println("Destination not found")
	}
}

func (m *MazeSolver) dfs(x, y int) bool {
	fmt.Println("Analyzing ", x, y)
	if x == (len(m.maze)-1) && y == (len(m.maze[x])-1) {
		return true
	}

	if m.isFeaseable(x, y) {
		m.visited[x][y] = true
		// go down
		if m.dfs(x+1, y) {
			return true
		}

		// go right
		if m.dfs(x, y+1) {
			return true
		}

		// go up
		if m.dfs(x-1, y) {
			return true
		}

		// go left
		if m.dfs(x, y-1) {
			return true
		}

		m.visited[x][y] = false

	}

	return false
}

func (m *MazeSolver) isFeaseable(x, y int) bool {
	if x < 0 || x > len(m.maze)-1 {
		return false
	}

	if y < 0 || y > len(m.maze[x])-1 {
		return false
	}

	fmt.Println(x, y)
	if m.visited[x][y] {

		return false
	}

	return m.maze[x][y] != 1
}
