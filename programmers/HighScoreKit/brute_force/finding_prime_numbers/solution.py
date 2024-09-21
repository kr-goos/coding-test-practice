import math
from itertools import permutations

def solution(numbers):
    number_set = set()

    for i in range(1, len(numbers) + 1):
        for perm in permutations(numbers, i):
            number = int(''.join(perm))
            number_set.add(number)
        
    prime_count = sum(1 for num in number_set if is_prime(num))

    return prime_count

def is_prime(num):
    if num < 2:
        return False
    
    for i in range(2,int(math.sqrt(num)) + 1):
        if num % i == 0:
            return False
    return True

numbers = "011"
result = 2

if result == solution(numbers):
    print("success")
else:
    print("failure")