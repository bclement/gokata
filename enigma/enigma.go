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
const LEFT = 2
const MID = 1
const RIGHT = 0

/*
Holds lookup for rotor and current position.
Lookup string has step value as index 26
*/
type Rotor struct{
    arr string
    pos byte
}

/*
Holds rotors and reflector.
Rotors go in execution order (right, mid, left)
*/
type Machine struct{
    rotors []Rotor
    reflector string
    dstep bool
}

/*
Create a new machine.
Rotors must be 27 in length. Last item is step value.
Cypher path is right->middle->left->reflect and back
*/
func New(reflector, left, middle, right string) (*Machine){
    rs := []Rotor{Rotor{right,0},Rotor{middle,0},Rotor{left,0}}
    return &Machine{rs, reflector, false}
}

/*
Initialize rotor positions.
Init string must be 3 in length.
*/
func (m *Machine) Set(init string) error{
    size := len(init)
    if size != len(m.rotors){
        return fmt.Errorf("Invalid init string: %s", init)
    }
    // TODO check if ascii A-Z
    upper := strings.ToUpper(init)
    for i := range m.rotors{
        // we store rotors in reverse
        m.rotors[i].pos = upper[size-1-i] - 'A'
    }
    m.dstep = false
    return nil
}

/*
return rotor settings
*/
func (m *Machine) Get() string{
    var rval bytes.Buffer
    for i := len(m.rotors)-1; i >= 0; i-=1{
        r := m.rotors[i]
        rval.WriteByte(r.pos + 'A')
    }
    return rval.String()
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
    for _, x := range msg{
        m.step()
        rval.WriteByte(m.codec(byte(x)))
    }
    return rval.String()
}

/*
step rotors
*/
func (m *Machine) step(){
    // always step right
    m.rotors[RIGHT].step()
    // check if right will step mid
    midStepped := true //optimistic
    if m.rotors[RIGHT].willStep(){
        m.rotors[MID].step()
        // mid steps twice in a row
        m.dstep = true
    } else if m.dstep{
        m.rotors[MID].step()
        m.dstep = false
    } else {
        midStepped = false
    }
    // check if mid stepped and will step left
    if midStepped && m.rotors[MID].willStep(){
        m.rotors[LEFT].step()
    }
}

/*
rotate the rotor by 1/26th
*/
func (r *Rotor) step(){
    r.pos = (r.pos + 1) % 26
}

/*
return true if this rotor should step the following rotor when rotating
*/
func (r *Rotor) willStep() bool{
    return r.pos + 'A' == r.arr[STEP]
}

/*
add offset while keeping in valid range
offset can be negative
*/
func adjust(b byte, offset byte) byte{
    return (b + offset + 26) % 26
}

/*
b is index from 0-25
encode towards reflector
takes rotation position into account
*/
func (r *Rotor) rightToLeft(b byte) byte{
    i := adjust(b, r.pos)
    return adjust(r.arr[i] - 'A', -r.pos)
}

/*
b is index from 0-25
encode away from reflector
takes rotation position into acount
*/
func (r *Rotor) leftToRight(b byte) byte{
    i := adjust(b, r.pos)
    b = byte(strings.IndexRune(r.arr, rune(i + 'A')))
    return adjust(b, -r.pos)
}

/*
b is ascii A-Z
internal function for encoding and decoding
*/
func (m *Machine) codec(b byte) byte{
    b -= 'A'
    for _, r := range m.rotors{
        b = r.rightToLeft(b)
    }
    b = m.reflector[b] - 'A'
    for i := len(m.rotors)-1; i >= 0; i-=1{
        r:= m.rotors[i]
        b = r.leftToRight(b)
    }
    return b + 'A'
}
