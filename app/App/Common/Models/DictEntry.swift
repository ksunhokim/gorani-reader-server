//
//  DictEntry.swift
//  app
//
//  Created by Sunho Kim on 08/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation

class DictEntry {
    var id: Int
    var word: String
    var pron: String
    var defs: [DictDefinition] = []
    
    init(id: Int, word: String, pron: String) {
        self.id = id
        self.word = word
        self.pron = pron.unstressed
    }
}

class DictEntryRedirect: DictEntry {
    var verbType: VerbType
    
    init(id: Int, word: String, pron: String, type: VerbType) {
        self.verbType = type
        super.init(id: id, word: word, pron: pron)
    }
}

class DictDefinition {
    var id: Int
    var pos: POS?
    var def: String
    var examples: [DictExample] = []
    
    init(id: Int, word: DictEntry, pos: POS?, def: String) {
        self.id = id
        self.pos = pos
        self.def = def
    }
}

class DictExample {
    var first: String
    var second: String
    
    init(first: String, second: String) {
        self.first = first
        self.second = second
    }
}
