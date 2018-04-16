import re
import os

with open('./result/colons.txt', 'w') as out:
  for dirpath, dnames, fnames in os.walk('./raw'):
      for f in fnames:
        if f.endswith('.dic'):
          with open('./raw/' + f, 'r') as file:
            line = file.readline()
            while line:
              arr = line.split(':')
              if len(arr) != 2:
                print(line, file=out)
              line=file.readline()
