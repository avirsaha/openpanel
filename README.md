
# ⚙️ OpenPanel — A FOSS Alternative to WHM/cPanel

**OpenPanel** is a modern, lightweight, open-source server control panel designed for developers, hosting providers, and sysadmins. It enables easy management of Linux-based web servers — including domains, email, DNS, databases, and SSL — without relying on proprietary solutions like WHM/cPanel.

> 🛠️ Built with Go, Nginx, PostgreSQL, and SvelteKit for speed, reliability, and extensibility.

---

## 🌍 Project Vision

OpenPanel is built with three goals in mind:

- 🔒 **Transparent** – 100% open source and auditable
- ⚡ **Efficient** – Minimal resource usage, near-zero latency
- 🧱 **Modular** – Easy to extend, integrate, or replace components

---

## 📦 MVP Scope

The first milestone focuses on building a core, production-ready control panel with:

### ✅ Admin Features
- [x] Root-level user dashboard
- [x] System resource monitoring (RAM, CPU, disk)
- [x] User creation and role assignment
- [x] Domain management (Add/Remove)
- [x] Automated Nginx virtual host configuration
- [x] Free SSL provisioning with Let's Encrypt (`acme.sh`)
- [x] PostgreSQL database creation per domain/user
- [x] FTP/SFTP account setup (via system user or vsftpd)
- [x] Email address provisioning (Postfix + Dovecot)
- [x] Optional DNS record management (via PowerDNS or CoreDNS)

### ✅ User Features
- [x] Domain dashboard with status and usage
- [x] Access to logs and backups
- [x] FTP/Email/DB management for assigned domains

---

## 🧱 Tech Stack

| Layer            | Technology         | Reason                                  |
|------------------|--------------------|-----------------------------------------|
| **Frontend**     | SvelteKit / Vue 3  | Lightweight, fast, reactive             |
| **Backend API**  | Go (Golang)        | High performance, concurrency, safe     |
| **Database**     | PostgreSQL / SQLite| Scalable, reliable                      |
| **Web Server**   | Nginx              | Reverse proxy and virtual host manager |
| **Automation**   | Go + Bash          | Fast system-level control               |
| **SSL**          | acme.sh (Let's Encrypt) | Zero-dependency, automated certs |
| **Mail Server**  | Postfix + Dovecot  | Proven, secure email infrastructure     |
| **DNS**          | PowerDNS / CoreDNS | Optional for full DNS management        |

---

## 🏗️ Project Structure

```

openpanel/
├── cmd/                  # Application entry points
│   └── server/main.go
├── internal/             # Main application logic
│   ├── domain/           # Interfaces and core logic
│   ├── service/          # Business logic implementations
│   ├── handler/          # HTTP route handlers
│   ├── config/           # App configuration
│   ├── store/            # Database implementations
│   ├── system/           # OS integration (nginx, ssl, etc.)
│   └── utils/            # Logging, error handling, etc.
├── web-ui/               # Frontend (SvelteKit or Vue 3)
├── scripts/              # Shell/Ansible scripts
├── deployments/          # Docker, Systemd, CI/CD files
├── tests/                # Unit and integration tests
├── go.mod / go.sum       # Go modules
└── README.md

````

---

## 🚀 Getting Started (Dev Setup)

### 🔧 Prerequisites
- Go >= 1.22
- Node.js >= 20 (for UI)
- PostgreSQL (or SQLite for quick dev)
- Nginx
- Bash (for system scripts)
- acme.sh (optional SSL setup)

### 🖥️ Backend
```bash
# Clone the repo
git clone https://github.com/yourusername/openpanel.git
cd openpanel

# Set up environment
cp .env.example .env

# Run the API server
go run ./cmd/server
````

### 🌐 Frontend

```bash
cd web-ui
npm install
npm run dev
```

### 🐳 Docker (optional)

```bash
docker-compose up --build
```

---

## 📌 Contributing

We welcome contributors! To get started:

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/new-feature`)
3. Commit changes (`git commit -am 'Add new feature'`)
4. Push and create a PR

Please follow our [contribution guide](CONTRIBUTING.md) for standards and tips.

---

## 🛡 Security

OpenPanel will include:

* HTTPS-by-default UI/API
* JWT-based user sessions
* Role-based access control
* Safe OS-level permission isolation
* Fail2Ban / CSF integration (future roadmap)

---

## 🗺 Roadmap Highlights

* [ ] Plugin system for custom modules
* [ ] gRPC support for high-throughput APIs
* [ ] OAuth2 SSO integration
* [ ] Multi-server clustering
* [ ] Web installer + diagnostics tool

---

## 📄 License

This project is licensed under the **MIT License**. See [LICENSE](LICENSE) for details.

---

## 🤝 Community

* Discussions: [GitHub Discussions](https://github.com/yourusername/openpanel/discussions)
* Issues: [GitHub Issues](https://github.com/yourusername/openpanel/issues)
* Roadmap: [Projects tab](https://github.com/yourusername/openpanel/projects)

---

## 🙋‍♂️ Maintainers

**OpenPanel** is developed and maintained by:

* You (and contributors like you!)


