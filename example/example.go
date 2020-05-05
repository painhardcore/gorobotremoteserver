package main

import (
	"fmt"
	gobot "gorobotremoteserver"
)


type hw struct {
	name string
	size int
}


func (m *hw) Name() string {
	return fmt.Sprintf("%s %d",m.name,m.size)
}

func (m *hw) Run(args []string) gobot.Result{
	l := len(args)
	if l < m.size {
		return gobot.Result{Status:"PASS"}
	}
	error := fmt.Sprintf("should be less than %d, current size %d \nContent: %v", m.size,l, args)
	return gobot.Result{Status: "FAIL",Error: error}

}

func main () {
	fmt.Println("starting")
	testSrv := gobot.New()
	m3 := &hw{name: "Args should be less than",size: 3}
	m5 := &hw{name: "Args should be less than",size: 5}
	testSrv.Register(m3)
	testSrv.Register(m5)
	testSrv.Start()
}
