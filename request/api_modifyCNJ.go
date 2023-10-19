package request

import (
	"github.com/Darklabel91/CNJ_Validate/CNJ"
)

func modifyCNJ(cnjNumber string) (string, error) {
	decomposed, _ := CNJ.DecomposeCNJ(cnjNumber)
	newCNJ := decomposed.LawsuitNumber + decomposed.VerifyingDigit + decomposed.ProtocolYear + decomposed.Segment + decomposed.Court + decomposed.SourceUnit
	return newCNJ, nil
}
