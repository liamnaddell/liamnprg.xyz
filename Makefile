all: sass
	sudo go run ./main.go

sass: clean
	./sass.sh

clean:
	rm -f static/css/*.css -v 
	rm -vf logfile
	rm -vf static/*.gz static/img/*.gz static/css/*.gz
