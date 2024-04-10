
type IPAddress struct {
    Address     string
    Primary     bool
}


type Director struct {
    b Builder
}

func (d *Director) SetBuilder(b Builder) {
    d.b = b
}

func (d *Director) L3PortContruct(port string,ip string,is_primary bool, description string, speed Speed){
    d.b.SetPort(port).SetIP(port,is_primary).SetDescription(description).SetSpeed(speed)
}

func (d *Director) L2PortContruct(port string,description string, speed Speed){
    d.b.SetPort(port).SetDescription(description).SetSpeed(speed)
}

type Builder interface {
    SetPort(string)             Builder
    SetIP(string,bool)            Builder
    SetDescription(string)      Builder
    SetSpeed(Speed)             Builder
}

type Speed int

const (
    none Speed = iota
    Ethernet
    GigabitEthernet
    TenGigabitEthernet
)

type Componet struct {
    Port            string
    IP              IPAddress
    Description     string
    Speed           Speed
}

type PointToPointBuilder struct {
    component       Componet
}

func (p *PointToPointBuilder) SetPort(s string) Builder {
    p.component.Port = s
    return p
}

func (p *PointToPointBuilder) SetIP(ip string,is_primary bool) Builder {
    p.component.IP.Address = ip
    p.component.IP.Primary = is_primary
    return p
}

func (p *PointToPointBuilder) SetDescription(s string) Builder {
    p.component.Description = s
    return p
}

func (p *PointToPointBuilder) SetSpeed(s Speed) Builder {
    p.component.Speed = s
    return p
}

func (p *PointToPointBuilder) Build() Componet {
    return p.component
}


