import re
import os
import codecs

prons = dict()
with codecs.open('./raw/cmudict-0.7b', encoding='utf-8', errors='ignore') as file:
  line = file.readline()
  while line:
    arr = line.split('  ')
    if len(arr)==2:
      prons[arr[0].lower()]=arr[1]
    line = file.readline()

for dirpath, dnames, fnames in os.walk('./raw'):
    for f in fnames:
      if f.endswith('.dic'):
        with open('./raw/' + f, 'r') as file:
          line = file.readline()
          while line:
            if len(line.split(':'))!=2:
              print(line)
            line = file.readline()
          
            