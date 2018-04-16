import re
import os

with open('./result/mispatching_parentheses.txt', 'w') as out:
  for dirpath, dnames, fnames in os.walk('./raw'):
      for f in fnames:
        if f.endswith('.dic'):
          print(f, file=out)
          with open('./raw/' + f, 'r') as file:
            line = file.readline()
            while line:
              if line.count('(') != line.count(')'):
                print(line, file=out)
              line=file.readline()