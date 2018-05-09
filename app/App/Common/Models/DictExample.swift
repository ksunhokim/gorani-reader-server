//
//  DictExample.swift
//  app
//
//  Created by Sunho Kim on 09/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation
import SQLite

fileprivate let examplesTable = Table("examples")
fileprivate let defIdField = Expression<Int>("def_id")
fileprivate let firstField = Expression<String>("first")
fileprivate let secondField = Expression<String>("second")

class DictExample {
    var first: String
    var second: String
    
    init(first: String, second: String) {
        self.first = first
        self.second = second
    }

    static func fetch(_ connection: Connection, def: DictDefinition) {
        let query = examplesTable.where(defIdField == def.id)
        if let results = try? connection.prepare(query) {
            for result in results {
                do {
                    def.examples.append(DictExample(first: try result.get(firstField), second: try result.get(secondField)))
                } catch {}
            }
        }
    }
}
