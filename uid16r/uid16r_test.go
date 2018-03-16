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
	v1, err := g.FromString("yyysksp6kunvwyrl0gd0jpaeir")
	chk(t, err)
	v2, err := g.FromString("sksp6kunvwyrl0gd0jpaeir")
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
	v1, err := g.FromString("yyysksp6kunvwyrl0gd0jpaeir")
	chk(t, err)
	v2 := v1.Shorten()
	chkEqual(t, "sksp6kunvwyrl0gd0jpaeir", v2)
}

func TestMax(t *testing.T) {
	v1 := maxId16r()
	chkEqual(t, "yyyyyyyyyyyyyyyyyyyyyyyyyv", v1.String())

}

func TestMin(t *testing.T) {
	v1 := minId16r()
	chkEqual(t, "00000000000000000000000000", v1.String())


}

func TestVoid(t *testing.T) {
	g := NewUId16rGen()
	v1, err := g.FromString("")
	chk(t, err)
	s1 := v1.String()
	chkEqual(t, "yyyyyyyyyyyyyyyyyyyyyyyyyv", s1)
	v2 := v1.Shorten()
	chkEqual(t, "v", v2)
}
func TestVoid2(t *testing.T) {
	g := NewUId16rGen()
	v1, err := g.FromString("y")
	chk(t, err)
	s1 := v1.String()
	chkEqual(t, "yyyyyyyyyyyyyyyyyyyyyyyyyv", s1)
	v2 := v1.Shorten()
	chkEqual(t, "v", v2)
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
	chkEqual(t, "yyyyyyyyyyyywyr00000000000", v1.String())
	v2 := g.New()
	chkEqual(t, "yyyyyyyyyyyywyj00000000000", v2.String())

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
