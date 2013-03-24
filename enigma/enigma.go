package enigma

import "strings"
import "fmt"
import "bytes"

// rotors, index 26 is step position
const I =   "EKMFLGDQVZNTOWYHXUSPAIBRCJR"
const II =  "AJDKSIRUXBLHWTMCQGZNPYFVOEF"
const III = "BDFHJLCPRTXVZNYEIWGAKMUSQOW"
const IV =  "ESOVPZJAYQUIRHXLNFTGKDCMWBK"
const V =   "VZBRGITYUPSDNHLXAWMJQOFECKA"

// reflectors
const A =   "EJMZALYXVBWFCRQUONTSPIKHGD"
const B =   "YRUHQSLDPXNGOKMIEBFZCWVJAT"
const C =   "FVPJIAOYEDRZXWGCTKUQSBNMHL"

const STEP = 26

/*
Holds lookup for rotor and current position.
Lookup string has step value as index 26
*/
type Rotor struct{
    arr string
    pos int
}

/*
Holds rotors and reflector.
Rotors go in execution order (right, mid, left)
*/
type Machine struct{
    rotors []Rotor
    reflector string
}

/*
Create a new machine.
Rotors must be 27 in length. Last item is step value.
Cypher path is right->middle->left->reflect and back
*/
func New(reflector, left, middle, right string) (*Machine){
    rs := []Rotor{Rotor{right,0},Rotor{middle,0},Rotor{left,0}}
    return &Machine{rs, reflector}
}

/*
Initialize rotor positions.
Init string must be 3 in length.
*/
func (m *Machine) Set(init string) error{
    if len(init) != 3{
        return fmt.Errorf("Invalid init string: %s", init)
    }
    // TODO check if ascii A-Z
    upper := strings.ToUpper(init)
    // we store rotors in reverse
    for i := 2; i >= 0; i -= 1{
        m.rotors[i].pos = int(upper[i] - 'A')
    }
    return nil
}

/*
Encode string
*/
func (m *Machine) Enc(msg string) (string, error){
    upper := strings.ToUpper(msg)
    upper = strings.Replace(upper, " ", "X", -1)
    // TODO format output
    return m.codecStr(upper), nil
}

/*
Decode string
*/
func (m *Machine) Dec(msg string) (string, error){
    upper := strings.ToUpper(msg)
    // TODO unformat input
    return m.codecStr(upper), nil
}

/*
Internal function for encoding and decoding
*/
func (m *Machine) codecStr(msg string) string{
    var rval bytes.Buffer
    for x := range msg{
        rval.WriteByte(m.codec(msg[x]))
        // TODO step
    }
    return rval.String()
}

/*
internal function for encoding and decoding
*/
func (m *Machine) codec(b byte) byte{
    for i := range m.rotors{
        r := m.rotors[i]
        b = byte(strings.IndexRune(r.arr, rune(b))+'A')
    }
    b = m.reflector[b-'A']
    for i := len(m.rotors)-1; i >= 0; i-=1{
        r:= m.rotors[i]
        b = r.arr[b-'A']
    }
    return b
}
