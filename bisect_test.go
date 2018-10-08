package bisect

import (
	"github.com/theodesp/bisect"
	"testing"
)

func TestIntSlice_InsortRight(t *testing.T) {
	ints := createTestIntSlice(10)

	originalLenght := ints.Len()
	ints.InsortRight(2)

	if ints.Len() == originalLenght {
		t.Errorf("expecting int slice items count increase after InsortRight")
	}

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 2}

	for i := range expected {
		if ints[i] != expected[i] {
			t.Errorf("expected value %d in position %d. got %d instead", expected[i], i, ints[i])
		}
	}

}

func TestIntSlice_InsortLeft(t *testing.T) {
	ints := createTestIntSlice(10)

	originalLenght := ints.Len()
	ints.InsortLeft(2)

	if ints.Len() == originalLenght {
		t.Errorf("expecting int slice items count increase after InsortLeft")
	}

	expected := []int{0, 1, 2, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := range expected {
		if ints[i] != expected[i] {
			t.Errorf("expected value %d in position %d. got %d instead", expected[i], i, ints[i])
		}
	}

}

func TestEmptyIntSlice_InsortLeft(t *testing.T) {
	ints := createTestIntSlice(0)

	originalLenght := ints.Len()
	ints.InsortLeft(2)

	if ints.Len() == originalLenght {
		t.Errorf("expecting int slice items count increase after InsortLeft")
	}
}

func TestEmptyIntSlice_Insorts(t *testing.T) {
	ints := createTestIntSlice(0)

	originalLenght := ints.Len()
	ints.InsortLeft(2)
	ints.InsortLeft(3)
	ints.InsortLeft(5)
	ints.InsortLeft(2)
	ints.InsortRight(99)
	ints.InsortRight(-33)

	if ints.Len() == originalLenght {
		t.Errorf("expecting int slice items count increase after insort operation")
	}

	expected := []int{99, 2, 2, 3, 5, -33}

	for i := range expected {
		if ints[i] != expected[i] {
			t.Errorf("expected value %d in position %d. got %d instead", expected[i], i, ints[i])
		}
	}
}

func createTestIntSlice(size int) bisect.IntSlice {
	var ints bisect.IntSlice
	for i := 0; i < size; i++ {
		ints = append(ints, i)
	}
	return ints
}
