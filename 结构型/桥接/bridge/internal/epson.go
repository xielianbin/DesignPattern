package internal

import "fmt"

//  爱普生 打印机
type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}
