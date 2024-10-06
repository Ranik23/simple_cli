package benchmarks

import (
	_ "net/http/pprof"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

func BenchmarkCatFile1GB(b *testing.B) {
    benchmarkCat(b, "test_files/file_1GB.txt")
}

func BenchmarkCatFile50MB(b *testing.B) {
    benchmarkCat(b, "test_files/file_50MB.txt")
}

func BenchmarkCountFile1GB(b *testing.B) {
    benchmarkCount(b, "test_files/file_1GB.txt")
}

func BenchmarkCountFile50MB(b *testing.B) {
    benchmarkCount(b, "test_files/file_50MB.txt")
}

func BenchmarkCountSFile50MB(b *testing.B) {
    benchmarkCountS(b, "test_files/file_50MB.txt")
}

func BenchmarkCountSFile1GB(b *testing.B) {
    benchmarkCountS(b, "test_files/file_1GB.txt")
}

func benchmarkCat(b *testing.B, filename string) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    b.Logf("Before reading file: Alloc = %v KiB", m.Alloc/1024)

    for i := 0; i < b.N; i++ {
        cmd := exec.Command("cli", "cat", filename)
        cmd.Env = append(os.Environ(), "CONFIG=../config.yml")

        output, err := cmd.CombinedOutput() 
        if err != nil {
            b.Errorf("Error running command: %v, output: %s", err, output)
            return
        }
    }
    runtime.ReadMemStats(&m)
    b.Logf("After reading file: Alloc = %v KiB", m.Alloc/1024)
}

func benchmarkCount(b *testing.B, filename string) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    b.Logf("Before reading file: Alloc = %v KiB", m.Alloc/1024)

    for i := 0; i < b.N; i++ {
        cmd := exec.Command("cli", "count", filename)

        cmd.Env = append(os.Environ(), "CONFIG=../config.yml")

        output, err := cmd.CombinedOutput() 
        if err != nil {
            b.Errorf("Error running command: %v, output: %s", err, output)
            return
        }
    }
    runtime.ReadMemStats(&m)
    b.Logf("After reading file: Alloc = %v KiB", m.Alloc/1024)
}

func benchmarkCountS(b *testing.B, filename string) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    b.Logf("Before reading file: Alloc = %v KiB", m.Alloc/1024)

    for i := 0; i < b.N; i++ {
        cmd := exec.Command("cli", "countS", filename)

        cmd.Env = append(os.Environ(), "CONFIG=../config.yml")

        output, err := cmd.CombinedOutput() 
        if err != nil {
            b.Errorf("Error running command: %v, output: %s", err, output)
            return
        }
    }
    runtime.ReadMemStats(&m)
    b.Logf("After reading file: Alloc = %v KiB", m.Alloc/1024)
}

