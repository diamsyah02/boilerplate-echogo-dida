# Boilerplate-EchoGo-Dida
Boilerplate Echo Golang

# Tech Stack
- Golang
- Echo
- GORM
- JWT
- Zerolog
- Assert
- Asynq
- Go Mail
- Asynqmon

# Build App in Linux
```bash
GOOS=linux GOARCH=amd64 go build -o app ./cmd/app/
```

# Build Worker in Linux
```bash
GOOS=linux GOARCH=amd64 go build -o worker ./cmd/worker/
```

# Build App in Windows
```bash
GOOS=windows GOARCH=amd64 go build -o app.exe ./cmd/app/
```

# Build Worker in Windows
```bash
GOOS=windows GOARCH=amd64 go build -o worker.exe ./cmd/worker/
```

# Run Auto Start
- Buat file /etc/systemd/system/app.service
```bash
[Unit]
Description=Boilerplate Echo Golang
After=network.target

[Service]
User=boilerplate
WorkingDirectory=/home/boilerplate/app
ExecStart=/home/boilerplate/app/app
EnvironmentFile=/home/boilerplate/app/.env
Restart=always
RestartSec=5
Environment=PORT=5000
StandardOutput=append:/home/boilerplate/app/logs/app.log
StandardError=append:/home/boilerplate/app/logs/app.err.log

[Install]
WantedBy=multi-user.target
```

- Reload systemd
```bash
sudo systemctl daemon-reload
```

- Restart service
```bash
sudo systemctl daemon-reexec
```

- Start service
```bash
sudo systemctl start app
```

- Enable service
```bash
sudo systemctl enable --now app
```
