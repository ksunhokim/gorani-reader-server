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
}
