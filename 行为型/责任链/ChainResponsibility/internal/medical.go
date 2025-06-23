package internal
import "fmt"
type Medical struct {
	next Department	
}
func (m *Medical) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("开药已完成，转到下一个部门")
		m.next.Execute(p)
		return
	}
	fmt.Println("正在进行开药...")
	p.MedicineDone = true
	fmt.Println("开药完成，转到下一个部门")
	if m.next != nil {
		m.next.Execute(p)
	} else {
		fmt.Println("没有下一个部门，流程结束")
	}
}
func (m *Medical) SetNext(next Department) {
	m.next = next
}