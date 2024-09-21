def solution(citations):

    citations.sort(reverse=True)

    for i, citation in enumerate(citations):
        if citation < i + 1:
            return i

    return len(citations)

testcase = [
    [3, 0, 6, 1, 5],
    [10, 8, 5, 4, 3],
    [25, 8, 5, 3, 3],
]

testsuccess = (3,4,3)

for i, tc in enumerate(testcase):
    result = solution(tc)
    print(f"test[{i+1}] : ", end='')
    if result == testsuccess[i]:
        print("success")
    else:
        print("failure")



