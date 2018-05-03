//
//  Epub.swift
//  app
//
//  Created by sunho on 2018/05/03.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation

internal func contentsOfDirectory(path: String) -> [String]? {
    guard let paths = try? FileManager.default.contentsOfDirectory(atPath: path) else { return nil}
    return paths.map { content in (path as NSString).appendingPathComponent(content)}
}

internal let sharedDir = FileManager.default.containerURL(forSecurityApplicationGroupIdentifier: "group.sunho.app")!
internal let sharedBooks = sharedDir.appendingPathComponent("books")
