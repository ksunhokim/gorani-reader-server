//
//  DataBase.swift
//  app
//
//  Created by Sunho Kim on 05/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation
import SQLite

class UserData {
    var connection: Connection
    
    static let shared = try! UserData(url: FileUtill.userDataURL)
    
    private init(url: URL) throws {
        self.connection = try! Connection(url.path)
        try KnownWord.prepare(self.connection)
    }
  
    func getKnownWord(word: String) -> KnownWord? {
        return KnownWord.get(self.connection, word: word)
    }
    
    func addKnownWord(word: KnownWord) throws {
        try word.add(self.connection)
    }
    
    func addKnownWords(html: String) throws {
        try KnownWord.add(self.connection, html: html)
    }
    
    func deleteKnownWord(word: KnownWord) throws {
        try word.delete(self.connection)
    }
}
