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
