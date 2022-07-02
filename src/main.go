package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Cloud Wizard")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Informe o nome do Projto: ")
	V_Projeto, _ := reader.ReadString('\n')
	fmt.Println(V_Projeto)
}
