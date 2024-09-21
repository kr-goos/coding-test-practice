def solution(answers):
    pattern1 = [1,2,3,4,5]
    pattern2 = [2,1,2,3,2,4,2,5]
    pattern3 = [3,3,1,1,2,2,4,4,5,5]
    l = [0,0,0]
    for i,v in enumerate(answers):
        if pattern1[i % len(pattern1)] == v:
            l[0] +=1

        if pattern2[i % len(pattern2)] == v:
            l[1] +=1

        if pattern3[i % len(pattern3)] == v:
            l[2] +=1
    
    maxv = max(l)

    answer = []
    for i, v in enumerate(l):
        if v == maxv:
            answer.append(i+1)            
    return answer