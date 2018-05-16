//
//  TabViewController.swift
//  app
//
//  Created by Sunho Kim on 15/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit

class TabViewController: UIViewController {
    var selectedIndex: Int = 0
    
    var viewControllers: [UIViewController] = []

    @IBOutlet weak var titleLabel: UILabel!
    
    @IBOutlet weak var contentView: UIView!
    @IBOutlet weak var tabBarView: UIView!
    
    @IBOutlet weak var bookTabButton: UIButton!
    @IBOutlet weak var wordbookTabButton: UIButton!
    @IBOutlet var buttons: [UIButton]!
    

    override func viewDidLoad() {
        super.viewDidLoad()
        UIUtill.dropShadow(self.tabBarView, offset: CGSize(width: 0, height: -4), radius: 4)
        
        let storyboard = UIStoryboard(name: "Main", bundle: nil)
        let bookViewController = storyboard.instantiateViewController(withIdentifier: "BookMainViewController")
        let wordbookViewController = storyboard.instantiateViewController(withIdentifier: "WordbookMainViewController")
        
        self.viewControllers = [bookViewController, wordbookViewController]
        
        self.didPressTab(self.buttons[0])
    }
    
    @IBAction func didPressTab(_ sender: UIButton) {
        self.buttons[self.selectedIndex].isSelected = false
        let previousVC = viewControllers[self.selectedIndex]
        previousVC.willMove(toParentViewController: nil)
        previousVC.view.removeFromSuperview()
        previousVC.removeFromParentViewController()
        
        self.selectedIndex = sender.tag
        
        let vc = self.viewControllers[self.selectedIndex]
        
        self.addChildViewController(vc)
        vc.view.frame = self.contentView.bounds
        self.contentView.addSubview(vc.view)
        vc.didMove(toParentViewController: self)
        
        sender.isSelected = true
        self.titleLabel.text = vc.title
    }
}
