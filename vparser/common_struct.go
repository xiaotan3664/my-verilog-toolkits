package vparser

const (
	NET_TYPE_WIRE = "wire"
	NET_TYPE_REG  = "reg"
)

type Param struct {
	Name  string
	Value string
}

type Signal struct {
	Name      string
	NetType   string
	Width     string //paramter or digits
	IndexFrom string //paramter or digits
	IndexTo   string //paramter or digits
	Drivers   []*Signal
	Connects  []*Signal
	Instances []*Module
}

type Port struct {
	Signal
	Direction string //input, output, inout
}

type Module struct {
	Name   string
	Params []*Param
	Ports  []*Port
}
