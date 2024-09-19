def solution(numbers):
    strNums = sorted(list(map(str, numbers)),key=lambda x: x*3, reverse=True)
    if strNums[0] == '0':
        return '0'
    return ''.join(strNums)


numbers	= [
    [6, 10, 2],
    [3, 30, 34, 5, 9],
]
result = [
    "6210",
	"9534330",
]

print("test","success"if [solution(num) for num in numbers] == result else "failure")
