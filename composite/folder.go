package composite

import "fmt"

type folder struct {
	components []iComponent
	name       string
}

func (f *folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *folder) add(c iComponent) {
	f.components = append(f.components, c)
}
