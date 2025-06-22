package internal

type Patient struct {
	Name string
	RegistrationDone bool // 是否挂号
	DoctorCheckUpDone bool // 是否医生检查
	MedicineDone bool // 是否取药
	PaymentDone bool // 是否支付
}