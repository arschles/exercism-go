package protein

// TODO:
// - comment the public symbols (var/functions, etc...)
// - implement FromRNA as a loop, not recursively

// HEY!
// We did this on a twitch stream, and one of the people watching
// was a COMPUTATIONAL BIOLOGIST. BOOYAH

import "errors"

var (
	ErrStop        error = errors.New("STOP")
	ErrInvalidBase error = errors.New("Invalid")
)

func FromCodon(str string) (string, error) {

	codons := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCG": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}

	protein, ok := codons[str]
	if !ok {
		return "", ErrInvalidBase
	} else if protein == "STOP" {
		return "", ErrStop
	} else {
		return protein, nil
	}
}

func FromRNA(str string) ([]string, error) {
	if len(str) < 3 {
		return nil, ErrInvalidBase
	}
	curCodon := str[0:3]
	protein, err := FromCodon(curCodon)
	if err == ErrStop {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	restOfString := str[3:]
	if restOfString == "" {
		return []string{protein}, nil
	}
	remainingProteins, err := FromRNA(str[3:])
	if err == ErrStop {
		return nil, nil
	} else if err != nil {
		return []string{protein}, err
	}

	return append([]string{protein}, remainingProteins...), nil
}
