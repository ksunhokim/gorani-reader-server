import Foundation
import SQLite

fileprivate let table = Table("wordbooks")
fileprivate let idField = Expression<Int64>("id")
fileprivate let addedDateField = Expression<Date>("added_date")
fileprivate let updatedDateField = Expression<Date>("updated_date")
fileprivate let nameField = Expression<String>("name")

fileprivate let UnknownWordbookId: Int64 = 1

enum WordbookEntryFilter {
    case day(Range<Date>)
    case word(String)
    case book(String)
}

class Wordbook {
    var id: Int64?
    var name: String
    var addedDate: Date
    var updatedData: Date
    
    static let unknown: Wordbook = {
        let wordbook = Wordbook(name: "")
        wordbook.id = UnknownWordbookId
        return wordbook
    }()
    
    // to prevent useless accesses to entries
    var count: Int {
        return WordbookEntry.count(self)
    }
    
    lazy var entries: [WordbookEntry] = {
        return WordbookEntry.get(self)
    }()

    init(name: String) {
        self.name = name
        self.addedDate = Date()
        self.updatedData = Date()
    }
    
    class func get() -> [Wordbook] {
        let query = table.order(addedDateField.desc)
        
        guard let results = try? UserData.shared.connection.prepare(query) else {
            return []
        }
            
        var wordbooks: [Wordbook] = []
        for result in results {
            do{
                let wordbook = Wordbook(name: try result.get(nameField))
                wordbook.id = try result.get(idField)
                wordbook.updatedData = try result.get(updatedDateField)
                wordbook.addedDate = try result.get(addedDateField)
                wordbooks.append(wordbook)
            } catch {}
        }
        
        return wordbooks
    }
    
    func filteredEntries(_ filters: [WordbookEntryFilter]) -> [WordbookEntry] {
        return entries.filter {
            for filter in filters {
                switch filter{
                case .day(let day):
                    if !day.contains($0.addedDate) {
                        return false
                    }
                case .book(let book):
                    if book != $0.book {
                        return false
                    }
                case .word(let word):
                    if word == "" {
                        return false
                    }
                    if $0.word.lowercased().range(of:word) == nil {
                        return false
                    }
                }
                
            }
            return true
        }
    }

    func add() throws {
        self.id = try UserData.shared.connection.run(table.insert(
            addedDateField <- self.addedDate,
            updatedDateField <- self.updatedData,
            nameField <- self.name
        ))
    }
    
    func addEntry(_ entry: WordbookEntry) throws {
        try entry.add(self)
        self.entries.insert(entry, at: 0)
    }
    
    func deleteEntry(at: Int) throws {
        if at >= self.entries.count {
            assertionFailure()
            return
        }
        
        try entries[at].delete(self)
        self.entries.remove(at: at)
    }
    
    class fileprivate func createUnkownWordbook(_ connection: Connection) throws {
        let counts = try connection.scalar(table.count)
        if counts == 0 {
            let id = try connection.run(table.insert(
                addedDateField <- Date(),
                updatedDateField <- Date(),
                nameField <- ""
            ))
            assert(id == UnknownWordbookId)
        }
    }
    
    class func prepare(_ connection: Connection) throws {
        try connection.run(table.create(ifNotExists: true) { t in
            t.column(idField, primaryKey: true)
            t.column(updatedDateField)
            t.column(addedDateField)
            t.column(nameField)
        })
        try connection.run(table.createIndex(addedDateField, unique: false, ifNotExists: true))
        try connection.run(table.createIndex(nameField, unique: false, ifNotExists: true))
        try createUnkownWordbook(connection)
    }
}
