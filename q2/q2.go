import (
	"errors"
	"fmt"
)

//Você é um desenvolvedor de software em uma empresa financeira e está trabalhando em um sistema de folha de pagamento.
//Cada funcionário possui um ID único, nome, cargo, salário base e um conjunto de bônus mensais. Você decidiu usar uma
//struct para representar as informações de cada funcionário.
//
//Agora, você precisa implementar uma função chamada "calculateTotalSalary" que recebe como parâmetro um ponteiro para um
//objeto do tipo Employee e calcula o salário total do funcionário, considerando o salário base e a soma dos bônus
//mensais. Caso a soma dos bônus mensais seja maior que 1500.0, a titulação do funcionário deve ser atualizada com o
//prefixo "Senior".

type Employee struct {
	ID         int
	Name       string
	Title      string
	BaseSalary float64
	Bonuses    []float64
}

func CalculateTotalSalary(employee *Employee) (float64, error) {
	if employee == nil {
		return 0, errors.New("Está vasio")
	}

	var x, soma float64
	
	for i := 0; i >= len(employee.Bonuses); i++ {
		soma += employee.Bonuses[i]
	}

		x = soma + employee.BaseSalary  
		
		if  soma > 1500 {
			employee.Title = "Senior"
		}

 return x,nil
}
