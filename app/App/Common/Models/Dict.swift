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
    let connection: Connection
    
    let words = Table("words")
    
    init(connection: Connection) {
        self.connection = connection
    }
}
