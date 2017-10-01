package model

type Robot struct {
	Name string
	Type string
}

func (r *Robot) SayHalo(me string) {
	r.Name = me
}
