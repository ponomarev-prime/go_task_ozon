# docker run on winows

ubuntu
```
docker run -it --rm -v .\:/home/ubuntu/workspace ubuntu bash
```

go
```
docker run -it --rm -v .\:/app golang:1.22 go run /app/main.go
```