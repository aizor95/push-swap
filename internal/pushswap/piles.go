package pushswap

type Piles struct {
	a          []int
	b          []int
	operations []string
}

func NewCheckerPiles(numbers []int) *Piles {
	copie := append([]int(nil), numbers...)
	return &Piles{a: copie}
}

func (p *Piles) aEstTriee() bool {
	for i := 0; i < len(p.a)-1; i++ {
		if p.a[i] > p.a[i+1] {
			return false
		}
	}
	return true
}

func (p *Piles) IsSorted() bool {
	return p.aEstTriee()
}

func (p *Piles) BEmpty() bool {
	return len(p.b) == 0
}