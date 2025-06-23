package internal
import "fmt"
type Cashier struct {
	next Department
}
func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("付款已完成，转到下一个部门")
		c.next.Execute(p)
		return
	}
	fmt.Println("正在进行付款...")
	p.PaymentDone = true
	fmt.Println("付款完成，转到下一个部门")
	if c.next != nil {
		c.next.Execute(p)
	} else {
		fmt.Println("没有下一个部门，流程结束")
	}
}
func (c *Cashier) SetNext(next Department) {
	c.next = next
}	
