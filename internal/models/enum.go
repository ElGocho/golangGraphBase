package models


type Priority byte 

const (
	Priority1 Priority = iota 
	Priority2
	Priority3
)


func (p *Priority) Valid() {
	if p == nil {
		p.Base()
		return
	}

	switch(*p){
		case Priority1: 
		case Priority2: 
		case Priority3: break
		default: p.Base()
	}
}

func (p *Priority) Base() {
	base := Priority1

	p = &base
}
