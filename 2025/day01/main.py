import utils.parser as parser

# Parses an operation of the form LXX or RXX where XX are numbers.
# Returns a tuple {Direction, Number}
def parse_operation(operation: str) -> tuple[str, int]:
    if len(operation) == 0:
        raise Exception("invalid operation: ", operation)
    direction = operation[0]
    val = int(operation[1:])
    
    if (direction not in ["L", "R"]) or (val < 0):
        raise Exception("invalid operation: ", operation)
    return (direction, val)

def rotate(pos, direction, value):
    if direction == 'L':
        return (pos - value) % 100
    elif direction == 'R':
        return (pos + value) % 100
    else:
        raise ValueError("Invalid direction")

# Rotate for Part 2.
def rotate2(pos: int, direction: str, value: int, res: list[int]):
    num_rotations = value//100
    new_value = value%100
    res[0] += num_rotations

    new_pos = 0
    if direction == 'L':
        new_pos = pos - new_value
        if pos <= new_value and pos > 0:
            res[0] += 1 
    elif direction == 'R':
        new_pos = pos + new_value
        if new_pos >= 100:
            res[0] += 1

    return new_pos % 100

def run():
    print("Day 1, Hooray")

    print("Part 1")

    dial_pos = 50
    num_zeroes = 0

    operations = parser.file_to_string_list("day01/input.txt")
    for operation in operations:
        direction, value = parse_operation(operation)
        dial_pos = rotate(dial_pos, direction, value)

        if dial_pos == 0:
            num_zeroes += 1

    print(num_zeroes)

    print("Part 2")
    # In part 2, we count the number of times it always reaches 0, despite it being mid operation

    dial_pos = 50
    res = [0]
    for operation in operations:
        direction, value = parse_operation(operation)
        dial_pos = rotate2(dial_pos, direction, value, res)

    print(res[0])

run()

