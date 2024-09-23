def solution(brown, yellow):
    area = brown + yellow
    for h in range(1, int(area**(1/2)) + 1):
        if area % h == 0:  
            w = area // h
            if 2 * (w + h) - 4 == brown:
                return [w, h]
    return []

brown = [10,8,24]
yellow = [2,1,24]
result = [
    [4, 3],
    [3, 3],
    [8, 6],
]
for i in range(3):
    r = solution(brown[i],yellow[i])
    print(f"testcase [{i}] : ",end="")
    if r == result[i]:
        print("sucecss")
    else:
        print("failure")
        print("solution result : ", r)