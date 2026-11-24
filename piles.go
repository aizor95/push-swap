package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------- Les piles et les opérations (séance 2) ----------

type Piles struct {
	a          []int
	b          []int
	operations []string
}

func (p *Piles) sa() {
	if len(p.a) < 2 {
		return
	}
	p.a[0], p.a[1] = p.a[1], p.a[0]
	p.operations = append(p.operations, "sa")
}

func (p *Piles) sb() {
	if len(p.b) < 2 {
		return
	}
	p.b[0], p.b[1] = p.b[1], p.b[0]
	p.operations = append(p.operations, "sb")
}

func (p *Piles) pa() {
	if len(p.b) == 0 {
		return
	}
	p.a = append([]int{p.b[0]}, p.a...)
	p.b = p.b[1:]
	p.operations = append(p.operations, "pa")
}

func (p *Piles) pb() {
	if len(p.a) == 0 {
		return
	}
	p.b = append([]int{p.a[0]}, p.b...)
	p.a = p.a[1:]
	p.operations = append(p.operations, "pb")
}

func (p *Piles) ra() {
	if len(p.a) < 2 {
		return
	}
	p.a = append(p.a[1:], p.a[0])
	p.operations = append(p.operations, "ra")
}

func (p *Piles) rb() {
	if len(p.b) < 2 {
		return
	}
	p.b = append(p.b[1:], p.b[0])
	p.operations = append(p.operations, "rb")
}

func (p *Piles) rra() {
	if len(p.a) < 2 {
		return
	}
	dernier := p.a[len(p.a)-1]
	p.a = append([]int{dernier}, p.a[:len(p.a)-1]...)
	p.operations = append(p.operations, "rra")
}

func (p *Piles) rrb() {
	if len(p.b) < 2 {
		return
	}
	dernier := p.b[len(p.b)-1]
	p.b = append([]int{dernier}, p.b[:len(p.b)-1]...)
	p.operations = append(p.operations, "rrb")
}

func (p *Piles) ss() {
	effectue := false

	if len(p.a) >= 2 {
		p.a[0], p.a[1] = p.a[1], p.a[0]
		effectue = true
	}

	if len(p.b) >= 2 {
		p.b[0], p.b[1] = p.b[1], p.b[0]
		effectue = true
	}

	if effectue {
		p.operations = append(p.operations, "ss")
	}
}

func (p *Piles) rr() {
	effectue := false

	if len(p.a) >= 2 {
		p.a = append(p.a[1:], p.a[0])
		effectue = true
	}

	if len(p.b) >= 2 {
		p.b = append(p.b[1:], p.b[0])
		effectue = true
	}

	if effectue {
		p.operations = append(p.operations, "rr")
	}
}

func (p *Piles) rrr() {
	effectue := false

	if len(p.a) >= 2 {
		dernier := p.a[len(p.a)-1]
		p.a = append([]int{dernier}, p.a[:len(p.a)-1]...)
		effectue = true
	}

	if len(p.b) >= 2 {
		dernier := p.b[len(p.b)-1]
		p.b = append([]int{dernier}, p.b[:len(p.b)-1]...)
		effectue = true
	}

	if effectue {
		p.operations = append(p.operations, "rrr")
	}
}

// aEstTriee renvoie true si chaque élément de a est inférieur ou égal au suivant.
func (p *Piles) aEstTriee() bool {
	for i := 0; i < len(p.a)-1; i++ {
		if p.a[i] > p.a[i+1] {
			return false
		}
	}
	return true
}

// ---------- Lire et vérifier les arguments ----------

// lireNombres transforme les arguments en nombres,
// et refuse un argument qui n'est pas un nombre entier.
func lireNombres(args []string) ([]int, error) {
	var nombres []int
	for _, texte := range args {
		valeurs := strings.Fields(texte)
		for _, valeur := range valeurs {
			nombre, err := strconv.Atoi(valeur)
			if err != nil {
				return nil, fmt.Errorf("%q n'est pas un nombre entier", valeur)
			}
			nombres = append(nombres, nombre)
		}
	}
	return nombres, nil
}

// verifierDoublons renvoie une erreur si un nombre apparaît deux fois.
func verifierDoublons(nombres []int) error {
	vus := map[int]bool{}

	for _, nombre := range nombres {
		if vus[nombre] {
			return fmt.Errorf("le nombre %d apparaît plusieurs fois", nombre)
		}
		vus[nombre] = true
	}

	return nil
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

type etatOptimal struct {
	valeurs [7]uint8 // A suivie de B
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
		_ = p.executerInstruction(nomsOps[chemin[i]])
	}
}

// ---------- Les rangs ----------

// rangs remplace chaque nombre par sa position dans l'ordre croissant :
// 0 pour le plus petit, n-1 pour le plus grand.
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

// ---------- Le programme ----------

// trier choisit la méthode selon le nombre d'éléments.
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

func executer(args []string) error {
	if len(args) == 0 {
		return nil
	}

	nombres, err := lireNombres(args)
	if err != nil {
		return err
	}

	err = verifierDoublons(nombres)
	if err != nil {
		return err
	}

	p := Piles{a: rangs(nombres)}
	p.trier()

	if !p.aEstTriee() || len(p.b) != 0 {
		return errors.New("la pile n'a pas été triée")
	}

	for _, operation := range p.operations {
		fmt.Println(operation)
	}
	return nil
}

// executerInstruction exécute une instruction reçue en argument.
func (p *Piles) executerInstruction(instruction string) error {
	switch instruction {
	case "sa":
		p.sa()
	case "sb":
		p.sb()
	case "ss":
		p.ss()
	case "pa":
		p.pa()
	case "pb":
		p.pb()
	case "ra":
		p.ra()
	case "rb":
		p.rb()
	case "rr":
		p.rr()
	case "rra":
		p.rra()
	case "rrb":
		p.rrb()
	case "rrr":
		p.rrr()
	default:
		return fmt.Errorf("instruction inconnue : %s", instruction)
	}
	return nil
}

func lireInstructions(p *Piles) error {
	for {
		var instruction string
		n, err := fmt.Fscanln(os.Stdin, &instruction)
		if err != nil {
			if err.Error() == "EOF" {
				return nil
			}
			return err
		}
		if n != 1 {
			return fmt.Errorf("instruction invalide")
		}
		err = p.executerInstruction(instruction)
		if err != nil {
			return err
		}
	}
}

func verifier(args []string) (bool, error) {
	nombres, err := lireNombres(args)
	if err != nil {
		return false, err
	}
	err = verifierDoublons(nombres)
	if err != nil {
		return false, err
	}
	p := Piles{
		a: nombres,
	}
	err = lireInstructions(&p)
	if err != nil {
		return false, err
	}
	if p.aEstTriee() && len(p.b) == 0 {
		return true, nil
	}
	return false, nil
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	cheminExe, err := os.Executable()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	nomProgramme := strings.ToLower(cheminExe)
	if strings.HasSuffix(nomProgramme, "checker.exe") || strings.HasSuffix(nomProgramme, "checker") {
		ok, err := verifier(os.Args[1:])

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}

		if ok {
			fmt.Println("OK")
		} else {
			fmt.Println("KO")
		}
		return
	}
	err = executer(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
}
