import XCTest
@testable import app

class UserDataTest: XCTestCase {
    override class func setUp() {
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
    
    func testKnownWordFromHTML() {
        let html = """
            <html>
                <body>
                    <p>helloo <chunk>from</chunk> the other side</p>
                </body>
            </html>
        """
        try! UserData.shared.addKnownWords(html: html)
        let word = UserData.shared.getKnownWord(word: "helloo")
        XCTAssertNotNil(word)
        
        let word2 = UserData.shared.getKnownWord(word: "from")
        XCTAssertNil(word2)

        let word3 = UserData.shared.getKnownWord(word: "side")
        XCTAssertNotNil(word3)
    }

    func testKnownWordFromSpecialHTML() {
        let html = """
            <html>
                <body>
                    <p>helloo, <chunk>from.</chunk> !the other-side.</p><p>...</p>
                </body>
            </html>
        """
        try! UserData.shared.addKnownWords(html: html)
        let word = UserData.shared.getKnownWord(word: "helloo")
        XCTAssertNotNil(word)
        
        let word2 = UserData.shared.getKnownWord(word: "from")
        XCTAssertNil(word2)
        
        let word3 = UserData.shared.getKnownWord(word: "other-side")
        XCTAssertNotNil(word3)
        
    }
    
}
