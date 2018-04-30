//
//  TabBarController.swift
//  app
//
//  Created by sunho on 2018/04/30.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class TabBarController: UITabBarController, UITabBarDelegate {

    override func viewDidLoad() {
        super.viewDidLoad()

        self.setValue(UITabBar(), forKey: "tabBar")
        // Do any additional setup after loading the view.
    }
    
    override func tabBar(tabBar: UITabBar, didSelectItem item: UITabBarItem) {
        print("Selected item")
    }
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }
    
    func asdf() {
    }

}
