package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	result := fmt.Sprintf("Имя: %s\n", p.Name) +
		fmt.Sprintf("Вес: %.2f кг.\n", p.Weight) +
		fmt.Sprintf("Рост: %.2f м.", p.Height)

	fmt.Println(result)
}
