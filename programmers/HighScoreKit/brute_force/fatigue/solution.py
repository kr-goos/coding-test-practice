from itertools import permutations

def solution(k, dungeons):
    answer = -1
    for tup in permutations(dungeons,len(dungeons)):
        fatigue = k
        count = 0
        for required, consumed in tup:
            if fatigue < required:
                break
            fatigue -= consumed
            count += 1
        answer = max(answer, count)

    return answer

dungeons = [[80,20],[30,10],[50,40]]
solution(80,dungeons)