package abstractFactory

type iShirt interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}
