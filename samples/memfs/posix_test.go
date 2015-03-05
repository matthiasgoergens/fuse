// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Tests for the behavior of os.File objects on plain old posix file systems,
// for use in verifying the intended behavior of memfs.

package memfs_test

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/jacobsa/ogletest"
)

func TestPosix(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Boilerplate
////////////////////////////////////////////////////////////////////////

type PosixTest struct {
	dir string
}

var _ SetUpInterface = &PosixTest{}
var _ TearDownInterface = &PosixTest{}

func init() { RegisterTestSuite(&PosixTest{}) }

func (t *PosixTest) SetUp(ti *TestInfo) {
	var err error

	// Create a temporary directory.
	t.dir, err = ioutil.TempDir("", "posix_test")
	if err != nil {
		panic(err)
	}
}

func (t *PosixTest) TearDown() {
	// Remove the temporary directory.
	err := os.RemoveAll(t.dir)
	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////////////////////////
// Test functions
////////////////////////////////////////////////////////////////////////

func (t *PosixTest) DoesFoo() {
	AssertTrue(false, "TODO")
}
