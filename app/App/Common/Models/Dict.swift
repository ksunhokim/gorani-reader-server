//
//  db.swift
//  app
//
//  Created by Sunho Kim on 05/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation
import SQLite

fileprivate let wordsTable = Table("words")
fileprivate let id = Expression<Int>("id")
fileprivate let word = Expression<String>("word")
fileprivate let pron = Expression<String?>("pron")

fileprivate let defsTable = Table("defs")
fileprivate let word_id = Expression<Int>("word_id")
fileprivate let pos = Expression<String?>("pos")
fileprivate let def = Expression<String>("def")

fileprivate let examplesTable = Table("examples")
fileprivate let def_id = Expression<Int>("def_id")
fileprivate let first = Expression<String>("first")
fileprivate let second = Expression<String>("second")


class Dict {
    typealias EntrySortPolicy = (_ word: String, _ entries: [DictDefinition], _ pos: POS?) -> [DictDefinition]
    let connection: Connection
    
    var entrySortPolicy: EntrySortPolicy?
    
    init(connection: Connection) {
        self.connection = connection
    }
    
    func get(word wordstr: String, pos: POS? = nil) -> DictEntry? {
        let query = wordsTable.where(word == wordstr)
        do {
            if let entry = try self.connection.pluck(query) {
                let entry = DictEntry(id: try entry.get(id), word: try entry.get(word), pron: try entry.get(pron) ?? "")
                fecthDefs(entry, pos)
                return entry
            }
        } catch {}
        
        return nil
    }

    fileprivate func fecthDefs(_ entry: DictEntry, _ pos2: POS?) {
        let query = defsTable.where(word_id == entry.id)
                             .order(pos, id)
        
        var defs: [DictDefinition] = []
        if let results = try? self.connection.prepare(query) {
            for result in results {
                do {
                    let defi = DictDefinition(id: try result.get(id), word: entry, pos: POS(rawValue: try result.get(pos) ?? ""), def: try result.get(def))
                    self.fetchExamples(defi)
                    defs.append(defi)
                } catch{}
            }
        }
        if let entrySortPolicy = entrySortPolicy {
            defs = entrySortPolicy(entry.word, defs, pos2)
        }
        entry.defs = defs
    }

    fileprivate func fetchExamples(_ def: DictDefinition) {
        let query = examplesTable.where(def_id == def.id)
        if let results = try? self.connection.prepare(query) {
            for result in results {
                do {
                    def.examples.append(DictExample(first: try result.get(first), second: try result.get(second)))
                } catch {}
            }
        }
    }

    func search(word: String, pos: POS? = nil, type: VerbType? = nil) -> [DictEntry] {
        if word == "" {
            return []
        }
        
        var entries: [DictEntry] = []
        let candidates = VerbType.candidates(word: word)
        for candidate in candidates {
            if let entry = self.get(word: candidate.0, pos: pos) {
                let entry = DictEntryRedirect(entry: entry, type: candidate.1)
                if candidate.1 == type {
                    entries.insert(entry, at: 0)
                } else {
                    entries.append(entry)
                }
            }
        }
        if let entry = self.get(word: word, pos: pos) {
            entries.append(entry)
        }
        
        return entries
    }
    
    
}
