//
//  File.swift
//  copyEpub
//
//  Created by sunho on 2018/05/03.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation

let fileManager = FileManager.default
let sharedDir: URL = fileManager.containerURL(forSecurityApplicationGroupIdentifier: "group.sunho.app")!
let booksDir: URL = {
    let url = sharedDir.appendingPathComponent("books")
    if !fileManager.fileExists(atPath: url.path) {
       try! fileManager.createDirectory(atPath: url.path, withIntermediateDirectories: true, attributes: nil)
    }
    return url
}()

