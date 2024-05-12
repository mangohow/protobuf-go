package internal_gengo

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	names := []string{
		"my_variable_name",
		"myVariableName",
		"MyVariableName",
	}
	for _, name := range names {
		fmt.Println("Original:", name)
		fmt.Println("Camel Case:", ToCamelCase(name))
		fmt.Println("Pascal Case:", ToPascalCase(name))
		fmt.Println("Snake Case:", ToSnakeCase(name))
		fmt.Println("----------------------------------")
	}

}
