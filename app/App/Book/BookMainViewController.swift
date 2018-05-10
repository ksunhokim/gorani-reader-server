//
//  BookMainViewController.swift
//  app
//
//  Created by sunho on 2018/05/02.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit
import FolioReaderKit

fileprivate let MinActulReadRate = 0.7
class BookMainViewController: UIViewController, UITableViewDataSource, UITableViewDelegate, FolioReaderDelegate, FolioReaderCenterDelegate {
    @IBOutlet weak var tableView: UITableView!
    
    var books: [Epub]!
    var dict: Dict!
    var folioReader = FolioReader()
    var currentHTML: String?
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.books = Epub.getLocalBooks()
        
        self.tableView.delegate = self
        self.tableView.dataSource = self
        self.folioReader.delegate = self
    }
    
    func presentDictView(bookName: String, page: Int, scroll: CGFloat, sentence: String, word: String, index: Int) {
        let viewController = DictViewController(word: word, sentence: sentence, index: index)
        self.folioReader.readerContainer?.present(viewController, animated:  true)
        print(bookName, page, scroll, sentence, word, index)
    }
    
    fileprivate func calculateKnownWords() {
        if let html = self.currentHTML {
            if self.folioReader.readerCenter!.actualReadRate > MinActulReadRate {
                
            }
        }
    }

    func folioReaderDidClose(_ folioReader: FolioReader) {
        self.calculateKnownWords()
        self.currentHTML = nil
    }
    
    func htmlContentForPage(_ page: FolioReaderPage, htmlContent: String) -> String {
        self.calculateKnownWords()
        self.currentHTML = htmlContent
        return htmlContent
    }
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return self.books.count
    }
    
    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        tableView.deselectRow(at: indexPath, animated: true)
        let item = self.books[indexPath.row]
        let config = FolioReaderConfig()
        config.tintColor = UIUtill.blue
        config.canChangeScrollDirection = false
        config.hideBars = false
        config.scrollDirection = .horizontal
        self.folioReader.presentReader(parentViewController: self, book: item.book!, config: config)
        self.folioReader.readerCenter!.delegate = self
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "BooksTableCell") as! BooksTableCell
        
        let item = self.books[indexPath.row]
        cell.titleLabel.text = item.title
        cell.coverImage.image = item.cover
        
        return cell
    }
}
