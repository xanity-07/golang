package intermediate

// import "fmt"

// type stack[T any] struct {
// 	elements []T
// }

// func (s *stack[T]) push(item T) {
// 	s.elements = append(s.elements, item)
// }

// func (s *stack[T]) pop() (T, bool) {
// 	if len(s.elements) == 0 {
// 		var zero T
// 		return zero, false
// 	}

// 	length := len(s.elements)
// 	lastEl := s.elements[length-1]
// 	s.elements = s.elements[0 : length-1]
// 	return lastEl, true

// }

// func (s *stack[T]) isEmpty() bool {
// 	return len(s.elements) == 0
// }

// func (s stack[T]) printAll() {
// 	if len(s.elements) == 0 {
// 		fmt.Println("Stack is empty.")
// 		return
// 	}
// 	fmt.Print("Stack elements: ")
// 	for _, el := range s.elements {
// 		fmt.Print(el)
// 	}
// 	fmt.Println()
// }

// func IntroGenerics() {
// 	s := stack[int]{
// 		elements: []int{},
// 	}
// 	s.push(1)
// 	s.push(4)
// 	// s.push("hello") //* Wont work because its contraint to int
// 	fmt.Println(s.pop())
// 	fmt.Println(s.elements)
// 	fmt.Println(s.isEmpty())
// 	s.printAll()
// }
