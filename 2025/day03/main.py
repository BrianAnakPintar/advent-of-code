import utils.parser as parser

def get_best_joltage(bank: str):
    n = len(bank)
    best_val = 0
    # Naive algorithm, I might be able to speed it up w/ DP
    for i in range(n):
        for j in range(i+1, n):
            val_str = bank[i] + bank[j]
            best_val = max(int(val_str), best_val)

    return best_val
        

def get_best_joltage2(bank: str, length: int):
    # Greedy Trick, keep taking the largest value as long as
    # there are still numbers left
    best = ""
    last_best_idx = -1
    # Try to find the best value for each index.
    for i in range(length):
        # Put simply, find the best value for index i s.t. we still have numbers rem
        end_idx = len(bank) - length + i
        best_idx = last_best_idx+1
        for j in range(last_best_idx+1, end_idx):
            if bank[j] > bank[best_idx]:
                best_idx = j
        last_best_idx = best_idx
        best += bank[last_best_idx]
    return int(best)

    

def run():
    banks = parser.file_to_string_list("day03/input.txt")

    print("Part 1")
    res = 0
    for bank in banks:
        res += get_best_joltage(bank)
    print(res)

    print("Part 2")
    res = 0
    for bank in banks:
        res += get_best_joltage2(bank, 12)
    print(res)

run()
