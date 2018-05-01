//
//  BookMainViewController.swift
//  app
//
//  Created by sunho on 2018/05/02.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class BookMainViewController: UINavigationController {
    @IBOutlet weak var tableView: UITableView!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        if let dir = FileManager.default.containerURL(forSecurityApplicationGroupIdentifier: "group.sunho.app") {
            let filename = dir.appendingPathComponent("copy3.txt")
            let text2 = try? String(contentsOf: filename, encoding: .utf8)
            print(text2)
        }
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

}
