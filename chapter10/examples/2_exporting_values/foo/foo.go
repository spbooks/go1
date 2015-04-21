package foo

import "fmt"

type Bar struct {
	Visible bool
	hidden  bool
}

func NewBar() Bar {
	return Bar{
		Visible: true,
		hidden:  true,
	}
}

func (bar Bar) String() string {
	return bar.status()
}
func (bar Bar) status() string {
	return fmt.Sprintf("Visible is %T and hidden is %T", bar.Visible, bar.hidden)
}
