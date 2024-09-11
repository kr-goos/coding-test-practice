import collections

# 정확성: 58.3
# 효율성: 41.7
# 합계: 100.0 / 100.0
def solution1(participant, completion):
    pd = collections.Counter(participant)
    pc = collections.Counter(completion)

    for k,v in pd.items():
        if v > pc[k]:
            return k
    return ''

# 정확성: 58.3
# 효율성: 41.7
# 합계: 100.0 / 100.0
def solution2(participant, completion):
    hash_table = {}
    total = 0
    
    for p in participant:
        hash_value = hash(p)
        hash_table[hash_value] = p
        total += hash_value

    for c in completion:
        total -= hash(c)

    return hash_table[total]


solution2(["mislav", "stanko", "mislav", "ana"],["stanko", "ana", "mislav"])