import sqlite3
import json

conn = sqlite3.connect('result/dict.db')
cursor = conn.cursor()
cursor.execute('''CREATE TABLE words(
    id integer PRIMARY KEY,
    word text NOT NULL,
    pron text 
)''')
cursor.execute('''CREATE TABLE defs(
    id integer PRIMARY KEY,
    word_id integer NOT NULL,
    pos text, 
    def text NOT NULL, 
    FOREIGN KEY (word_id) REFERENCES words(id)
)''')
cursor.execute('''CREATE TABLE examples(
    def_id integer NOT NULL, 
    first text NOT NULL, 
    second text NOT NULL, 
    FOREIGN KEY (def_id) REFERENCES defs(id)
)''')

with open('raw/proned_crawled.json') as f:
    data = json.load(f)

for word, value in data.items():
    cursor.execute('''INSERT INTO words 
    (word, pron) values (?, ?)''', (word, value['pron']))
    word_id = cursor.lastrowid
    if 'defs' in value:
        for defi in value['defs']:
            cursor.execute('''INSERT INTO defs
            (word_id, pos, def) values (?, ?, ?)''', (word_id, defi['pos'], defi['def']))
            def_id = cursor.lastrowid
            if 'examples' in defi:
                for example in defi['examples']:
                    cursor.execute('''INSERT INTO examples
                    (def_id, first, second) values (?, ?, ?)''', (def_id, example['first'], example['second']))
conn.commit()
    