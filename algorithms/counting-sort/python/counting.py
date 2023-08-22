def countingSort(array):
    size = len(array)
    output = [0] * size
    outputPosition = 0

    count = [0] * (max(array) + 1)
    
    for i in range(0, size):
        count[array[i]] += 1

    for i in range(0, len(count)):
        if count[i] != 0:
            while count[i]> 0:
                output[outputPosition] = i
                outputPosition += 1
                count[i] -= 1

    for i in range(0, size):
        array[i] = output[i]


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

print("\nIniciando Execucao do Counting Sort em Python: \n")
countingSort(numbers)

for i in range(len(numbers)):
    print(numbers[i],end=" ")
print("\n")