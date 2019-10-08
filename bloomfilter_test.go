package bloomfilter_test

import (
	"github.com/kenju/go-bloomfilter"
	"testing"
)

func TestNew(t *testing.T) {
	size := uint(1024)

	bf := bloomfilter.New(size)

	if bf.Size() != 0 {
		t.Errorf("expected 0, got %d\n", bf.Size())
	}

	bf.Add([]byte("foo"))
	bf.Add([]byte("bar"))
	bf.Add([]byte("buz"))

	if bf.Size() != 3 {
		t.Errorf("expected 3, got %d\n", bf.Size())
	}

	if bf.Test([]byte("foo")) == false {
		t.Errorf("expected true, got false\n")
	}
	if bf.Test([]byte("bar")) == false {
		t.Errorf("expected true, got false\n")
	}
	if bf.Test([]byte("buz")) == false {
		t.Errorf("expected true, got false\n")
	}

	if bf.Test([]byte("helloworld")) == true {
		t.Errorf("expected false, got true\n")
	}
}

func BenchmarkBloomFilter_Add(b *testing.B) {
	size := uint(1024)
	bf := bloomfilter.New(size)

	for i := 0; i < b.N; i++ {
		bf.Add([]byte("foo"))
	}
}