package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("CLI Template Setup")

	originURL := detectOriginURL()

	projectParts := strings.Split(strings.TrimPrefix(originURL, "https://github.com/"), "/")
	repoUser := projectParts[0]
	repoName := projectParts[1]

	project := struct {
		Username    string
		Reponame    string
		ProjectURL  string
		ProjectName string
	}{}

	project.Username = repoUser
	project.Reponame = repoName
	project.ProjectURL = fmt.Sprintf("github.com/%s/%s", project.Username, project.Reponame)
	project.ProjectName = fmt.Sprintf("%s/%s", project.Username, project.Reponame)

	const cliTemplatePath = "KarolosLykos/cli-template"
	fmt.Printf("Replacing all '%s', with %s", cliTemplatePath, project.ProjectName)
	walkOverExt("go,mod", func(path string) {
		fmt.Printf("Replacing '%s' in %s with %s", cliTemplatePath, path, project.ProjectName)
		replaceAllInFile(path, cliTemplatePath, project.ProjectName)
	})

	replaceAllInFile("./cmd/root.go", `Use:   "cli-template",`, fmt.Sprintf(`Use:   "%s",`, project.Reponame))
	replaceAllInFile("./README.md", cliTemplatePath, project.ProjectName)
	replaceAllInFile("./.golangci.yml", cliTemplatePath, project.ProjectName)

	if err := os.RemoveAll(getPathTo("./setup")); err != nil {
		panic(err)
	}
}

func walkOverExt(exts string, f func(path string)) {
	_ = filepath.Walk(getPathTo(""), func(path string, info fs.FileInfo, err error) error {
		for _, ext := range strings.Split(exts, ",") {
			if filepath.Ext(path) == "."+ext {
				f(path)
			}
		}
		return nil
	})
}

func detectOriginURL() (url string) {
	out, err := exec.Command("git", "remote", "-v").Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Git output:\n%s", string(out))

	output := string(out)

	for _, s := range strings.Split(output, "\n") {
		s = strings.TrimSpace(strings.TrimPrefix(s, "origin"))
		if strings.HasPrefix(s, "https://github.com/") && strings.Contains(s, "push") {
			url = strings.TrimSpace(strings.TrimRight(s, "(push)"))

			return
		}

		if strings.HasPrefix(s, "git@github.com:") && strings.Contains(s, "push") {
			url = strings.TrimSpace(strings.TrimRight(s, "(push)"))

			return
		}
	}

	return
}

func replaceAllInFile(filepath, search, replace string) {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	content := string(fileBytes)
	content = strings.ReplaceAll(content, search, replace)

	if err := ioutil.WriteFile(
		filepath,
		[]byte(content),
		fs.ModePerm,
	); err != nil {
		panic(err)
	}
}

func getPathTo(file string) string {
	_, scriptPath, _, _ := runtime.Caller(1)
	return filepath.Join(scriptPath, "../../", file)
}
