package generals

import (
	"testing"
)

func Test_IntIntNewOptions(t *testing.T) {
	m1 := MapOptions[int, int]{1: 2, 2: 4, 4: 8, 8: 16}
	rs := NewOptions(m1)
	//fmt.Println(rs.Keys())
	key := rs.Option(1)
	t.Log(key)
	t.Log(rs.Options())
}

func Test_IntStringNewOptions(t *testing.T) {
	m1 := MapOptions[int, string]{1: "2", 2: "4", 4: "8", 8: "16"}
	rs := NewOptions(m1)
	t.Log(rs.Option(1))
	t.Log(rs.Options())
}

func TestOptions_Keys(t *testing.T) {
	m1 := MapOptions[int, string]{1: "2", 2: "4", 4: "8", 8: "16"}
	rs := NewOptions(m1)
	t.Log(rs.keys)
}

func Test_IntFloat(t *testing.T) {
	m1 := MapOptions[int, float64]{1: 0.011, 2: 2.1, 4: 12, 8: 01.3}
	rs := NewOptions(m1)
	t.Log(rs.keys)
	t.Log(rs.Options())
}
