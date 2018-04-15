import re
import os

pos_pat = re.compile(r'\s([a-z-]+)\,')
pos_candidates = ['s', 'd', 'auxil', 'dn', 'q', 'b', 'c', 'm', 'pt', 'pi', 'rel', 'pron', 'xprep', 'xconj', 'xbent', 't', 'vn', 'xa', '-va', 'as', 'inits', 'avt', 'nad', 'vy', 'xvt'] 

with open('./result/pos_candidate_sentences.txt', 'w') as out:
  for dirpath, dnames, fnames in os.walk('./raw'):
      for f in fnames:
        if f.endswith('.dic'):
          print(f, file=out)
          with open('./raw/' + f, 'r') as file:
            line = file.readline()
            while line:
              line = re.sub(r'^.*:','', line)
              part = line[:15]
              poss = pos_pat.findall(part)
              for pos in poss:
                if pos in pos_candidates:
                  print(line, file=out)
              line=file.readline()