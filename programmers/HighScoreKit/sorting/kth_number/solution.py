
# 채점 결과
# 정확성: 100.0
# 합계: 100.0 / 100.0
def solution(array, commands):
    answer = []
    for cmd in commands:
        answer.append(sorted(array[cmd[0]-1:cmd[1]])[cmd[2]-1]) 
    
    return answer

array = [1, 5, 2, 6, 3, 7, 4]
commands = [[2, 5, 3], [4, 4, 1], [1, 7, 3]]
result = [5, 6, 3]


print("test", "success" if solution(array,commands) == result else "failure")

# # map 을 이용한 방법
# def solution2(array, commands):
#     return list(map(lambda cmd:sorted(array[cmd[0]-1:cmd[1]])[cmd[2]-1], commands))

# # comprehension
# def solution3(array, commands):
#     return [sorted(array[a[0]-1:a[1]])[a[2]-1] for a in commands]