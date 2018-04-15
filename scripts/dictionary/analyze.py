import re
import os

pos_pat = re.compile(r'\s([a-z-]+)\,')
pos_count = dict()
for dirpath, dnames, fnames in os.walk('./raw'):
    for f in fnames:
      if f.endswith('.dic'):
        with open('./raw/' + f, 'r') as file:
          line = file.readline()
          while line:
            line = re.sub(r'^.*:','', line)
            part = line[:15]
            #pos
            poss = pos_pat.findall(part)
            for pos in poss:
              if pos not in pos_count:
                pos_count[pos] = 1
              else:
                pos_count[pos] += 1
            
            line=file.readline()

pos_count = sorted(pos_count.items(), key=lambda x: x[1])

with open('./result/pos_count.txt', 'w') as file:
  for count in pos_count:
    print(count, file=file)