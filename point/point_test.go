package point

import (
	"fmt"
	"testing"
)

func TestMine(t *testing.T) {
	x := int16(59)
	y := int16(-79)

	{
		sum := sum_digits(x)
		expected := int16(5 + 9)
		if sum != expected {
			fmt.Println(fmt.Sprintf("The sum has failed ! expected : %d result : %d", expected, sum))
			t.Fail()
		}
	}

	{
		sum := sum_digits(y)
		expected := int16(7 + 9)
		if sum != expected {
			fmt.Println(fmt.Sprintf("The sum has failed ! expected : %d result : %d", expected, sum))
			t.Fail()
		}
	}

	if t.Failed() {
		return
	}

	if !New(x, y).IsMine() {
		fmt.Println("There should have a mine at this point : ", x, y)
		t.Fail()
	}
}

func TestNoMine(t *testing.T) {
	x := int16(-113)
	y := int16(-104)

	{
		sum := sum_digits(x)
		expected := int16(1 + 1 + 3)
		if sum != expected {
			fmt.Println(fmt.Sprintf("The sum has failed ! expected : %d result : %d", expected, sum))
			t.Fail()
		}
	}

	{
		sum := sum_digits(y)
		expected := int16(1 + 0 + 4)
		if sum != expected {
			fmt.Println(fmt.Sprintf("The sum has failed ! expected : %d result : %d", expected, sum))
			t.Fail()
		}
	}

	if t.Failed() {
		return
	}

	if New(x, y).IsMine() {
		fmt.Println("There should not have a mine at this point : ", x, y)
		t.FailNow()
	}
}

func TestPath(t *testing.T) {
	//  32767 -> 0111111111111111 -> ‭32767‬
	// -32768 -> 1000000000000000 -> ‭32768

	// Path (32767, -32768) -> 01111111111111111000000000000000 -> ‭2147450880‬
	if point := New(int16(32767), int16(-32768)); point.Path() != uint32(2147450880) {
		fmt.Printf("X : %d -> %b\n", point.x, uint16(point.x))
		fmt.Printf("Y : %d -> %b\n", point.y, uint16(point.y))
		fmt.Printf("PATH : %d -> %b\n", point.Path(), point.Path())
		t.Fail()
	}

	// Path (-32768, 32767) -> 10000000000000000111111111111111 -> ‭2147516415‬
	if point := New(int16(-32768), int16(32767)); point.Path() != uint32(2147516415) {
		fmt.Printf("X : %d -> %b\n", point.x, uint16(point.x))
		fmt.Printf("Y : %d -> %b\n", point.y, uint16(point.y))
		fmt.Printf("PATH : %d -> %b\n", point.Path(), point.Path())
		t.Fail()
	}

	//  1 -> 0000000000000001 -> 1
	// -1 -> 1111111111111111 -> ‭65535

	// Path (1, -1) -> 00000000000000011111111111111111 -> 131071
	if point := New(int16(1), int16(-1)); point.Path() != uint32(131071) {
		fmt.Printf("X : %d -> %b\n", point.x, uint16(point.x))
		fmt.Printf("Y : %d -> %b\n", point.y, uint16(point.y))
		fmt.Printf("PATH : %d -> %b\n", point.Path(), point.Path())
		t.Fail()
	}

	// Path (-1, 1) -> 11111111111111110000000000000001 -> ‭4294901761
	if point := New(int16(-1), int16(1)); point.Path() != uint32(4294901761) {
		fmt.Printf("X : %d -> %b\n", point.x, uint16(point.x))
		fmt.Printf("Y : %d -> %b\n", point.y, uint16(point.y))
		fmt.Printf("PATH : %d -> %b\n", point.Path(), point.Path())
		t.Fail()
	}

	// Path (32767, 32767) -> 01111111111111110111111111111111 -> ‭2147450879‬
	if point := New(int16(32767), int16(32767)); point.Path() != uint32(2147450879) {
		fmt.Printf("X : %d -> %b\n", point.x, uint16(point.x))
		fmt.Printf("Y : %d -> %b\n", point.y, uint16(point.y))
		fmt.Printf("PATH : %d -> %b\n", point.Path(), point.Path())
		t.Fail()
	}

	// Path (-32768, -32768) -> 10000000000000001000000000000000 -> ‭2147516416‬
	if point := New(int16(-32768), int16(-32768)); point.Path() != uint32(2147516416) {
		fmt.Printf("X : %d -> %b\n", point.x, uint16(point.x))
		fmt.Printf("Y : %d -> %b\n", point.y, uint16(point.y))
		fmt.Printf("PATH : %d -> %b\n", point.Path(), point.Path())
		t.Fail()
	}
}
