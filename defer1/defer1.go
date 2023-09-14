package main

type Foo struct {
	v int
}

func MakeFoo(n *int) Foo {
	print(*n)
	return Foo{}
}

func (Foo) Bar(n *int) {
	print(*n)
}

func main() {
	var x = 1
	var p = &x
	defer MakeFoo(p).Bar(p) // line 19
	x = 2
	p = new(int) // line 21
	MakeFoo(p)
}
