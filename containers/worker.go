package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func processJob(id string) {
	// initial setup for running input data
	mu.Lock()
	job := jobs[id]
	job.Status = "running"
	jobs[id] = job
	mu.Unlock()

	//creation of file and job to run

	var tmpfile string
	var outfile string //only for c , cpp  , java binary file

	var cmd *exec.Cmd

	switch job.Language {
	case "python":
		tmpfile = fmt.Sprintf("Job_%s.py", job.Id)
		os.WriteFile(tmpfile, []byte(job.Source), 0644)
		cmd = exec.Command("python", tmpfile)

	case "go":
		tmpfile = fmt.Sprintf("Job_%s.go", job.Id)
		os.WriteFile(tmpfile, []byte(job.Source), 0644)
		cmd = exec.Command("go", "run", tmpfile)
	case "node":
		tmpfile = fmt.Sprintf("Job_%s.js", job.Id)
		os.WriteFile(tmpfile, []byte(job.Source), 0644)
		cmd = exec.Command("node", tmpfile)
	case "c":
		tmpfile = fmt.Sprintf("Job_%s.c", job.Id)
		outfile = fmt.Sprintf("Job_%s.out", job.Id)
		os.WriteFile(tmpfile, []byte(job.Source), 0644)

		//compile
		var compileOut bytes.Buffer
		compile := exec.Command("gcc", tmpfile, "-o", outfile)
		compile.Stderr = &compileOut

		if err := compile.Run(); err != nil {
			saveErrorWithStderr(id, err, compileOut.String())
			return
		}
		cmd = exec.Command("./" + outfile)

	case "cpp":
		tmpfile = fmt.Sprintf("Job_%s.cpp", job.Id)
		outfile = fmt.Sprintf("Job_%s.out", job.Id)
		os.WriteFile(tmpfile, []byte(job.Source), 0644)

		//compile
		compile := exec.Command("g++", tmpfile, "-o", outfile)
		if err := compile.Run(); err != nil {
			saveError(id, err)
			return
		}
		cmd = exec.Command("./" + outfile)

	case "java":
		tmpfile = "Main.java"
		className := "Main"
		source := "public class " + className + "{\n " + job.Source + "\n}\n"
		os.WriteFile(tmpfile, []byte(source), 0644)

		compile := exec.Command("javac", tmpfile)
		if err := compile.Run(); err != nil {
			saveErrorWithStderr(id, err, compile.String())
			return
		}
		dir := filepath.Dir(tmpfile)
		cmd = exec.Command("java", "-cp", dir, className)

	default:
		mu.Lock()
		job.Status = "done"
		job.Stderr = "Unsupported language"
		job.ExitCode = 1
		jobs[id] = job
		mu.Unlock()
		return
	}

	if job.Stdin != "" {
		cmd.Stdin = bytes.NewBufferString(job.Stdin)
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	exitcode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitcode = exitErr.ExitCode()
		} else {
			exitcode = 1
		}
	}
	//cleaning up resources
	if tmpfile != "" {
		_ = os.Remove(tmpfile)
	}
	if outfile != "" {
		_ = os.Remove(outfile)
	}
	//update Job in result
	mu.Lock()
	job.Status = "done"
	job.Stdout = stdout.String()
	job.ExitCode = exitcode
	job.Stderr = stderr.String()
	jobs[id] = job
	mu.Unlock()
}

func worker() {
	for id := range Jobqueue {
		processJob(id)
	}
}

func saveError(id string, err error) {
	mu.Lock()
	job := jobs[id]
	job.Status = "done"
	job.ExitCode = 1
	job.Stderr = err.Error()
	jobs[id] = job
	mu.Unlock()
}
func saveErrorWithStderr(id string, err error, stderr string) {
	mu.Lock()
	job := jobs[id]
	job.Status = "done"
	job.Stderr = stderr
	job.ExitCode = 1
	jobs[id] = job
	mu.Unlock()
}
