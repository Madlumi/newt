
outname = newt
run:
	go run ./src/...
build:
	-mkdir out
	go build -o out ./src/...
	mv out/src out/$(outname)
local: build
	mv out/$(outname) ~/scripts/runnable/
clean:
	rm out/$(outname)
	install: build
	echo "not implemented"

