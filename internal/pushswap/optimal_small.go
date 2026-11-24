package pushswap

type etatOptimal struct {
	valeurs [7]uint8
	tailleA uint8
}

type noeudOptimal struct {
	etat   etatOptimal
	parent int
	op     uint8
}

var nomsOps = [...]string{
	"sa", "sb", "ss", "pa", "pb",
	"ra", "rb", "rr", "rra", "rrb", "rrr",
}

func appliquerOp(etat etatOptimal, n, op int) etatOptimal {
	tailleA := int(etat.tailleA)

	gauche := func(debut, fin int) {
		if fin-debut < 2 {
			return
		}
		premier := etat.valeurs[debut]
		copy(etat.valeurs[debut:fin-1], etat.valeurs[debut+1:fin])
		etat.valeurs[fin-1] = premier
	}

	droite := func(debut, fin int) {
		if fin-debut < 2 {
			return
		}
		dernier := etat.valeurs[fin-1]
		copy(etat.valeurs[debut+1:fin], etat.valeurs[debut:fin-1])
		etat.valeurs[debut] = dernier
	}

	switch op {
	case 0: // sa
		if tailleA >= 2 {
			etat.valeurs[0], etat.valeurs[1] = etat.valeurs[1], etat.valeurs[0]
		}
	case 1: // sb
		if n-tailleA >= 2 {
			i := tailleA
			etat.valeurs[i], etat.valeurs[i+1] = etat.valeurs[i+1], etat.valeurs[i]
		}
	case 2: // ss
		if tailleA >= 2 {
			etat.valeurs[0], etat.valeurs[1] = etat.valeurs[1], etat.valeurs[0]
		}
		if n-tailleA >= 2 {
			i := tailleA
			etat.valeurs[i], etat.valeurs[i+1] = etat.valeurs[i+1], etat.valeurs[i]
		}
	case 3: // pa
		if tailleA < n {
			topB := etat.valeurs[tailleA]
			copy(etat.valeurs[1:tailleA+1], etat.valeurs[:tailleA])
			etat.valeurs[0] = topB
			tailleA++
		}
	case 4: // pb
		if tailleA > 0 {
			gauche(0, tailleA)
			tailleA--
		}
	case 5: // ra
		gauche(0, tailleA)
	case 6: // rb
		gauche(tailleA, n)
	case 7: // rr
		gauche(0, tailleA)
		gauche(tailleA, n)
	case 8: // rra
		droite(0, tailleA)
	case 9: // rrb
		droite(tailleA, n)
	case 10: // rrr
		droite(0, tailleA)
		droite(tailleA, n)
	}

	etat.tailleA = uint8(tailleA)
	return etat
}

func etatTrie(etat etatOptimal, n int) bool {
	if int(etat.tailleA) != n {
		return false
	}
	for i := 0; i < n; i++ {
		if etat.valeurs[i] != uint8(i) {
			return false
		}
	}
	return true
}

func (p *Piles) trierOptimalPetit() {
	n := len(p.a)

	var depart etatOptimal
	depart.tailleA = uint8(n)
	for i, valeur := range p.a {
		depart.valeurs[i] = uint8(valeur)
	}

	noeuds := []noeudOptimal{{etat: depart, parent: -1}}
	vus := map[etatOptimal]int{depart: 0}
	indexTrie := -1

	for tete := 0; tete < len(noeuds); tete++ {
		if etatTrie(noeuds[tete].etat, n) {
			indexTrie = tete
			break
		}

		for op := range nomsOps {
			suivant := appliquerOp(noeuds[tete].etat, n, op)
			if _, existe := vus[suivant]; existe {
				continue
			}

			vus[suivant] = len(noeuds)
			noeuds = append(noeuds, noeudOptimal{
				etat:   suivant,
				parent: tete,
				op:     uint8(op),
			})
		}
	}

	if indexTrie == -1 {
		return
	}

	var chemin []int
	for index := indexTrie; noeuds[index].parent != -1; index = noeuds[index].parent {
		chemin = append(chemin, int(noeuds[index].op))
	}

	for i := len(chemin) - 1; i >= 0; i-- {
		_ = p.ApplyInstruction(nomsOps[chemin[i]])
	}
}