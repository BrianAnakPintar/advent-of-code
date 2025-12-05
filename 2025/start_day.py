import os
import sys

def create_day(day_num):
    day_str = f"day{day_num:02d}"
    base_path = os.path.join(day_str)
    
    os.makedirs(base_path, exist_ok=True)
    
    # Create __init__.py
    open(os.path.join(base_path, "__init__.py"), 'a').close()
    
    # Create input.txt
    open(os.path.join(base_path, "input.txt"), 'a').close()
    
    # Create solution.py with template
    solution_path = os.path.join(base_path, "main.py")
    open(solution_path, 'a').close()

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python start_day.py <day_number>")
    else:
        create_day(int(sys.argv[1]))
