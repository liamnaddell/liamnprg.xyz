all: sass pyg
	sudo go run ./main.go | tee logfile

sass: clean
	./sass.sh

pyg:
	./pyg.sh
clean:
	rm -f static/css/*.css -v 
	rm -vf logfile
	rm -vf static/code/*.html
	find . -type f -name '*.gz' -exec rm -vf {} \;
