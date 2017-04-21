all:
	echo "Use lint, bench or prof objectives"

bench:
	cd ./serv/; go test -bench "Bench*"; cd -

prof: api-benchmark
	

api-benchmark:
	go build

lint:
	gometalinter --disable-all --enable=misspell --enable=errcheck --enable=goconst --enable=gosimple --enable=deadcode --enable=aligncheck --enable=unconvert --enable=gas --deadline=100s ./

