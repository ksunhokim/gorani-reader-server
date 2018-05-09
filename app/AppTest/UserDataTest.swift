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
        let word = UserData.shared.getKnownWord(word: "142sdf089hyxcsv")
        XCTAssert(word == nil)
        
        let word2 = KnownWord(word: "hello")
        try! UserData.shared.addKnownWord(word: word2)
        
        let word3 = UserData.shared.getKnownWord(word: "hello")!
        XCTAssert(word3.word == "hello")
        
        try! UserData.shared.deleteKnownWord(word: word3)
        let word4 = UserData.shared.getKnownWord(word: "hello")
        XCTAssert(word4 == nil)
    }

    func testPerformanceExample() {
    }
    
}
