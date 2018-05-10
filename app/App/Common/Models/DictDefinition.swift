//
//  DictDefinition.swift
//  app
//
//  Created by Sunho Kim on 09/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation
import SQLite

fileprivate let defsTable = Table("defs")
fileprivate let idField = Expression<Int>("id")
fileprivate let wordIdField = Expression<Int>("word_id")
fileprivate let posField = Expression<String?>("pos")
fileprivate let defField = Expression<String>("def")

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
    
    class func fetch(_ connection: Connection, entry: DictEntry, pos pos2: POS?, policy: Dict.EntrySortPolicy?) {
        let query = defsTable.where(wordIdField == entry.id)
            .order(posField, idField)
        
        var defs: [DictDefinition] = []
        guard let results = try? connection.prepare(query) else {
            return
        }
        for result in results {
            do {
                let defi = DictDefinition(id: try result.get(idField), word: entry, pos: POS(rawValue: try result.get(posField) ?? ""), def: try result.get(defField))
                DictExample.fetch(connection, def: defi)
                if pos2 != nil && pos2 == defi.pos {
                    defs.insert(defi, at: 0)
                } else {
                    defs.append(defi)
                }
            } catch{}
        }
        if let policy = policy {
            defs = policy(entry.word, defs, pos2)
        }
        entry.defs = defs
    }
}
