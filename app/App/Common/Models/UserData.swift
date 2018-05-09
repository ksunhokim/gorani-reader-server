//
//  DataBase.swift
//  app
//
//  Created by Sunho Kim on 05/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation
import SQLite


fileprivate var userData_: Connection? = nil

var userData: Connection {
    if userData_ == nil {
        userData_ = try! Connection(userDataURL.path)
        try! KnownWord.prepare(userData_!)
    }
    
    return userData_!
}

class UserData {
    var connection: Connection
    
    static let shared = try! UserData(url: userDataURL)
    
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
    
    func deleteKnownWord(word: KnownWord) throws {
        try word.delete(self.connection)
    }
}
