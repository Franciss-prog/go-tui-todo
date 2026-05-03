package main

type todo struct {
	task     []string
	cursor   int
	done     bool
	selected map[int]struct{}
}

func initTodo() todo {
	return todo{
		task:     []string{},
		cursor:   0,
		done:     false,
		selected: make(map[int]struct{}),
	}
}
