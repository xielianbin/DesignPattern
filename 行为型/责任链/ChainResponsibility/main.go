package main

import "fmt"
import "chainResponsibility/internal"
func main() {
	// 创建各个部门
	reception := &internal.Reception{}
	doctor:= &internal.Doctor{}
	medical := &internal.Medical{}
	cashier := &internal.Cashier{}

	// 设置责任链
	reception.SetNext(doctor)
	doctor.SetNext(medical)
	medical.SetNext(cashier)

	// 创建患者
	patient := &internal.Patient{Name: "张三"}

	// 执行责任链
	fmt.Println("开始处理患者流程...")
	reception.Execute(patient)
	fmt.Println("患者流程处理完毕")

}
