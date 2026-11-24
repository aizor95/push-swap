package pushswap

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseNumbers(args []string) ([]int, error) {
	var nombres []int
	vus := make(map[int]struct{})

	for _, argument := range args {
		for _, texte := range strings.Fields(argument) {
			nombre, err := strconv.Atoi(texte)
			if err != nil {
				return nil, fmt.Errorf("%q n'est pas un nombre entier", texte)
			}

			if _, existe := vus[nombre]; existe {
				return nil, fmt.Errorf("le nombre %d apparaît plusieurs fois", nombre)
			}

			vus[nombre] = struct{}{}
			nombres = append(nombres, nombre)
		}
	}

	return nombres, nil
}