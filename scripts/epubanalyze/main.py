from os import listdir
from bs4 import BeautifulSoup
import zipfile
import operator

files = ['./epubs/'+file for file in listdir('./epubs') if file.endswith('.epub')]
print('These files will be analyzed:')
print(*files, sep='\n')

def htmltexts(epubfile):
  zf = zipfile.ZipFile(epubfile)
  files = [info.filename for info in zf.infolist() if info.filename.endswith('.html')]
  return [str(zf.read(file), 'utf-8') for file in files]

counts = dict()
for file in files:
  for text in htmltexts(file):
    soup = BeautifulSoup(text, 'html.parser')
    for tag in soup.findAll():
      text = tag.find(text=True,recursive=False)
      if text is not None:
        if len(text) > 10:
          counts[tag.name.lower()] = counts.get(tag.name, 0) + 1

print('(tag, number of appearance)')
sorted_counts = sorted(counts.items(), key=operator.itemgetter(1), reverse=True)
print(*sorted_counts, sep='\n')
