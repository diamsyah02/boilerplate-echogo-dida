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

# Build Linux
```bash
GOOS=linux GOARCH=amd64 go build -o app ./cmd/app/
```

# Build Windows
```bash
GOOS=windows GOARCH=amd64 go build -o app.exe ./cmd/app/
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
Restart=always
Environment=PORT=5000

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
