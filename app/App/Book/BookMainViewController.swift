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
    
    var books: [String]!
    override func viewDidLoad() {
        super.viewDidLoad()
        
        var tBooks = contentsOfDirectory(path: sharedBooks.path)
        if tBooks == nil {
            try! FileManager.default.createDirectory(atPath: sharedBooks.path, withIntermediateDirectories: true, attributes: nil)
            tBooks = contentsOfDirectory(path: sharedBooks.path)
        }
        self.books = tBooks
        
        self.tableView.delegate = self
        self.tableView.dataSource = self
    }
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return self.books.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "BookMainCell")
        
        let item = self.books[indexPath.row]
        cell!.textLabel!.text = item
        
        return cell!
    }
    

}
