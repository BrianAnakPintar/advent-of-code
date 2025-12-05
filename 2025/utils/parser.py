def file_to_string_list(path: str):
    """
    Takes a path and returns a list of string representing each line.
    """
    res = []
    with open(path, 'r') as f:
        for line in f:
            res.append(line)
    return res

