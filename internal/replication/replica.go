package replication

import (
	"fmt"
	"os/exec"
)


var (
	ErrFailedToInitDB = fmt.Errorf("failed to init database")
)


func InitDB(filePath string) {

	cmd := exec.Command("initdb", "-D", filePath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
	}
}

func Pg_BaseBackUp(host string, user string, filepath string, port string) {

	Port := "--port=" + port

	cmd := exec.Command("pg_basebackup",
		"-h", host,
		"-U", user,
		"--checkpoint=fast",
		"-D", filepath,
		"-R",
		"--slot=some_name",
		"-C",
		Port,
	)
	output, _ := cmd.CombinedOutput()
	
	fmt.Println(string(output))
}

func Pg_Ctl_Stop(filepath string) {
	cmd := exec.Command("pg_ctl", "-D", filepath, "stop")

    output, _ := cmd.CombinedOutput()
 
    fmt.Println(string(output))
}

func Pg_Ctl_Start(filepath string) {
	cmd := exec.Command("pg_ctl", "-D", filepath, "start")

    output, _ := cmd.CombinedOutput()
 
    fmt.Println(string(output))
}


