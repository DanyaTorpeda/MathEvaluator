package stack

// Stack структура для стека
type Stack struct {
	items []string
}

func NewStack() *Stack {
	return &Stack{}
}

// Push добавляет элемент в стек
func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

// Pop удаляет верхний элемент из стека и возвращает его
func (s *Stack) Pop() (string, bool) {
	if len(s.items) == 0 {
		return "", false // Пустой стек
	}
	// Удаляем последний элемент
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

// Peek возвращает верхний элемент без удаления
func (s *Stack) Peek() (string, bool) {
	if len(s.items) == 0 {
		return "", false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
