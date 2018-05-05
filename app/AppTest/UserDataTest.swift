//
//  UserDataTest.swift
//  AppTest
//
//  Created by Sunho Kim on 05/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import XCTest
@testable import app

class UserDataTest: XCTestCase {
    override func setUp() {
        super.setUp()
        try? app.fileManager.removeItem(at: userDataURL)
    }
    
    override func tearDown() {
        super.tearDown()
    }
    
    func testKnownWord() {
        let word = KnownWord.get(userData, word: "142sdf089hyxcsv")
        XCTAssert(word == nil)
        
        let word2 = KnownWord(word: "hello")
        try! word2.add(userData)
        
        let word3 = KnownWord.get(userData, word: "hello")!
        XCTAssert(word3.word == "hello")
        
        try! word3.delete(userData)
        let word4 = KnownWord.get(userData, word:"hello")
        XCTAssert(word4 == nil)
    }

    func testPerformanceExample() {
    }
    
}
