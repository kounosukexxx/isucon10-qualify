all: isuumo

isuumo: *.go
	git pull
	go build -o isuumo
	sudo systemctl restart isuumo.go.service
