package main

func main() {

	n := 10
	println(read(&n))
}

func read(p *int) (v int) {
	v = *p
	return
}
