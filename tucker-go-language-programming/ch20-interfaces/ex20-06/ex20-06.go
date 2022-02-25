package main

type Attacker interface {
	Attack()
}

func main() {
	var attacker Attacker
	attacker.Attack()
}
