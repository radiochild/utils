package text

var exists = struct{}{}

type NameSet map[string]struct{}

func (ns NameSet) Add(value string) {
	ns[value] = exists
}

func (ns NameSet) Remove(value string) {
	delete(ns, value)
}

func (ns NameSet) Contains(value string) bool {
	_, ok := ns[value]
	return ok
}

func (ns NameSet) Members() []string {
	members := make([]string, len(ns))
	i := 0
	for k := range ns {
		members[i] = k
		i++
	}
	return members
}

func FromStrings(strs []string) NameSet {
	ns := NameSet{}
	for _, str := range strs {
		ns[str] = exists
	}
	return ns
}
