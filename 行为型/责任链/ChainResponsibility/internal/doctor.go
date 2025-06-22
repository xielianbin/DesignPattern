package internal
import "fmt"
type Doctor struct {
	next Department
}
func (d *Doctor) Execute(p *Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("医生检查已完成，转到下一个部门")
		d.next.Execute(p)
		return
	}
	fmt.Println("正在进行医生检查...")
	p.DoctorCheckUpDone = true
	fmt.Println("医生检查完成，转到下一个部门")
	if d.next != nil {
		d.next.Execute(p)
	} else {
		fmt.Println("没有下一个部门，流程结束")
	}
}
func (d *Doctor) SetNext(next Department) {
	d.next = next
}