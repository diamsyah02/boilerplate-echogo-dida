# boilerplate-echogo-dida
Boilerplate Echo Golang

# Tech Stack
- Golang
- Echo
- GORM
- JWT
- Zerolog

# Build
```bash
GOOS=linux GOARCH=amd64 go build -o myapp main.go
```

# Run Auto Start
- Buat file /etc/systemd/system/myapp.service
```bash
[Unit]
Description=Boilerplate Echo Golang
After=network.target

[Service]
User=boilerplate
WorkingDirectory=/home/boilerplate/app
ExecStart=/home/boilerplate/app/myapp
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
sudo systemctl start myapp
```

- Enable service
```bash
sudo systemctl enable --now myapp
```
