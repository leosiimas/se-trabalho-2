
all: generate compile main

compile:
	gcc -o ./algorithms/bucket-sort/c/bucket.o ./algorithms/bucket-sort/c/bucket.c
	go build -o ./algorithms/bucket-sort/go/bucket.o ./algorithms/bucket-sort/go/bucket.go
	rustc -o ./algorithms/bucket-sort/rust/bucket.o ./algorithms/bucket-sort/rust/bucket.rs
	gcc -o ./algorithms/counting-sort/c/counting.o ./algorithms/counting-sort/c/counting.c
	go build -o ./algorithms/counting-sort/go/counting.o ./algorithms/counting-sort/go/counting.go
	rustc -o ./algorithms/counting-sort/rust/counting.o ./algorithms/counting-sort/rust/counting.rs
	gcc -o ./algorithms/radix-sort/c/radix.o ./algorithms/radix-sort/c/radix.c
	go build -o ./algorithms/radix-sort/go/radix.o ./algorithms/radix-sort/go/radix.go
	rustc -o ./algorithms/radix-sort/rust/radix.o ./algorithms/radix-sort/rust/radix.rs

main:
	go run ./main.go

generate:
	go run ./number-generator/number-generator.go

bucket_c:
	gcc -o ./algorithms/bucket-sort/c/bucket.o ./algorithms/bucket-sort/c/bucket.c
	./algorithms/bucket-sort/c/bucket.o

bucket_go:
	go build -o ./algorithms/bucket-sort/go/bucket.o ./algorithms/bucket-sort/go/bucket.go
	./algorithms/bucket-sort/go/bucket.o

bucket_rs:
	rustc -o ./algorithms/bucket-sort/rust/bucket.o ./algorithms/bucket-sort/rust/bucket.rs
	./algorithms/bucket-sort/rust/bucket.o

bucket: bucket_c bucket_go bucket_rs

counting_c:
	gcc -o ./algorithms/counting-sort/c/counting.o ./algorithms/counting-sort/c/counting.c
	./algorithms/counting-sort/c/counting.o

counting_go:
	go build -o ./algorithms/counting-sort/go/counting.o ./algorithms/counting-sort/go/counting.go
	./algorithms/counting-sort/go/counting.o

counting_rs:
	rustc -o ./algorithms/counting-sort/rust/counting.o ./algorithms/counting-sort/rust/counting.rs
	./algorithms/counting-sort/rust/counting.o

counting: counting_c counting_go counting_rs

radix_c:
	gcc -o ./algorithms/radix-sort/c/radix.o ./algorithms/radix-sort/c/radix.c
	./algorithms/radix-sort/c/radix.o

radix_go:
	go build -o ./algorithms/radix-sort/go/radix.o ./algorithms/radix-sort/go/radix.go
	./algorithms/radix-sort/go/radix.o

radix_rs:
	rustc -o ./algorithms/radix-sort/rust/radix.o ./algorithms/radix-sort/rust/radix.rs
	./algorithms/radix-sort/rust/radix.o

radix: radix_c radix_go radix_rs

c: bucket_c counting_c  radix_c

go: bucket_go counting_go  radix_go

rs: bucket_rs counting_rs  radix_rs

.PHONY : generate bucket_c bucket_go bucket_rs bucket counting_c counting_go counting_rs counting main compile
