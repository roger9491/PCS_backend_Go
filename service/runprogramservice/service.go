package runprogramservice
import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"gorm.io/gorm"
)

var (
	programLanguage = "python"
)

func RunProgram(programText string, db *gorm.DB) (output string, err error) {

	file, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(file.Name())

	_, err = file.WriteString(programText)
	if err != nil {
		log.Println(err)
		return
	}

	cmd := exec.Command(programLanguage, file.Name())

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return
	}
	output = string(out)

	return
}
