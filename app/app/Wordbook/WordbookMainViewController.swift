//
//  WordbookMainViewController.swift
//  app
//
//  Created by sunho on 2018/04/26.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class WordbookMainViewController: UIViewController, UITableViewDataSource, UITableViewDelegate {

    @IBOutlet var tableView: UITableView!
    
    let asdf = ["ASfd"]
    override func viewDidLoad() {
        super.viewDidLoad()
        
        tableView.tableFooterView = UIView()
        tableView.delegate = self
        tableView.dataSource = self
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }

    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return asdf.count
    }
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "cell")
        cell?.textLabel?.text = asdf[indexPath.row]
        cell?.accessoryType = .disclosureIndicator
        return cell!
    }
    

}
