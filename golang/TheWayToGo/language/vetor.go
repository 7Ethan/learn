package main

type Element interface {
}

type Vetor struct {
	a []Element
}

func (p *Vetor) At(i int) Element {
	return p.a[i]
}

func (p *Vetor) Set(i int, data Element) error {
	p.a[i] = data
	return nil
}
