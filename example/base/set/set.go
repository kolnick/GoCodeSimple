package set

type Set struct {
	set map[interface{}]bool
}

func (this *Set) init() {
	this.set = make(map[interface{}]bool)
}

func (this *Set) Add(value interface{}) {
	this.set[value] = true
}
func (this *Set) Remove(value interface{}) {
	delete(this.set, value)
}
