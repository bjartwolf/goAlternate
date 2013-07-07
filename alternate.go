package main
import ("fmt"
	"testing"
	"flag"
	"os"
	"log"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func AlternateBranch(b *testing.B) {
    a := 13
    c := 9
    x := a
	for i := 0; i < b.N; i++ {
		if (x == a) {
			x = c
		} else {
			x = a
		}
	}
}

func AlternateXor(b *testing.B) {
    a := 13
    c := 9
    x := a
	for i := 0; i < b.N; i++ {
		x = a^c^x
	}
}

func main () {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	fmt.Println("Normal switch") 
    fmt.Println(testing.Benchmark(AlternateBranch))
	fmt.Println("Fast switch")
	fmt.Println(testing.Benchmark(AlternateXor))
}
