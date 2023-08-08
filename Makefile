
all: generate bucket counting

generate:
	go run ./number-generator/number-generator.go

bucket_c:
	gcc -o ./bucket-sort/c/bucket ./bucket-sort/c/bucket.c
	./bucket-sort/c/bucket

bucket_go:
	go run ./bucket-sort/go/bucket.go

bucket_py:
	python3 ./bucket-sort/python/bucket.py

bucket: bucket_c bucket_go bucket_py

counting_c:
	gcc -o ./counting-sort/c/counting ./counting-sort/c/counting.c
	./counting-sort/c/counting

counting_go:
	go run ./counting-sort/go/counting.go

counting_py:
	python3 ./counting-sort/python/counting.py

counting: counting_c counting_go counting_py

radix_c:
	gcc -o ./radix-sort/c/radix ./radix-sort/c/radix.c
	./radix-sort/c/radix

radix_go:
	go run ./radix-sort/go/radix.go

radix_py:
	python3 ./radix-sort/python/radix.py

radix: radix_c radix_go radix_py

c: bucket_c counting_c  radix_c

go: bucket_go counting_go  radix_go

py: bucket_py counting_py  radix_py



.PHONY : generate bucket_c bucket_go bucket_py bucket counting_c counting_go counting_py counting
