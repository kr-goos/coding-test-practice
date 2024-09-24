# node : 시작노드
# graph : 인덱스와 연결되어있는 인덱스의 리스트를 저장하는 리스트
# visited : 전체 노드개수보다 1 큰 리스트 (인덱스 1 부터 n + 1 까지 사용)
# 예 ) n 이 10 이면 0 ~ 10 까지 생성하여 1번부터 10번 송전탑으로 이용

def dfs(node, graph, visited):
    visited[node] = True
    count = 1
    for neighbor in graph[node]:
        if not visited[neighbor]:
            count += dfs(neighbor, graph, visited)
    return count

def solution(n, wires):
    
    graph = [[] for _ in range(n + 1)]
    for v1, v2 in wires:
        graph[v1].append(v2)
        graph[v2].append(v1)

    min_difference = float('inf')

    for v1, v2 in wires:
        # 전선 끊기
        graph[v1].remove(v2)
        graph[v2].remove(v1)

        # 첫 번째 전력망의 크기 구하기
        visited = [False] * (n + 1)
        power_1_size = dfs(1, graph, visited)

        power_2_size = n - power_1_size

        # 두 전력망의 차이 계산
        min_difference = min(min_difference, abs(power_1_size - power_2_size))

        # 전선 복구
        graph[v1].append(v2)
        graph[v2].append(v1)
    
    return min_difference


n = [9, 4, 7]
wires = [
    [[1,3],[2,3],[3,4],[4,5],[4,6],[4,7],[7,8],[7,9]],
    [[1,2],[2,3],[3,4]],
    [[1,2],[2,7],[3,7],[3,4],[4,5],[6,7]],
]
result = [3,0,1]

for i in range(3):
    print(f"testcase {i+1} : ", end='')
    r = solution(n[i], wires[i])
    if r == result[i]:
        print("success")
    else:
        print("failure")
        print("===============")
        print("solution result : ", r)
        print("result : ", result[i])
        print("===============")
