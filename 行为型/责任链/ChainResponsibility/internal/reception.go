package internal
import "fmt"
type Reception struct {
	next Department	
}
func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone{
		fmt.Println("挂号已完成，转到下一个部门")
		r.next.Execute(p)
		return
	}
	fmt.Println("正在挂号...")
	p.RegistrationDone = true
	fmt.Println("挂号完成，转到下一个部门")
	if r.next != nil {
		r.next.Execute(p)
	} else {
		fmt.Println("没有下一个部门，流程结束")
	}
}
func (r *Reception) SetNext(next Department) {
	r.next = next
}