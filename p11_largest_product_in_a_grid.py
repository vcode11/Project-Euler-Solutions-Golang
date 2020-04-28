with open('p10.txt') as f:
    l = f.read()
    l = l.split('\n')
    l = [i.split(' ') for i in l]
    l = l[:-1]
    for i in range(len(l)):
        l[i] = list(map(int,l[i]))
    prod = 1
    for i in range(len(l)):
        for j in range(len(l[i])):
            if i+3 < len(l):
                prod = max(prod, l[i][j]*l[i+1][j]*l[i+2][j]*l[i+3][j])
            if j+3 < len(l[i]):
                prod = max(prod, l[i][j]*l[i][j+1]*l[i][j+2]*l[i][j+3])
            if i+3 < len(l) and j+3 < len(l[i]):
                prod = max(prod, l[i][j]*l[i+1][j+1]*l[i+2][j+2]*l[i+3][j+3])
            if i+3 < len(l) and j >=3:
                prod = max(prod, l[i][j]*l[i+1][j-1]*l[i+2][j-2]*l[i+3][j-3])
    print(prod) 
            
