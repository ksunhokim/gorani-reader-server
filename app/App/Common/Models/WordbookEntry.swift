import Foundation
import SQLite

fileprivate let table = Table("wordbook_entries")
fileprivate let wordbookTable = Table("wordbooks")
fileprivate let wordbookIdField = Expression<Int64>("wordbook_id")
fileprivate let idField = Expression<Int64>("id")
fileprivate let addedDateField = Expression<Date>("added_date")
fileprivate let sentenceField = Expression<String>("sentence")
fileprivate let indexField = Expression<Int>("index")
fileprivate let wordField = Expression<String>("word")
fileprivate let defField = Expression<String>("def")
fileprivate let bookField = Expression<String>("book")

class WordbookEntry {
    var id: Int64?
    var addedDate: Date
    var sentence: String
    var index: Int
    var word: String
    var def: String
    var book: String
    
    init(sentence: String, index: Int, word: String, def: String, book: String) {
        self.addedDate = Date()
        self.sentence = sentence
        self.index = index
        self.word = word
        self.def = def
        self.book = book
    }
    
    class func count(_ wordbook: Wordbook) -> Int {
        guard let wordbookId = wordbook.id else {
            fatalError("wordbook hasn't been initialized")
        }
        
        let query = table.where(wordbookIdField == wordbookId).count
        if let counts = try? UserData.shared.connection.scalar(query) {
            return counts
        }
        
        return 0
    }
    
    class func get(_ wordbook: Wordbook) -> [WordbookEntry] {
        guard let wordbookId = wordbook.id else {
            fatalError("wordbook hasn't been initialized")
        }
        
        let query = table.where(wordbookId == wordbookIdField)
        var entries: [WordbookEntry] = []
        do {
            for entry in try UserData.shared.connection.prepare(query) {
                let wEntry = WordbookEntry(
                    sentence: try entry.get(sentenceField),
                    index: try entry.get(indexField),
                    word: try entry.get(wordField),
                    def: try entry.get(defField),
                    book: try entry.get(bookField)
                    )
                wEntry.id = try entry.get(idField)
                wEntry.addedDate = try entry.get(addedDateField)
                entries.append(wEntry)
            }
        } catch {}
        
        return entries
    }

    func add(_ wordbook: Wordbook) throws {
        guard let wordbookId = wordbook.id else {
            fatalError("wordbook hasn't been initialized")
        }
        
        self.id = try UserData.shared.connection.run(table.insert(
            wordbookIdField <- wordbookId,
            addedDateField <- self.addedDate,
            sentenceField <- self.sentence,
            indexField <- self.index,
            wordField <- self.word,
            defField <- self.def,
            bookField <- self.book
        ))
    }
    
    func delete(_ wordbook: Wordbook) throws {
        guard let id = self.id else {
            fatalError("wordbookentry hasn't been initialized")
        }
        
        let me = table.where(idField == id)
        try UserData.shared.connection.run(me.delete())
    }
    
    class func perpare(_ connection: Connection) throws {
        try connection.run(table.create(ifNotExists: true) { t in
            t.column(wordbookIdField, references: wordbookTable, idField)
            t.column(idField, primaryKey: true)
            t.column(addedDateField)
            t.column(sentenceField)
            t.column(indexField)
            t.column(wordField)
            t.column(defField)
            t.column(bookField)
        })
        try connection.run(table.createIndex(addedDateField, ifNotExists: true))
        try connection.run(table.createIndex(wordField, ifNotExists: true))
        try connection.run(table.createIndex(bookField, ifNotExists: true))
    }
}
