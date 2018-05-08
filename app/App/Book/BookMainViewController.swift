//
//  BookMainViewController.swift
//  app
//
//  Created by sunho on 2018/05/02.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit
import FolioReaderKit

class BookMainViewController: UIViewController, UITableViewDataSource, UITableViewDelegate {
    @IBOutlet weak var tableView: UITableView!
    
    var books: [Epub]!
    var folioReader = FolioReader()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.books = Epub.getLocalBooks()
        
        self.tableView.delegate = self
        self.tableView.dataSource = self
    }
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return self.books.count
    }
    
    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        tableView.deselectRow(at: indexPath, animated: true)
        let item = self.books[indexPath.row]
        let config = FolioReaderConfig()
        config.tintColor = UIColor(rgba: "#007AFF")
        config.canChangeScrollDirection = false
        config.hideBars = false
        config.scrollDirection = .horizontal
        self.folioReader.presentReader(parentViewController: self, book: item.book!, config: config)
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "BooksTableCell") as! BooksTableCell
        
        let item = self.books[indexPath.row]
        cell.titleLabel.text = item.title
        cell.coverImage.image = item.cover
        
        return cell
    }
}
