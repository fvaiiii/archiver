build:
	go build -o archiver .

help: build
	./archiver --help

compress-test: build
	./archiver compress README.md readme.arc --window 32768

decompress-test: build
	./archiver decompress readme.arc readme_restored.md --force

clean:
	rm -f archiver readme.arc readme_restored.md

test-cli: build compress-test decompress-test