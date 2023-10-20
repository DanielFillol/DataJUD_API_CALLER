package request

import (
	"github.com/DanielFillol/CNJ_Validate/CNJ"
	"strings"
)

func defineTJ(cnjNumber string) (string, error) {
	decomposedCNJ, err := CNJ.DecomposeCNJ(cnjNumber)
	if err != nil {
		return "", err
	}

	return strings.ToLower(decomposedCNJ.TJ), nil
}
