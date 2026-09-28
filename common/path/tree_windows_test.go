// Copyright 2011-2018 Paul Ruane.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package path

import (
	"testing"
)

// Drive-letter paths must come back out of the tree intact; previously the
// synthetic "/" root turned C:\a\b into \C:a\b, breaking `tmsu status`.
func TestPathsWindowsVolume(test *testing.T) {
	tree := NewTree()
	tree.Add(`C:\root\a\x.txt`, false)
	tree.Add(`C:\root\b`, true)
	tree.Add(`D:\other\y.txt`, false)

	paths := tree.Paths()
	expected := []string{`C:\root\a\x.txt`, `C:\root\b`, `D:\other\y.txt`}
	if len(paths) != len(expected) {
		test.Fatalf("Expected %v paths but got %v: %v", len(expected), len(paths), paths)
	}
	for i := range expected {
		if paths[i] != expected[i] {
			test.Fatalf("Expected path %v to be '%v' but was '%v'.", i, expected[i], paths[i])
		}
	}

	top := tree.TopLevel().Paths()
	expectedTop := []string{`C:\root\a\x.txt`, `C:\root\b`, `D:\other\y.txt`}
	for i := range expectedTop {
		if top[i] != expectedTop[i] {
			test.Fatalf("Expected top-level path %v to be '%v' but was '%v'.", i, expectedTop[i], top[i])
		}
	}
}

func TestPathsWindowsVolumeRootOnly(test *testing.T) {
	tree := NewTree()
	tree.Add(`C:`, true)

	paths := tree.Paths()
	if len(paths) != 1 || paths[0] != `C:\` {
		test.Fatalf(`Expected [C:\] but got %v`, paths)
	}
}
