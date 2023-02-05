package mypackage


type Pc struct{
	Cpu_name string 
	Gpu_name string
	Ram_amount int 
}
func NewPc() *Pc{
	return &Pc{Cpu_name: "Intel i3", Gpu_name: "Nvidia GTX1650", Ram_amount: 8}
}
func ChangeCpu(pc *Pc, newCpu string){
	pc.Cpu_name = newCpu
}
func ChangeGpu(pc *Pc, newGpu string) {
	pc.Cpu_name = newGpu
}
func AddRam(pc *Pc, amount int){
	pc.Ram_amount += amount
}