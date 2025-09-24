package web

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "text/template"
)

const nginxTemplate = `
server {
    listen 80;
    server_name {{.Domain}};

    root /home/{{.Username}}/public_html;
    index index.html index.htm;

    location / {
        try_files $uri $uri/ =404;
    }
}
`

type WebsiteConfig struct {
    Domain   string
    Username string
}

func CreateWebsite(config WebsiteConfig) error {
    // 1. Generate config file content
    tpl, err := template.New("nginx").Parse(nginxTemplate)
    if err != nil {
        return fmt.Errorf("failed to parse template: %w", err)
    }

    configPath := filepath.Join("/etc/nginx/sites-available", config.Domain+".conf")
    file, err := os.Create(configPath)
    if err != nil {
        return fmt.Errorf("failed to create config file: %w", err)
    }
    defer file.Close()

    if err := tpl.Execute(file, config); err != nil {
        return fmt.Errorf("failed to write config: %w", err)
    }

    // 2. Symlink to sites-enabled
    enabledPath := filepath.Join("/etc/nginx/sites-enabled", config.Domain+".conf")
    if err := os.Symlink(configPath, enabledPath); err != nil && !os.IsExist(err) {
        return fmt.Errorf("failed to enable site: %w", err)
    }

    // 3. Reload nginx
    if err := exec.Command("systemctl", "reload", "nginx").Run(); err != nil {
        return fmt.Errorf("failed to reload nginx: %w", err)
    }

    return nil
}
