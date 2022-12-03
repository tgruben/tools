package main

import (
	"fmt"
	"os"
	"time"

	"encoding/binary"

	"github.com/tgruben/tools/lists"
)

func DiskWrite(path string, b []byte) {
	defer ExecTime("Write")()
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	l, err := f.Write(b)
	fmt.Println(l, err)
}

func DiskWriteInt(path string, b []uint64) {
	defer ExecTime("WriteUint")()
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for _, v := range b {
		binary.Write(f, binary.LittleEndian, v)
	}
}

func ExecTime(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s execution time: %v\n", name, time.Since(start))
	}
}
func Build() []uint64 {
	defer ExecTime("Build")()
	big := make([]uint64, 268435456)
	for i := range big {
		big[i] = uint64(i)
	}
	return big
}
func Shuffle(l []uint64) {
	defer ExecTime("Shuffle")()
	lists.Shuffle(l)
}
func Sum(l []uint64) uint64 {
	defer ExecTime("Sum")()
	return lists.Sum(l)
}

func Rotate(l []uint64) {
	defer ExecTime("Rotate")()
	lists.Rotate(l)
}

func Fold(l []uint64) uint64 {
	defer ExecTime("Fold")()
	return lists.Fold(l, func(x, y uint64) uint64 {
		return x + y
	})
}

func main() {
	l := Build()
	Shuffle(l)
	x := Sum(l)
	fmt.Println(x)
	Rotate(l)
	x = Fold(l)
	fmt.Println(x)
	//sz := int(len(l) * int(unsafe.Sizeof(l[0])))
	//2147483648
	//b := (*[2147483648]byte)(unsafe.Pointer(&l[0]))[:2147483648:2147483648]
	DiskWriteInt("test", l)

}
