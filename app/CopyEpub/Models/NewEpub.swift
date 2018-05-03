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

class NewEpub {
    private let parser = FREpubParser()
    
    let tempURL: ManagedEpubURL
    let book: FRBook
    
    var name: String = ""
    var cover: UIImage?
    
    init(epub: URL) throws {
        let tempURL = ManagedEpubURL(epub: epub)
        guard tempURL.isNew() else {
            tempURL.keep = true
            throw ShareError.notNew
        }
        self.tempURL = tempURL
        
        self.book = try self.parser.readEpub(epubPath: epub.path, removeEpub: false, unzipPath: booksDir.path)
        try self.parse()
    }
    
    private func parse() throws {
        if let image = self.book.coverImage {
            self.cover = UIImage(contentsOfFile: image.fullHref)
        }
    
        guard let name = self.book.name else {
            throw ShareError.notProperEpub
        }
        self.name = name
    }
}
