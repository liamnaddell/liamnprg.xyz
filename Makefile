all: sass pyg
	sudo go run ./main.go | tee logfile

sass: clean
	./sass.sh

pyg:
	./pyg.sh
clean:
	rm -f static/css/*.css -v 
	rm -vf logfile
	rm -vf static/*.gz static/img/*.gz static/css/*.gz
