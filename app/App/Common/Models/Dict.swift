//
//  db.swift
//  app
//
//  Created by Sunho Kim on 05/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation
import SQLite



class Dict {
    fileprivate static let dictURL: URL = {
        let path = Bundle.main.path(forResource: "dict", ofType: "db")!
        return URL(string: path)!
    }()
    
    static let shared = Dict(url: dictURL)
    
    let connection: Connection
    
    typealias EntrySortPolicy = (_ word: String, _ entries: [DictDefinition], _ pos: POS?) -> [DictDefinition]
    var entrySortPolicy: EntrySortPolicy?

    private init(url: URL) {
        self.connection = try! Connection(url.path)
    }
    
    func get(word: String, pos: POS? = nil) -> DictEntry? {
        return DictEntry.get(self.connection, word: word, pos: pos, policy: self.entrySortPolicy)
    }
    
    func search(word: String, pos: POS? = nil, type: VerbType? = nil) -> [DictEntry] {
        return DictEntry.search(connection, word: word, pos: pos, type: type, policy: self.entrySortPolicy)
    }
}
