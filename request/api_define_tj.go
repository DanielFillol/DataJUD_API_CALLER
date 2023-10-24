package request

import (
	"github.com/DanielFillol/CNJ_Validate/CNJ"
	"strings"
)

func defineTJLawsuit(cnjNumber string) (string, error) {
	isValid, err := CNJ.ValidateCNJ(cnjNumber)
	if err != nil || !isValid {
		return "", err
	}

	decomposedCNJ, err := CNJ.DecomposeCNJ(cnjNumber)
	if err != nil {
		return "", err
	}

	return strings.ToLower(decomposedCNJ.TJ), nil
}
