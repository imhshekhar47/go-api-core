EXE_NAME=api-core

clean:
	if [ "./build" ]; then rm -rf ./build; fi;

test: 
	if [ ! -d "./build/tests" ]; then mkdir -p "./build/tests"; fi;
	gotestsum --format testname \
			--junitfile "./build/tests/unit-tests.xml" \
			-- -coverprofile=build/coverage.out ./...
build: test
	if [ ! -d "./build/bin/" ]; then mkdir -p ./build/bin; fi;
	go build -o ./build/bin/${EXE_NAME}