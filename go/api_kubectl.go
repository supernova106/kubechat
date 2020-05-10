/*
 * Swagger Kubechat
 *
 * Wrapper API of kubectl CLI command
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kubechat

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// RunKubectl godoc
// @Summary execute kubectl command
// @Description get kubernetes objects with kubectl
// @Produce  json
// @Success 200 {object} CmdOutput
// @Router /kubectl [get]
// @Param cmd query string true "use by kubectl"
func RunKubectl(c *gin.Context) {
	userCommand := strings.TrimSpace(c.Query("cmd"))
	kctl := &Kubectl{
		Cmd:          userCommand,
		Version:      "1.15.11",
		AllowPattern: `.*`,
	}
	if !validateCommand(kctl) {
		c.JSON(http.StatusOK, &Error{
			Code:    400,
			Message: "command is invalid or not supported",
		})
		return
	}

	cmdOutput := runCommand(userCommand)
	var status int

	if cmdOutput.Code != 0 {
		status = http.StatusBadRequest
	} else {
		status = http.StatusOK
	}

	c.JSON(status, cmdOutput)
}

// validateCommand godoc
func validateCommand(kctl *Kubectl) bool {
	// Only allow get resources
	matched, _ := regexp.MatchString(kctl.AllowPattern, kctl.Cmd)
	return matched
}

// runCommand godoc
func runCommand(command string) *CmdOutput {
	start := time.Now()
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	args := strings.Split(command, " ")
	bin := "kubectl"
	cmd := exec.Command(bin, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	elapsed := time.Since(start)

	output := &CmdOutput{
		Cmd:          command,
		Stderr:       readBuffer(stderr.String()),
		Stdout:       readBuffer(stdout.String()),
		Code:         cmd.ProcessState.ExitCode(),
		ResponseTime: elapsed.String(),
	}

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}

	return output
}

// readBuffer godoc
func readBuffer(data string) []string {
	lines := []string{}
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		l := scanner.Text()
		fmt.Println(l)
		if len(strings.TrimSpace(l)) != 0 {
			lines = append(lines, l)
			continue
		}
	}

	return lines
}
