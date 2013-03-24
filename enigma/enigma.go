package enigma

var I = []byte{1, 2, 3}

type Rotor struct{
    arr []byte
    step int
    pos int
}

type Machine struct{
    rotors []Rotor
    reflector []byte
}
