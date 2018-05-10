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
        self.book = try FREpubParser().readEpub(epubPath: epub.path, removeEpub: false, unzipPath: FileUtill.booksDir.path)
        try self.parse()
    }
    

    func calculateKnownWordRate() -> Double {
        var counts = 0
        var set = Set<String>()

        guard let resources = self.book?.resources.resources else {
            return 1
        }
            
        for (_, resource) in resources {
            if resource.mediaType == .xhtml {
                if let html = try? String(contentsOfFile: resource.fullHref) {
                    KnownWord.getWordsFromHTML(set: &set, html: html)
                }
            }
        }

        for word in set {
            if UserData.shared.getKnownWord(word: word) != nil {
                counts += 1
            }
        }
        
        return Double(counts) / Double(set.count)
    }

}
