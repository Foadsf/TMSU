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

package storage

import (
	"testing"

	"github.com/oniony/TMSU/entities"
)

// A file outside the root is stored with an absolute directory; on Windows that
// starts with a volume, and must not be prefixed with the root again.
func TestAbsPathWindows(test *testing.T) {
	store := &Storage{RootPath: `C:\archive`}

	inside := &entities.File{Directory: `2023\sub`, Name: "a.txt"}
	store.absPath(inside)
	if inside.Directory != `C:\archive\2023\sub` {
		test.Fatalf(`Expected C:\archive\2023\sub but got %v`, inside.Directory)
	}

	outside := &entities.File{Directory: `D:\Downloads\x`, Name: "b.jpg"}
	store.absPath(outside)
	if outside.Directory != `D:\Downloads\x` {
		test.Fatalf(`Expected D:\Downloads\x unchanged but got %v`, outside.Directory)
	}
}
