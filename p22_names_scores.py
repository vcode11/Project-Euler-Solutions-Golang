with open("names.txt") as f:
    names = f.read().split(',')
    names = [name.replace('"','') for name in names]
    names = [name.replace('\n','') for name in names]
    score = 0
    names.sort()
    print(names[937])
    for i,name in enumerate(names):
        for char in name:
            score+=(i+1)*(ord(char)-64)
    print(score)

