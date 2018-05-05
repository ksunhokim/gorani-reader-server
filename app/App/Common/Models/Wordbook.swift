//
//  Wordbook.swift
//  app
//
//  Created by Sunho Kim on 05/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation

class Wordbook {
    var id: Int
    var name: String
    var entries: [WordbookEntry]
    
    init(id: Int, name: String) {
        self.id = id
        self.name = name
    }
}

class WordbookEntry {
    
}
