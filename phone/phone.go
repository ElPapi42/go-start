package phone

import (
	"fmt"
)

// Phone type with it properties
type Phone struct {
	model      string
	make       string
	screenSize int
}

// New phone instance
func New(model string, make string, screenSize int) Phone {
	phone := Phone{model, make, screenSize}
	return phone
}

// Features of the phone
func (p Phone) Features() {
	fmt.Printf("%s %s\n%d Inches", p.make, p.model, p.screenSize)
}
