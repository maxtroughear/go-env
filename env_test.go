package goenv

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	godotenv.Load("test.env")
	os.Exit(m.Run())
}

func TestCanGetNotDefault(t *testing.T) {
	def := "fail"
	expect := "success"
	r := CanGet("testdefault", def)

	if r != expect {
		t.Errorf("Failed to get the correct env var for CanGet, got: %s, want: %s", r, expect)
	}
}

func TestCanGetDefault(t *testing.T) {
	def := "fail"
	expect := def
	r := CanGet("testnondefault", def)

	if r != expect {
		t.Errorf("Failed to get the default env var for CanGet, got: %s, want: %s", r, expect)
	}
}

func TestCanGetFileNotDefault(t *testing.T) {
	def := "fail"
	expect := "success text"
	r := CanGet("testfile", def)

	if r != expect {
		t.Errorf("Failed to get correct env var from file for CanGet, go: %s, want %s", r, expect)
	}
}

func TestCanGetFileDefault(t *testing.T) {
	def := "success"
	expect := def
	r := CanGet("testfile_notexistant", def)

	if r != expect {
		t.Errorf("Failed to get the default env var from file for CanGet, go: %s, want %s", r, expect)
	}
}

func TestCanGetSlice(t *testing.T) {
	def := []string{
		"default",
		"slice",
	}
	expect := []string{
		"one",
		"two",
		"three",
		"four",
	}
	r := CanGetSlice("testslice", def)

	if len(r) != len(expect) {
		t.Errorf("Failed to get the correct env var for CanGetSlice, got: %s, want: %s", r, expect)
	}
	for i, v := range r {
		if v != expect[i] {
			t.Errorf("Failed to get the correct env var for CanGetSlice, got: %s, want: %s", r, expect)
		}
	}
}

func TestCanGetSliceDefault(t *testing.T) {
	def := []string{
		"default",
		"slice",
	}
	expect := def
	r := CanGetSlice("testslicedefault", def)

	if len(r) != len(expect) {
		t.Errorf("Failed to get the correct env var for CanGetSlice, got: %s, want: %s", r, expect)
	}
	for i, v := range r {
		if v != expect[i] {
			t.Errorf("Failed to get the correct env var for CanGetSlice, got: %s, want: %s", r, expect)
		}
	}
}

func TestCanGetInt32(t *testing.T) {
	def := int32(10)
	expect := int32(5)
	r := CanGetInt32("testint32", def)

	if r != expect {
		t.Errorf("Failed to get the correct env var for CanGetInt32, got: %d, want: %d", r, expect)
	}
}

func TestCanGetInt32Default(t *testing.T) {
	def := int32(10)
	expect := def
	r := CanGetInt32("testint32default", def)

	if r != expect {
		t.Errorf("Failed to get the default env var for CanGetInt32, got: %d, want: %d", r, expect)
	}
}

func TestCanGetInt32Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Failed to panic CanGetInt32 with invalid var")
		}
	}()

	CanGetInt32("testint32panic", 10)
}

func TestMustGet(t *testing.T) {
	expect := "mustgetstring"
	r := MustGet("testmust")

	if r != expect {
		t.Errorf("Failed to get the default env var for CanGetInt32, got: %s, want: %s", r, expect)
	}
}

func TestMustGetPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Failed to panic MustGet with no var")
		}
	}()

	MustGet("testmustpanic")
}

func TestMustGetSlice(t *testing.T) {
	expect := []string{
		"four",
		"three",
		"two",
		"one",
	}
	r := MustGetSlice("testslicemust")

	if len(r) != len(expect) {
		t.Errorf("Failed to get the correct env var for MustGetSlice, got: %s, want: %s", r, expect)
	}
	for i, v := range r {
		if v != expect[i] {
			t.Errorf("Failed to get the correct env var for MustGetSlice, got: %s, want: %s", r, expect)
		}
	}
}

func TestMustGetSlicePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Failed to panic MustGet with no var")
		}
	}()
	MustGetSlice("testslicemustpanic")
}

func TestMustGetBool(t *testing.T) {
	expect := true
	r := MustGetBool("testboolmust")

	if r != expect {
		t.Errorf("Failed to get the default env var for CanGetInt32, got: %t, want: %t", r, expect)
	}
}

func TestMustGetBoolString(t *testing.T) {
	expect := true
	r := MustGetBool("testboolmuststring")

	if r != expect {
		t.Errorf("Failed to get the default env var for CanGetInt32, got: %t, want: %t", r, expect)
	}
}

func TestMustGetBoolNumber(t *testing.T) {
	expect := true
	r := MustGetBool("testboolmustnumber")

	if r != expect {
		t.Errorf("Failed to get the default env var for CanGetInt32, got: %t, want: %t", r, expect)
	}
}

func TestMustGetBoolPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Failed to panic MustGetBool with no var")
		}
	}()

	MustGetBool("testboolmustpanic")
}

func TestMustGetBoolPanicInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Failed to panic MustGetBool with invalid var")
		}
	}()

	MustGetBool("testboolmustpanicnumber")
}

func TestMustGetInt32(t *testing.T) {
	expect := int32(15)
	r := MustGetInt32("testint32must")

	if r != expect {
		t.Errorf("Failed to get the correct env var for MustGetInt32, got: %d, want: %d", r, expect)
	}
}

func TestMustGetInt32Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Failed to panic MustGetInt32 with no var")
		}
	}()

	MustGetInt32("testint32mustpanic")
}

func TestMustGetInt32PanicInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Failed to panic MustGetInt32 with invalid var")
		}
	}()

	MustGetInt32("testint32mustpanicinvalid")
}

func TestMustGetInt64(t *testing.T) {
	expect := int64(157816845615315818)
	r := MustGetInt64("testint64must")

	if r != expect {
		t.Errorf("Failed to get the correct env var for MustGetInt64, got: %d, want: %d", r, expect)
	}
}

func TestMustGetInt64Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Failed to panic MustGetInt64 with no var")
		}
	}()

	MustGetInt64("testint64mustpanic")
}

func TestMustGetInt64PanicInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Failed to panic MustGetInt64 with invalid var")
		}
	}()

	MustGetInt64("testint64mustpanicinvalid")
}
