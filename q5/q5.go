package q5

import "fmt"

//Você está desenvolvendo um sistema de gerenciamento de vendas para uma loja. Cada venda possui um código único, data,
//valor total e uma lista de produtos vendidos. Cada produto tem um código, nome e preço unitário. Você decidiu usar
//structs para representar as informações de cada venda e de cada produto.
//
//Agora, você precisa implementar uma função chamada "calculateTotalSales" que recebe como parâmetro um mapa onde as
//chaves são os códigos das vendas e os valores são ponteiros para objetos do tipo Sale. A função deve calcular o valor
//total de todas as vendas.
//
//Sua tarefa é escrever a função "calculateTotalSales" e usá-la para calcular o valor total de todas as vendas no mapa
//fornecido.

type Product struct {
	Code  string
	Date  int
	Name  string
	Price float64
}

type Sale struct {
	Code     string
	Products []*Product
}

func CalculateTotalSales(sales map[string]*Sale) float64 {
	total := 0.0
	for _, Sale := range sales {
		sum := 0.0
		for _, Product := range Sale.Products {
			sum += Product.Price
		}
		total += sum
	}
	return total
}

func main() {
	sale1 := &Sale{
		Code: "S001",
		//Date: "2023-05-01",
		Products: []*Product{
			&Product{Code: "P001", Name: "Product 1", Price: 50.0},
			&Product{Code: "P002", Name: "Product 2", Price: 25.0},
			&Product{Code: "P003", Name: "Product 3", Price: 25.0},
		},
	}

	sale2 := &Sale{
		Code: "S002",
		//Date: "2023-05-05",
		Products: []*Product{
			&Product{Code: "P004", Name: "Product 4", Price: 25.0},
			&Product{Code: "P005", Name: "Product 5", Price: 50.0},
		},
	}

	sales := map[string]*Sale{
		"S001": sale1,
		"S002": sale2,
	}

	total := CalculateTotalSales(sales)
	fmt.Println(total)

}
