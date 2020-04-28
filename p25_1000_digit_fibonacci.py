a = 0
b = 1
i = 1
while True:
    if len(str(b)) == 1000:
        print(i)
        break
    a, b = b, a+b
    i+=1


