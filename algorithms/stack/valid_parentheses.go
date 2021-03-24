package stack

// https://leetcode-cn.com/problems/valid-parentheses/

type Stack interface {
	Append(elem interface{})
	Pop() interface{}
}

type MyStack struct {
	Elements []interface{}
	index    int
}

func (m *MyStack) Append(elem interface{}) {
	m.Elements = append(m.Elements, elem)
	m.index++
}

func (m *MyStack) Pop() interface{} {
	if m.index == 0 {
		return nil
	}
	elem := m.Elements[m.index-1]
	m.Elements = m.Elements[:m.index-1]
	m.index--
	return elem
}

func isPair(one, other rune) bool {
	switch one {
	case ')':
		return other == '('
	case '}':
		return other == '{'
	case ']':
		return other == '['
	default:
		return false
	}
}

// '('，')'，'{'，'}'，'['，']'
func IsValid(s string) bool {
	sk := &MyStack{}
	for _, word := range s {
		if elem := sk.Pop(); elem == nil {
			sk.Append(word)
		} else {
			v := elem.(rune)
			if !isPair(word, v) {
				sk.Append(v)
				sk.Append(word)
			}
		}
	}
	if sk.index == 0 {
		return true
	}
	return false
}
