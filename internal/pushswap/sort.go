package pushswap

func Generate(numbers []int) []string {
	p := Piles{a: rangs(numbers)}
	p.trier()
	return append([]string(nil), p.operations...)
}

func (p *Piles) trier() {
	if p.aEstTriee() {
		return
	}
	if len(p.a) <= 7 {
		p.trierOptimalPetit()
		return
	}
	p.trierRadix()
}

func (p *Piles) trierRadix() {
	maxRang := len(p.a) - 1
	nombreDeBits := 0

	for maxRang > 0 {
		nombreDeBits++
		maxRang >>= 1
	}

	for bit := 0; bit < nombreDeBits; bit++ {
		taille := len(p.a)

		for i := 0; i < taille; i++ {
			if (p.a[0]>>bit)&1 == 1 {
				p.ra()
			} else {
				p.pb()
			}
		}

		for len(p.b) > 0 {
			p.pa()
		}
	}
}

func rangs(nombres []int) []int {
	resultat := make([]int, len(nombres))

	for i, nombre := range nombres {
		plusPetits := 0
		for _, autre := range nombres {
			if autre < nombre {
				plusPetits++
			}
		}
		resultat[i] = plusPetits
	}

	return resultat
}