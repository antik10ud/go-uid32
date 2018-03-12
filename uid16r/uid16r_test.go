package uid16r

import (
	"testing"
	"fmt"
	"time"
	"math/rand"
)

func TestGenerate(t *testing.T) {
	g := NewUId16rGen()
	v := g.New()
	v2, _ := g.FromString(v.String())
	chkEqual(t, v.String(), v2.String())
}

func TestVariableEncoding(t *testing.T) {
	g := NewUId16rGen()
	v1, err := g.FromString("~~~k~Hb3IAk2WZ1XkT3feG")
	chk(t, err)
	v2, err := g.FromString("k~Hb3IAk2WZ1XkT3feG")
	chk(t, err)
	chkEqual(t, v1.String(), v2.String())
}

func BenchmarkGenerate(b *testing.B) {
	g := NewUId16rGen()
	for i := 0; i < 1000000; i++ {
		g.New()
	}
}

func TestExampleGenerate(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	g := NewUId16rGen()
	for i := 0; i < 10; i++ {
		id := g.New()
		fmt.Println(id.String())
	}
	//// Output: hello
}

func TestShorten(t *testing.T) {
	g := NewUId16rGen()
	v1, err := g.FromString("~~~k~Hb3IAk2WZ1XkT3feG")
	chk(t, err)
	v2 := v1.Shorten()
	chkEqual(t, "k~Hb3IAk2WZ1XkT3feG", v2)
}

func TestMax(t *testing.T) {
	v1 := maxId16r()
	chkEqual(t, "~~~~~~~~~~~~~~~~~~~~~l", v1.String())

}

func TestMin(t *testing.T) {
	v1 := minId16r()
	chkEqual(t, "0000000000000000000000", v1.String())


}

func TestVoid(t *testing.T) {
	g := NewUId16rGen()
	v1, err := g.FromString("")
	chk(t, err)
	s1 := v1.String()
	chkEqual(t, "~~~~~~~~~~~~~~~~~~~~~l", s1)
	v2 := v1.Shorten()
	chkEqual(t, "l", v2)
}
func TestVoid2(t *testing.T) {
	g := NewUId16rGen()
	v1, err := g.FromString("~")
	chk(t, err)
	s1 := v1.String()
	chkEqual(t, "~~~~~~~~~~~~~~~~~~~~~l", s1)
	v2 := v1.Shorten()
	chkEqual(t, "l", v2)
}

func TestIncSeq(t *testing.T) {
	timeFunc := func() uint64 {
		return 1
	}
	randFunc := func(b []byte) {
		for i := 0; i < len(b); i++ {
			b[i] = 0
		}
	}
	g := UId16rGen{
		lastSeq:  0xff,
		timeFunc: timeFunc,
		randFunc: randFunc,
	}

	v1 := g.New()
	chkEqual(t, "~~~~~~~~~~w~0000000000", v1.String())
	v2 := g.New()
	chkEqual(t, "~~~~~~~~~~wz0000000000", v2.String())

}

func chk(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func chkEqual(t *testing.T, v1 string, v2 string) {
	if v1 != v2 {
		t.Error(fmt.Sprintf("expected [%s] but was [%s]", v1, v2))
		t.FailNow()
	}
}
