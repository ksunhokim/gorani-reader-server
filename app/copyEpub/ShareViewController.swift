//
//  ShareViewController.swift
//  copyEpub
//
//  Created by sunho on 2018/05/01.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit
import Social

@objc(EntryViewController)
class EntryViewController : UINavigationController {
    
    init() {
        super.init(rootViewController: ShareViewController())
    }
    
    required init?(coder aDecoder: NSCoder) {
        super.init(coder: aDecoder)
    }
    
    override init(nibName nibNameOrNil: String?, bundle nibBundleOrNil: Bundle?) {
        super.init(nibName: nibNameOrNil, bundle: nibBundleOrNil)
    }
}

class ShareViewController: UIViewController {
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.view.backgroundColor = UIColor.white
        self.navigationItem.title = "Share this"
        
        self.navigationItem.leftBarButtonItem = UIBarButtonItem(barButtonSystemItem: UIBarButtonSystemItem.cancel, target: self, action: "cancelButtonTapped:")
        self.navigationItem.rightBarButtonItem = UIBarButtonItem(barButtonSystemItem: UIBarButtonSystemItem.save, target: self, action: "saveButtonTapped:")
    }
    
    func saveButtonTapped(sender: UIBarButtonItem) {
        self.hideExtensionWithCompletionHandler { (Bool) -> Void in
            self.extensionContext!.completeRequest(returningItems: nil, completionHandler: nil)
        }
    }
    
    func cancelButtonTapped(sender: UIBarButtonItem) {
        self.hideExtensionWithCompletionHandler{ (Bool) -> Void in
            self.extensionContext!.cancelRequest(withError: NSError())
        }
    }
    
    func hideExtensionWithCompletionHandler(completion:(Bool) -> Void) {
    }
}
