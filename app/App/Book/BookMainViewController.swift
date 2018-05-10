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
    
    @objc func applicationWillEnterForeground(_ notification: NSNotification) {
        self.books = Epub.getLocalBooks()
        self.tableView.reloadData()
    }
    
    override func viewDidAppear(_ animated: Bool) {
        NotificationCenter.default.addObserver(self, selector:#selector(applicationWillEnterForeground(_:)), name:NSNotification.Name.UIApplicationWillEnterForeground, object: nil)
    }
    
    override func viewWillDisappear(_ animated: Bool) {
        super.viewWillDisappear(animated)
        NotificationCenter.default.removeObserver(self)
    }
    
    func presentDictView(bookName: String, page: Int, scroll: CGFloat, sentence: String, word: String, index: Int) {
        let viewController = DictViewController(word: word, sentence: sentence, index: index)
        self.folioReader.readerContainer?.present(viewController, animated:  true)
        print(bookName, page, scroll, sentence, word, index)
    }
    
    fileprivate func calculateKnownWords() {
        guard let html = self.currentHTML else {
            return
        }
        
        if self.folioReader.readerCenter!.actualReadRate > MinActulReadRate {
            DispatchQueue.global(qos: .default).async {
                try? UserData.shared.addKnownWords(html: html)
            }
        }
    }

    func folioReaderDidClose(_ folioReader: FolioReader) {
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
        let book = self.books[indexPath.row]
        
        let config = FolioReaderConfig()
        config.tintColor = UIUtill.blue
        config.canChangeScrollDirection = false
        config.hideBars = false
        config.scrollDirection = .horizontal
        
        self.folioReader.presentReader(parentViewController: self, book: book.book!, config: config)
        self.folioReader.readerCenter!.delegate = self
        
        self.tableView.deselectRow(at: indexPath, animated: true)
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "BooksTableCell") as! BooksTableCell
        
        let item = self.books[indexPath.row]
        cell.titleLabel.text = item.title
        cell.coverImage.image = item.cover
        
        return cell
    }
}
