//
//  ManagedURL.swift
//  copyEpub
//
//  Created by sunho on 2018/05/03.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation

class ManagedEpubURL {
    let contentURL: URL
    
    var keep: Bool = false
    var path: String {
        return contentURL.path
    }

    init(epub: URL) {
        self.contentURL = booksDir.appendingPathComponent(epub.lastPathComponent)
    }
    
    func isNew() -> Bool {
        return !fileManager.fileExists(atPath: self.contentURL.path)
    }
    
    deinit {
        if !self.keep {
            DispatchQueue.global(qos: .utility).async { [contentURL = self.contentURL] in
                try? fileManager.removeItem(at: contentURL)
            }
        }
    }
}
