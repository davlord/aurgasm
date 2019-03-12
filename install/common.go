package install

import(
	"os"
	"os/exec"
	"davlord.com/aurgasm/common"
)

func buildInfoURL(packageName string) string {
	return common.AurAPIBase + common.AurAPIInfoPath + packageName
}

func aurAPIInfoPackage(packageName string, searchResult *common.SearchResult) error {
	url := buildInfoURL(packageName)
	return common.GetJson(url, searchResult)
}

func runCommand(dir string, cmdName string, cmdArgs ...string) (error) {
	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Dir = dir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()	
}