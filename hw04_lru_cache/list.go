package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	l_len       int
	front, back *ListItem
}

func NewList() List {
	return new(list)
}

func (l list) Len() int {
	return l.l_len
}

func (l list) Front() *ListItem {
	if l.l_len == 0 {
		return nil
	}

	return l.front
}

func (l list) Back() *ListItem {
	if l.l_len == 0 {
		return nil
	}

	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newPointer := &ListItem{
		Value: v,
		Prev:  nil,
		Next:  nil,
	}

	if l.l_len == 0 {
		l.front = newPointer
		l.back = newPointer
	} else {
		old := l.front
		l.front = newPointer
		old.Prev = l.front
		l.front.Next = old
	}

	l.l_len = l.Len() + 1
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	newPointer := &ListItem{
		Value: v,
		Prev:  nil,
		Next:  nil,
	}

	if l.l_len == 0 {
		l.front = newPointer
		l.back = newPointer
	} else {
		old := l.back
		l.back = newPointer
		l.back.Prev = old
		old.Next = l.back
	}

	l.l_len = l.Len() + 1
	return l.front
}

func (l *list) Remove(i *ListItem) {

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	l.l_len = l.Len() - 1
}

func (l *list) MoveToFront(i *ListItem) {

	if (l.l_len < 2) || (i.Prev == nil) { // если элемент единственный или первый, то ничего не делаем
		return
	}

	//Меняем следующий элемент для элемента, предшествующего i
	i.Prev.Next = i.Next

	if i.Next == nil { // Если перемещаемый элемент - последний
		l.back = i.Prev
	} else { // Если не последний, то переназначем указатели
		i.Next.Prev = i.Prev
	}

	//теперь переназначаем указатели для нового первого узла
	i.Next = l.front
	i.Prev = nil
	l.front = i
}
