import utils.parser as parser

def parse_range(s: str):
    """
    Takes a string of form "xx-yy" where xx and yy are numbers 
    Returns: A tuple {xx, yy} describing the range
    """

    vals = s.split('-')
    return (int(vals[0]), int(vals[1]))

def is_invalid(id: str):
    if len(id) % 2 != 0:
        return False

    mid_idx = len(id)//2
    for i in range(mid_idx):
        if id[i] != id[i+mid_idx]:
            return False
    return True

def get_invalid_ids(lower: int, upper: int, res: list[int]):
    for idd in range(lower, upper+1):
        if is_invalid(str(idd)):
            res[0] += idd


def is_duplicate_of_size(id, size):
    if len(id) % size != 0:
        return False

    num_sizes = len(id)//size
    for outer in range(1, num_sizes):
        for i in range(size):
            if id[i] != id[i+size*outer]:
                return False
    return True

def is_invalid2(idd: str):
    mid_idx = len(idd)//2 

    for size in range(1, mid_idx+1):
        if is_duplicate_of_size(idd, size):
            return True
    return False
    

def get_invalid_ids2(lower: int, upper: int, res: list[int]):
    for idd in range(lower, upper+1):
        if is_invalid2(str(idd)):
            res[0] += idd


def run():
    line: str = parser.file_to_string_list("day02/input.txt")[0]
    arrs: list[str] = line.split(',')

    print("Part 1")
    sum_invalid_ids = [0]
    for s in arrs:
        lower_bound, upper_bound = parse_range(s)
        get_invalid_ids(lower_bound, upper_bound, sum_invalid_ids)
    print(sum_invalid_ids[0])

    print("Part 2")
    sum_invalid_ids = [0]
    for s in arrs:
        lower_bound, upper_bound = parse_range(s)
        get_invalid_ids2(lower_bound, upper_bound, sum_invalid_ids)
    print(sum_invalid_ids[0])

run()
