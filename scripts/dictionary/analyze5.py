import re
import os

pos_pat = re.compile(r'\s([a-z-]+)\,')
sentences = dict()

with open('./result/empty_defs.txt', 'w') as out:
  for dirpath, dnames, fnames in os.walk('./raw'):
      for f in fnames:
        if f.endswith('.dic'):
          with open('./raw/' + f, 'r') as file:
            line = file.readline()
            while line:
              line2 = re.sub(r'^.*:','', line)
              part = line2[:15]
              if len(part) <= 5:
                print(line, file=out)
              line=file.readline()

  for pos, sentence in sentences.items():
    for sen in sentence:
      print(sen, file=out)
    print('-------------------',file=out)