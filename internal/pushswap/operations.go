package pushswap

import "fmt"

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

func (p *Piles) ApplyInstruction(instruction string) error {
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