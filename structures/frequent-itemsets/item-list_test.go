package structures

import (
	"testing"
)

func TestHeaderOnly(t *testing.T) {

	var dataset *[]*[]uint = &[]*[]uint{
		&[]uint{1, 2, 4, 5, 6, 7, 8, 9},
		&[]uint{2, 3, 5, 12},
		&[]uint{1, 2, 4, 6, 8, 12},
		&[]uint{1, 2, 3, 4, 6, 7, 12},
		&[]uint{2, 7, 8, 12},
		&[]uint{1, 2, 4, 6, 9}}

	for _, tx := range *dataset {
		for _, e := range *tx {
			t.Logf("%v ", ToCharStr(e))
		}
		t.Logf("\n")
	}

	il := NewItemList(dataset, 3)

	t.Logf("%v", il)

	output := "B [6]: [A B D E F G H I] + [B C E L] + [A B D F H L] + [A B C D F G L] + [B G H L] + [A B D F I]\n" +
		"A [4]: [A B D E F G H I] + [A B D F H L] + [A B C D F G L] + [A B D F I]\n" +
		"D [4]: [A B D E F G H I] + [A B D F H L] + [A B C D F G L] + [A B D F I]\n" +
		"F [4]: [A B D E F G H I] + [A B D F H L] + [A B C D F G L] + [A B D F I]\n" +
		"L [4]: [B C E L] + [A B D F H L] + [A B C D F G L] + [B G H L]\n" +
		"G [3]: [A B D E F G H I] + [A B C D F G L] + [B G H L]\n" +
		"H [3]: [A B D E F G H I] + [A B D F H L] + [B G H L]\n"

	if il.String() != output {
		t.Fatalf("Output is not correct.\nIt is: \n%v\nExpected: \n%v\n", il, output)
	}
	// if pt.header.right != nil {
	// }
}
