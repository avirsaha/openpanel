package users

import (
    "fmt"
    "os/exec"
    "strings"
    "errors"
)

// Create a Linux user with home directory and password
func CreateUser(username, password string) error {
    if UserExists(username) {
        return errors.New("user already exists")
    }

    cmd := exec.Command("useradd", "-m", "-s", "/bin/bash", username)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to create user: %w", err)
    }

    // Set password (WARNING: use securely in production)
    passCmd := exec.Command("chpasswd")
    passCmd.Stdin = strings.NewReader(fmt.Sprintf("%s:%s", username, password))
    if err := passCmd.Run(); err != nil {
        return fmt.Errorf("failed to set password: %w", err)
    }

    return nil
}

func DeleteUser(username string) error {
    cmd := exec.Command("userdel", "-r", username)
    return cmd.Run()
}

func UserExists(username string) bool {
    cmd := exec.Command("id", username)
    return cmd.Run() == nil
}

func ListUsers() ([]string, error) {
    out, err := exec.Command("cut", "-d:", "-f1", "/etc/passwd").Output()
    if err != nil {
        return nil, err
    }

    lines := strings.Split(string(out), "\n")
    var panelUsers []string
    for _, u := range lines {
        if u == "root" || u == "" {
            continue
        }
        panelUsers = append(panelUsers, u)
    }
    return panelUsers, nil
}
