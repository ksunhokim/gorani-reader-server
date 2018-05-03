//
//  EpubTest.swift
//  AppTest
//
//  Created by sunho on 2018/05/04.
//  Copyright © 2018 sunho. All rights reserved.
//

import XCTest
@testable import app
import FolioReaderKit

class EpubTest: XCTestCase {
    var epubPath: URL!
    
    override func setUp() {
        super.setUp()
        let path = Bundle.main.path(forResource: "alice", ofType: "epub")!
        let book = try! FREpubParser().readEpub(epubPath: path, removeEpub: false, unzipPath: booksDir.path)
        self.epubPath = booksDir.appendingPathComponent(book.name!)
    }
    
    override func tearDown() {
        super.tearDown()
    }
    
    func testLoad() {
        let epub = try! Epub(bookBase: epubPath)
        XCTAssert(epub.title == "Alice's Adventures in Wonderland")
        XCTAssert(epub.cover == nil)
    }
 
    func testGetLocalBooks() {
        let epubs = Epub.getLocalBooks()
        var okay = false
        for epub in epubs {
            if epub.title == "Alice's Adventures in Wonderland" &&
                epub.cover == nil {
                okay = true
            }
        }
        XCTAssert(okay)
    }

}
