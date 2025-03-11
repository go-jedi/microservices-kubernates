### migrations:

#### create:
- `migrate create -ext sql -dir migrations -seq posts_table`

#### execute:
- `migrate -database postgresql://admin:test@localhost:5432/test_db?sslmode=disable -path migrations up`
- `migrate -database postgresql://admin:test@localhost:5432/test_db?sslmode=disable -path migrations down`

#### build application:
- `go build -ldflags="-s -w" -trimpath -buildvcs=false -o app cmd/app/main.go`

#### run application in systemd:
- `cd /etc/systemd/system`
- `создать openai-service.service`
- `sudo systemctl daemon-reload`
- `sudo systemctl start openai-service.service`
- `sudo systemctl status openai-service.service`
- `sudo systemctl enable openai-service.service`

#### Включить порт:
- `sudo ufw allow 50051/tcp`
- `sudo ufw reload`
- если при выполнении команды: sudo ss -tuln | grep 50051 у вас показывается:
  tcp    LISTEN  0       4096         127.0.0.1:50051        0.0.0.0:*
  , то это указывает, что сервис будет доступен только внутри сервера через localhost.
- Если нужно, чтобы можно было отправлять запросы из внешних источников:
- в сервисе указываем при запуске http сервера :50051
- выполняем sudo ss -tuln | grep 50051 и должно быть в ответе такой результат:
  tcp    LISTEN  0       4096           0.0.0.0:50051        0.0.0.0:*

#### remove all local branch without main:
- `git branch | grep -v "main" | xargs git branch -D`