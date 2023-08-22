package metrics

import (
	"strconv"
	"strings"
)

func Parse(valueStr string) (float64, string) {

	parts := strings.Fields(valueStr)

	// Substitui a vírgula por ponto para ter um formato numérico válido
	valueStr = strings.Replace(parts[0], ",", ".", 1)

	// Converte para float64
	valueFloat, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return 0.0, parts[1]
	}

	return valueFloat, parts[1]
}
