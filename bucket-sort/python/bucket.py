def bucketSort(array):
    largest = max(array)
    length = len(array)
    size = largest/length

    buckets = [[] for i in range(length)]

    for i in range(length):
        index = int(array[i]/size)
        if index != length:
            buckets[index].append(array[i])
        else:
            buckets[length - 1].append(array[i])

    for i in range(len(array)):
        buckets[i] = sorted(buckets[i])

    result = []
    for i in range(length):
        result = result + buckets[i]
            
    return result


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
            
print("\nIniciando Execucao do Bucket Sort em Python: \n")
bucketSort(numbers)

for i in range(len(numbers)):
    print(numbers[i],end=" ")
print("\n")

