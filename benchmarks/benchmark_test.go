package benchmarks

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

var names = []string{
    //"test_files/file_1GB.txt",
    "test_files/file_50MB.txt",
}

func BenchmarkCat(b *testing.B) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    b.Logf("Before reading file: Alloc = %v MiB", m.Alloc/1024/1024)

    for i := 0; i < b.N; i++ {
        for _, filename := range names {
            cmd := exec.Command("cli", "cat", filename)

            cmd.Env = append(os.Environ(), "CONFIG=../config.yml")

            output, err := cmd.CombinedOutput() 
            if err != nil {
                b.Errorf("Error running command: %v, output: %s", err, output)
                return
            }
        }
    }
    runtime.ReadMemStats(&m)
    b.Logf("After reading file: Alloc = %v MiB", m.Alloc/1024/1024)
}



func BenchmarkCount(b *testing.B) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Before reading file: Alloc = %v MiB\n", m.Alloc/1024/1024)
    
    for _, filename := range names {
        cmd := exec.Command("cli", "count", filename)

        cmd.Env = append(os.Environ(), "CONFIG=../config.yml")

        output, err := cmd.CombinedOutput()
        if err != nil {
            b.Errorf("Error running command: %v", string(output))
        }
    }
    
    runtime.ReadMemStats(&m)
    fmt.Printf("After reading file: Alloc = %v MiB\n", m.Alloc/1024/1024)
}

func BenchmarkCountS(b *testing.B) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Before reading file: Alloc = %v MiB\n", m.Alloc/1024/1024)
    
    for _, filename := range names {
        cmd := exec.Command("cli", "countS", filename)

        cmd.Env = append(os.Environ(), "CONFIG=../config.yml")
        
        output, err := cmd.CombinedOutput()
        if err != nil {
            b.Errorf("Error running command: %v", string(output))
        }
    }
    
    runtime.ReadMemStats(&m)
    fmt.Printf("After reading file: Alloc = %v MiB\n", m.Alloc/1024/1024)
}

