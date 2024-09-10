from typing import *

# 정확성: 83.3
# 효율성: 16.7
# 합계: 100.0 / 100.0
def solution1(phone_book: List[int]) -> bool:
    phone_book.sort()
    for i in range(len(phone_book)-1):
        if phone_book[i] == phone_book[i+1][:len(phone_book[i])]:
            return False
    return True

# 채점 결과
# 정확성: 83.3
# 효율성: 16.7
# 합계: 100.0 / 100.0
def solution2(phone_book: List[int]) -> bool:
    hash_map: Dict[str:int] = {k: 0 for k in phone_book}

    for phone_number in phone_book:
        prefix = ''
        for digit in phone_number:
            prefix += digit
            if prefix == phone_number:
                continue

            if prefix in hash_map:
                return False
            
    return True

print(solution1(["119", "1195524421", "97674223"]))
print(solution2(["119", "1195524421", "97674223"]))

