package builder

import "testing"

func TestBuilder(t *testing.T){

    director := Director{}
    L3PortBuilder := PointToPointBuilder{}
    L2PortBuilder := PointToPointBuilder{}

    director.SetBuilder(&L3PortBuilder)
    director.L3PortContruct("Gi0/0/1","192.168.1.1",true,"L3 PORT",Ethernet)

    director.SetBuilder(&L2PortBuilder)
    director.L2PortContruct("Gi0/0/2","L2 PORT",Ethernet)



    l3Port := L3PortBuilder.Build()
    if l3Port.Port != "Gi0/0/1" {
        t.Errorf("Want Gi0/0/1, Got %v\n",l3Port.Port)
    }

    l2Port := L2PortBuilder.Build()
    if l2Port.Port != "Gi0/0/2" {
        t.Errorf("Want Gi0/0/2, Got %v\n",l2Port.Port)
    }




}
