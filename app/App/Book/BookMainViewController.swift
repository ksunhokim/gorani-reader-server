//
//  BookMainViewController.swift
//  app
//
//  Created by sunho on 2018/05/02.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class BookMainViewController: UIViewController, UITableViewDataSource, UITableViewDelegate {
    @IBOutlet weak var tableView: UITableView!
    
    var books: [Epub]!
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.books = Epub.getLocalBooks()
        
        self.tableView.delegate = self
        self.tableView.dataSource = self
    }
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return self.books.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "BooksTableCell") as! BooksTableCell
        
        let item = self.books[indexPath.row]
        cell.titleLabel.text = item.title
        cell.coverImage.image = item.cover
        
        return cell
    }
    

}
