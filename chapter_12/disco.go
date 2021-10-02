package main

func main() {
	one()
}
func one()  {
	two()
}
func two()  {
	three()
}
func three()  {
	panic("This call stack's too deep for me.")
}
