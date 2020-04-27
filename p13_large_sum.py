with open('p13.txt') as f:
    numbers = f.read().split('\n')[:-1]
    numbers = list(map(int, numbers))
    print(str(sum(numbers))[:10])
