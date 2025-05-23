package crypto

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestClassDiscriminant(t *testing.T) {
	t12_11_3 := NewClassGroup(big.NewInt(12), big.NewInt(11), big.NewInt(3))
	require.Equal(t, "-23", t12_11_3.Discriminant().String(), "they should be equal")

	t93_109_32 := NewClassGroup(big.NewInt(93), big.NewInt(109), big.NewInt(32))
	require.Equal(t, "-23", t93_109_32.Discriminant().String(), "they should be equal")

	D := big.NewInt(-103)
	e_id := newClassGroup(bigOne, bigOne, D)
	require.Equal(t, NewClassGroup(big.NewInt(1), big.NewInt(1), big.NewInt(26)), e_id, "they should be equal")
	require.Equal(t, e_id.Discriminant(), D, "they should be equal")

	e := newClassGroup(big.NewInt(2), big.NewInt(1), D)
	require.Equal(t, NewClassGroup(big.NewInt(2), big.NewInt(1), big.NewInt(13)), e, "they should be equal")
	require.Equal(t, e.Discriminant(), D, "they should be equal")
}

func TestNormalized(t *testing.T) {
	f := NewClassGroup(big.NewInt(195751), big.NewInt(1212121), big.NewInt(1876411))
	require.Equal(t, NewClassGroup(big.NewInt(195751), big.NewInt(37615), big.NewInt(1807)), f.Normalized(), "they should be equal")
}

func TestReduced(t *testing.T) {
	f := NewClassGroup(big.NewInt(195751), big.NewInt(1212121), big.NewInt(1876411))
	require.Equal(t, NewClassGroup(big.NewInt(1), big.NewInt(1), big.NewInt(1)), f.Reduced(), "they should be equal")
}

func check(a, b, c *big.Int, t *testing.T) {
	r, s, _ := solveMod(a, b, c)
	b.Mod(b, c)

	for k := 0; k < 50; k++ {
		//a_coefficient = r + s * k
		a_coefficient := new(big.Int).Add(r, new(big.Int).Mul(s, big.NewInt(int64(k))))
		aac := a_coefficient.Mul(a_coefficient, a)
		aac.Mod(aac, c)
		require.Equal(t, aac, b, fmt.Sprintf("diff when k = %d", k))
	}
}

func TestSolveMod(t *testing.T) {
	check(big.NewInt(3), big.NewInt(4), big.NewInt(5), t)
	check(big.NewInt(6), big.NewInt(8), big.NewInt(10), t)
	check(big.NewInt(12), big.NewInt(30), big.NewInt(7), t)
	check(big.NewInt(6), big.NewInt(15), big.NewInt(411), t)
	check(big.NewInt(192), big.NewInt(193), big.NewInt(863), t)
	check(big.NewInt(-565721958), big.NewInt(740), big.NewInt(4486780496), t)
	check(big.NewInt(565721958), big.NewInt(740), big.NewInt(4486780496), t)
	check(big.NewInt(-565721958), big.NewInt(-740), big.NewInt(4486780496), t)
	check(big.NewInt(565721958), big.NewInt(-740), big.NewInt(4486780496), t)
}

func TestMultiplication1(t *testing.T) {
	t12_11_3 := NewClassGroup(big.NewInt(12), big.NewInt(11), big.NewInt(3))
	t93_109_32 := NewClassGroup(big.NewInt(93), big.NewInt(109), big.NewInt(32))

	a := t12_11_3.Multiply(t93_109_32)
	require.Equal(t, a, NewClassGroup(big.NewInt(1), big.NewInt(1), big.NewInt(6)), "they should be equal")
}

func TestMultiplication2(t *testing.T) {
	t12_11_3 := NewClassGroup(big.NewInt(12), big.NewInt(11), big.NewInt(3))
	t93_109_32 := NewClassGroup(big.NewInt(93), big.NewInt(109), big.NewInt(32))

	x := CloneClassGroup(t12_11_3)
	y := t12_11_3.Multiply(x)
	require.Equal(t, y, NewClassGroup(big.NewInt(2), big.NewInt(1), big.NewInt(3)), "they should be equal")

	x = CloneClassGroup(t93_109_32)
	y = t93_109_32.Multiply(x)
	require.Equal(t, y, NewClassGroup(big.NewInt(2), big.NewInt(-1), big.NewInt(3)), "they should be equal")
}

func TestMultiplication3(t *testing.T) {
	t12_11_3 := NewClassGroup(big.NewInt(12), big.NewInt(11), big.NewInt(3))
	t93_109_32 := NewClassGroup(big.NewInt(93), big.NewInt(109), big.NewInt(32))

	a := t12_11_3.Multiply(t93_109_32)
	require.Equal(t, a, NewClassGroup(big.NewInt(1), big.NewInt(1), big.NewInt(6)), "they should be equal")
}

func TestMultiplication4(t *testing.T) {
	x := NewClassGroup(big.NewInt(-565721958), big.NewInt(-740), big.NewInt(4486780496))
	y := NewClassGroup(big.NewInt(565721958), big.NewInt(740), big.NewInt(4486780496))

	a := x.Multiply(y)
	fmt.Println(a.a.String(), a.b.String(), a.c.String())
	require.Equal(t, a.a, big.NewInt(-1))
	require.Zero(t, a.b.Cmp(bigZero))
	require.Equal(t, a.c, big.NewInt(2538270247313468068))
}

func TestSquare1(t *testing.T) {
	x := NewClassGroup(big.NewInt(12), big.NewInt(11), big.NewInt(3))
	y := x.Square()
	require.Equal(t, y, NewClassGroup(big.NewInt(2), big.NewInt(1), big.NewInt(3)), "they should be equal")
}

func TestSquare2(t *testing.T) {
	x := NewClassGroup(big.NewInt(93), big.NewInt(109), big.NewInt(32))
	y := x.Square()
	require.Equal(t, y, NewClassGroup(big.NewInt(2), big.NewInt(-1), big.NewInt(3)), "they should be equal")
}
