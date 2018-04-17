import json
import codecs

prons = dict()
with codecs.open('./raw/cmudict-0.7b', encoding='utf-8', errors='ignore') as file:
  line = file.readline()
  while line:
    line = line[:-1] # \n
    arr = line.split('  ')
    if len(arr)==2:
      prons[arr[0].lower()]=arr[1]
    line = file.readline()

missed = 0
missed_list = []
dictionary = dict()
with open('./raw/crawled.json', encoding='utf8') as file:
  dictionary = json.load(file)
  for word, v  in dictionary.items():
    if word in prons:
      dictionary[word]['pron'] = prons[word]
    else:
      missed +=1
      missed_list.append(word)

with open('./result/missed.txt', 'w') as file:
  print(missed_list, sep='\n', file=file)

with open('./result/proned_crawled.json', 'w', encoding='utf8') as file:
  json.dump(dictionary, file, ensure_ascii=False)
