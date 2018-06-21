all: 
	docker build -t liamnprg-webserver .
	docker run -p 80:80 -it --rm --name liamnprg-running liamnprg-webserver
