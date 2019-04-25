package iostruct_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/nylo-andry/iostruct"
)

type Foo struct {
	Bar string
}

func TestReadWrite(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	filename := fmt.Sprintf("%v.testdata", r.Int())

	f := Foo{fmt.Sprintf("%v", r.Int())}
	err := iostruct.Write(filename, f)
	if err != nil {
		t.Errorf("Could not write test data file: %v", err)
	}

	var f2 Foo
	err = iostruct.Read(filename, &f2)
	if err != nil {
		t.Errorf("Could not read test data file: %v", err)
	}

	if f.Bar != f2.Bar {
		t.Error("The written and read struct didn't match")
	}

	os.Remove(filename)
}

func TestReadWriteEncrypted(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	filename := fmt.Sprintf("%v.testdata", r.Int())
	passphrase := fmt.Sprintf("%v", r.Int())

	f := Foo{fmt.Sprintf("%v", r.Int())}
	err := iostruct.WriteEncrypted(filename, passphrase, f)
	if err != nil {
		t.Errorf("Could not write test data file: %v", err)
	}

	var f2 Foo
	err = iostruct.ReadEncrypted(filename, passphrase, &f2)
	if err != nil {
		t.Errorf("Could not read test data file: %v", err)
	}

	if f.Bar != f2.Bar {
		t.Error("The written and read struct didn't match")
	}

	os.Remove(filename)
}
