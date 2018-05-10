import Foundation

class Wordbook {
    var id: Int
    var name: String
    var entries: [WordbookEntry] = []
    
    init(id: Int, name: String) {
        self.id = id
        self.name = name
    }
    
    func addEntry(_ entry: WordbookEntry) {
        
    }
}
