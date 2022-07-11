package utils

import "os/exec"

func ExecuteVbsScript(vdfPathScript string) error {
	var linkCreationCommandExecution *exec.Cmd = exec.Command(
		"wscript",

		vdfPathScript,
	)

	var err = linkCreationCommandExecution.Run()

	if err != nil {
		return err
	}

	return linkCreationCommandExecution.Wait()
}
