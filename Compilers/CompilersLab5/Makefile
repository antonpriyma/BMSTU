.PHONY: goimports
goimports: bin/goimports
	find .\
    -type f\
    -name '*.go'\
    \! -name '*mock_generated.go' | xargs bin/goimports -w


