def countingSort(arr, exp1):
    n = len(arr)

    output = [0] * (n)

    count = [0] * (10)

    for i in range(0, n):
        index = arr[i] // exp1
        count[index % 10] += 1

    for i in range(1, 10):
        count[i] += count[i - 1]

    i = n - 1
    while i >= 0:
        index = arr[i] // exp1
        output[count[index % 10] - 1] = arr[i]
        count[index % 10] -= 1
        i -= 1


    i = 0
    for i in range(0, len(arr)):
        arr[i] = output[i]


def radixSort(arr):
    max1 = max(arr)

    exp = 1
    while max1 / exp >= 1:
        countingSort(arr, exp)
        exp *= 10



with open("./number-generator/numbers.txt", 'r') as f:
    numbersFile = f.read()
    numberString = ""
    numbers = []
    for n in numbersFile:
        if n != ',':
            numberString = numberString + n
        else:
            numbers.append(int(numberString))
            numberString = ""
print("\nIniciando Execucao do Radix Sort em Python: \n")
radixSort(numbers)

for i in range(len(numbers)):
    print(numbers[i],end=" ")
print("\n")
