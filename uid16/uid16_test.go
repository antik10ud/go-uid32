package uid16

import (
	"testing"
	"fmt"
	"time"
	"math/rand"
)

func TestGenerate(t *testing.T) {
	g := NewFactory()
	v := g.New()
	v2, _ := g.FromString(v.String())
	chkEqual(t, v.String(), v2.String())
}

func TestVariableEncoding(t *testing.T) {
	g := NewFactory()
	v1, err := g.FromString("0008hpxy8cgub03qjull4frtkf")
	chk(t, err)
	v2, err := g.FromString("8hpxy8cgub03qjull4frtkf")
	chk(t, err)
	chkEqual(t, v1.String(), v2.String())
}

func BenchmarkGenerate(b *testing.B) {
	g := NewFactory()
	for i := 0; i < 1000000; i++ {
		g.New()
	}
}

func TestExampleGenerate(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	g := NewFactory()
	for i := 0; i < 10; i++ {
		id := g.New()
		fmt.Println(id.String())
	}
	//// Output: hello
}

func TestShorten(t *testing.T) {
	g := NewFactory()
	v1, err := g.FromString("0008hpxy8cgub03qjull4frtkf")
	chk(t, err)
	v2 := v1.Shorten()
	chkEqual(t, "8hpxy8cgub03qjull4frtkf", v2)
}

func TestMax(t *testing.T) {
	v1 := maxId16()
	chkEqual(t, "yyyyyyyyyyyyyyyyyyyyyyyyyv", v1.String())
}

func TestMin(t *testing.T) {
	v1 := minId16()
	chkEqual(t, "00000000000000000000000000", v1.String())


}

func TestVoid(t *testing.T) {
	g := NewFactory()
	v1, err := g.FromString("")
	chk(t, err)
	s1 := v1.String()
	chkEqual(t, "00000000000000000000000000", s1)
	v2 := v1.Shorten()
	chkEqual(t, "0", v2)
}
func TestVoid2(t *testing.T) {
	g := NewFactory()
	v1, err := g.FromString("0")
	chk(t, err)
	s1 := v1.String()
	chkEqual(t, "00000000000000000000000000", s1)
	v2 := v1.Shorten()
	chkEqual(t, "0", v2)
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
	g := Factory{
		lastSeq:  0xff,
		timeFunc: timeFunc,
		randFunc: randFunc,
	}

	v1 := g.New()
	chkEqual(t, "00000000000040000000000000", v1.String())
	v2 := g.New()
	chkEqual(t, "00000000000040b00000000000", v2.String())

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
