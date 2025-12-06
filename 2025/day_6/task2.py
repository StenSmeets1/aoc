def get_sign(problems: list[str], x: int):
    while x > 0 and x < len(problems[0]) and problems[-1][x] == " ":
        x -= 1
    return problems[-1][x]


def solve(problems: list[str]) -> int:
    total = 0
    curr_sign = get_sign(problems, len(problems[0]) - 1)
    curr_res = 1 if curr_sign == "*" else 0

    for x in reversed(range(0, len(problems[0]))):
        curr_num = ""
        for y in range(0, len(problems) - 1):
            if problems[y][x].isdigit():
                curr_num += problems[y][x]

        if not curr_num:
            total += curr_res
            curr_sign = get_sign(problems, x)
            curr_res = 1 if curr_sign == "*" else 0
            continue

        if curr_sign == "+":
            curr_res += int(curr_num)
        elif curr_sign == "*":
            curr_res *= int(curr_num)

        if x == 0:
            total += curr_res

    return total


with open('input.txt', 'r') as f:
    input_text = f.read()

lines = input_text.strip('\n').split('\n')

answer = solve(lines)
print(f"Grand total: {answer}")
