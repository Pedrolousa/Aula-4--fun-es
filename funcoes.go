package main
import "fmt"
func somar (n1 int8, n2 int8) int8 {

	return n1 + n2
}


	func calculosmatematicos(n1, n2 int8) (int8, int8){

		soma := n1 + n2
		subtracao := n1 - n2
		return soma, subtracao
	}





func main() {

	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("texto da função 1")
	fmt.Println(resultado)

	
	resultadosoma, _ := calculosmatematicos(10, 15)
	fmt.Println(resultadosoma)

	


}