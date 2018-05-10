import XCTest
@testable import app
import SQLite

class DictTest: XCTestCase {
    var dict: Dict!

    override func setUp() {
        super.setUp()
        self.dict = Dict.shared
    }
    
    override func tearDown() {
        super.tearDown()
    }
    
    func testGet() {
        let entry = self.dict.get(word: "go")!
        XCTAssert(entry.word == "go")
        XCTAssert(entry.pron == "G OW")
        XCTAssert(entry.pron.ipa == "ɡoʊ")
        
        let entry2 = self.dict.get(word: "dajsfkfa")
        XCTAssert(entry2 == nil)
    }
    
    func testDefs() {
        let entry = self.dict.get(word: "go")!
        let defs = entry.defs
        XCTAssert(defs.count == 40)
        
        let verbs = defs.filter { $0.pos == .verb }
        XCTAssert(verbs.count == 35)
    }
    
    func testExamples() {
        let entry = self.dict.get(word: "go")!
        let defs = entry.defs
        print(defs[0].def)
        XCTAssert(defs[0].examples.count == 8)
        
    }
    
    func testSearchWithBase() {
        let entries = self.dict.search(word: "go", pos: .verb)
        let entry = entries[0]
        XCTAssert(entry.word == "go")
        XCTAssert(entry.defs[0].pos == .verb)
    }
    
    func testSearchWithVariant() {
        let entries = self.dict.search(word: "went", pos: .verb)
        let entry = entries[0] as! DictEntryRedirect
        XCTAssertEqual(entry.word, "go")
        XCTAssertEqual(entry.verbType, .past)

        let entries2 = self.dict.search(word: "goes", pos: .verb)
        let entry2 = entries2[0] as! DictEntryRedirect
        XCTAssertEqual(entry2.word, "go")
        XCTAssertEqual(entry2.verbType, nil)

        let entries3 = self.dict.search(word: "going", pos: .adj)
        let entry3 = entries3[0] as! DictEntryRedirect
        XCTAssertEqual(entry3.word, "go")
        XCTAssertEqual(entry3.verbType, .present)
    }
    
    func testMockSearch() {
        let entries = self.dict.search(word: "", pos: .verb)
        XCTAssert(entries.count == 0)
        
        let entries2 = self.dict.search(word: "ㅇ아아아", pos: .verb)
        XCTAssert(entries2.count == 0)
        
        let entries3 = self.dict.search(word: "🇰🇷", pos: .noun)
        XCTAssert(entries3.count == 0)
    }
}
