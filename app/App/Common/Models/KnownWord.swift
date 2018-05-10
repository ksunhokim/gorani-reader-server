//
//  KnownWord.swift
//  app
//
//  Created by Sunho Kim on 05/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation
import SQLite
import SwiftSoup
import Regex

fileprivate let table = Table("known_words")
fileprivate let wordField = Expression<String>("word")
fileprivate let specialCharPattern = "[^1-9a-zA-Z-\\s]"

class KnownWord {
    var word: String

    init(word: String) {
        self.word = word
    }
    

    func add(_ connection: Connection) throws {
        do {
            try connection.run(table.insert( wordField <- self.word ))
        } catch let Result.error(_, code, _) where code == SQLITE_CONSTRAINT {
        }
    }
    
    func delete(_ connection: Connection) throws {
        let me = table.where(wordField == self.word)
        try connection.run(me.delete())
    }
    
    static func getWordsFromHTML(html: String) -> [String] {
        var arr: [String] = []
        do {
            let doc = try SwiftSoup.parse(html)
            let ps = try doc.select("p")
            try ps.select("chunk").remove()
            for ele in ps.array() {
                let text = try ele.text()
                let replaced = specialCharPattern.r?.replaceAll(in: text, with: "")
                if let words = replaced?.components(separatedBy: " ") {
                    arr += words
                }
            }
        } catch {}
        
        return arr
    }

    static func add(_ connection: Connection, html: String) throws {
        let words = getWordsFromHTML(html: html)
        let nWords = words.map { word in
            return KnownWord(word: word)
        }
        for word in nWords {
            try word.add(connection)
        }
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
