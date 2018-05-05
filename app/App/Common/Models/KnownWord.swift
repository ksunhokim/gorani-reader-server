//
//  KnownWord.swift
//  app
//
//  Created by Sunho Kim on 05/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation
import SQLite

fileprivate let table = Table("known_words")
fileprivate let wordField = Expression<String>("word")

class KnownWord {
    var word: String

    init(word: String) {
        self.word = word
    }

    func add(_ connection: Connection) throws {
        try connection.run(table.insert( wordField <- self.word ))
    }
    
    func delete(_ connection: Connection) throws {
        let me = table.where(wordField == self.word)
        try connection.run(me.delete())
    }


    static func get(_ connection: Connection, word: String) -> KnownWord? {
        let query = table.where(wordField == word)
        do {
            if let known = try connection.pluck(query) {
                return KnownWord(word: try known.get(wordField))
            }
        } catch {}
        
        return nil
    }

    static func prepare(_ conenction: Connection) throws {
        try conenction.run(table.create(ifNotExists: true) { t in
            t.column(wordField, unique: true)
        })
    }
}
