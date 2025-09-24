package system

import (
    "fmt"
    "os/exec"
)

// CheckDependencies checks if required system tools are installed
func CheckDependencies() error {
    deps := []string{"nginx", "php", "mysql", "certbot"}

    for _, dep := range deps {
        if !isCommandAvailable(dep) {
            return fmt.Errorf("missing required dependency: %s", dep)
        }
    }

    return nil
}

// isCommandAvailable checks if a command exists in PATH
func isCommandAvailable(name string) bool {
    _, err := exec.LookPath(name)
    return err == nil
}

