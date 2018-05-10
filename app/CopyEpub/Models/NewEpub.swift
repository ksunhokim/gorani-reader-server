//
//  NewEpub.swift
//  copyEpub
//
//  Created by sunho on 2018/05/03.
//  Copyright © 2018 sunho. All rights reserved.
//

import Foundation
import UIKit
import FolioReaderKit

class NewEpub: Epub {
    let tempURL: ManagedEpubURL
    
    init(epub: URL) throws {
        let tempURL = ManagedEpubURL(epub: epub)
        guard tempURL.isNew() else {
            tempURL.keep = true
            throw ShareError.notNew
        }
        self.tempURL = tempURL
        
        super.init()
        self.book = try FREpubParser().readEpub(epubPath: epub.path, removeEpub: false, unzipPath: booksDir.path)
        try self.parse()
    }
    
    func calculateKnownWordRate() -> Int {
        var counts = 0
        var wordSet = Set<String>()
        if let resources = self.book?.resources.resources {
            for (_, resource) in resources{
                if resource.mediaType == .xhtml {
                    if let html = try? String(contentsOfFile: resource.fullHref) {
                        let words = KnownWord.getWordsFromHTML(html: html)
                        for word in words {
                                wordSet.insert(word)
                        }
                    }
                }
            }
        }
        for word in wordSet {
            if UserData.shared.getKnownWord(word: word) != nil {
                counts += 1
            }
        }
        return counts
    }

}
