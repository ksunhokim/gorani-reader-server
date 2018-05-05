pasts = []
completes = []

with open('irregular.txt') as file:
    for line in file.readlines():
        line = line.rstrip()
        arr = line.split(' – ')
        if len(arr) != 3:
            print(line)
        
        word = arr[0]
        past = arr[1].split('/')
        complete = arr[2].split('/')
        
        for p in past:
            if p not in [verb[0] for verb in pasts]:
                pasts.append((p, word))
        
        for c in complete:
            if c not in [verb[0] for verb in completes]:
                completes.append((c, word))

with open('pasts.txt', 'w') as file:
    file.write('[')
    for v in pasts:
        if v[0] == v[1]:
            continue
        file.write('"{}":"{}",'.format(v[0], v[1]))
    file.write(']')

with open('completes.txt', 'w') as file:
    file.write('[')
    for v in completes:
        if v[0] == v[1]:
            continue
        file.write('"{}":"{}",'.format(v[0], v[1]))
    file.write(']')