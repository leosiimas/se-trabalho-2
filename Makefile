
all: generate bucket counting radix

generate:
	go run ./number-generator/number-generator.go

bucket_c:
	gcc -o ./bucket-sort/c/bucket.o ./bucket-sort/c/bucket.c
	./bucket-sort/c/bucket.o

bucket_go:
	go build -o ./bucket-sort/go/bucket.o ./bucket-sort/go/bucket.go
	./bucket-sort/go/bucket.o

bucket_rs:
	rustc -o ./bucket-sort/rust/bucket.o ./bucket-sort/rust/bucket.rs
	./bucket-sort/rust/bucket.o

bucket: bucket_c bucket_go bucket_rs

counting_c:
	gcc -o ./counting-sort/c/counting.o ./counting-sort/c/counting.c
	./counting-sort/c/counting.o

counting_go:
	go build -o ./counting-sort/go/counting.o ./counting-sort/go/counting.go
	./counting-sort/go/counting.o

counting_rs:
	rustc -o ./counting-sort/rust/counting.o ./counting-sort/rust/counting.rs
	./counting-sort/rust/counting.o

counting: counting_c counting_go counting_rs

radix_c:
	gcc -o ./radix-sort/c/radix.o ./radix-sort/c/radix.c
	./radix-sort/c/radix.o

radix_go:
	go build -o ./radix-sort/go/radix.o ./radix-sort/go/radix.go
	./radix-sort/go/radix.o

radix_rs:
	rustc -o ./radix-sort/rust/radix.o ./radix-sort/rust/radix.rs
	./radix-sort/rust/radix.o

radix: radix_c radix_go radix_rs

c: bucket_c counting_c  radix_c

go: bucket_go counting_go  radix_go

rs: bucket_rs counting_rs  radix_rs

.PHONY : generate bucket_c bucket_go bucket_rs bucket counting_c counting_go counting_rs counting
