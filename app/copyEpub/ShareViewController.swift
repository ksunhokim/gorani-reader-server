//
//  ShareViewController.swift
//  copyEpub
//
//  Created by sunho on 2018/05/01.
//  Copyright © 2018 sunho. All rights reserved.
//

import UIKit
import Social
import MobileCoreServices

class ShareViewController: UIViewController {
    @IBOutlet weak var okayButton: UIButton!
    @IBOutlet weak var noButton: UIButton!
    @IBOutlet weak var titleLabel: UILabel!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.okayButton.layer.cornerRadius = 10
        self.okayButton.clipsToBounds = true
        
        let content = self.extensionContext!.inputItems[0] as! NSExtensionItem
        let attachment = content.attachments!.first as! NSItemProvider
        if attachment.hasItemConformingToTypeIdentifier("public.url") {
            attachment.loadItem(forTypeIdentifier: "public.url", options: nil,
                                completionHandler: {
                                    (coding:NSSecureCoding?, error:Error!) in
                                    let url = coding as? NSURL
                                    self.titleLabel.text = url?.absoluteString
            })
        }
    }
    

    @IBAction func okButtonTouch(_ sender: Any) {
        self.extensionContext!.completeRequest(returningItems: nil, completionHandler: nil)
    }
    
    @IBAction func noButtonTouch(_ sender: Any) {
        hideExtensionWithCompletionHandler(completion: { (Bool) in
            self.extensionContext!.cancelRequest(withError: NSError())
        })
    }
    
    func hideExtensionWithCompletionHandler(completion: @escaping (Bool) -> Void) {
        CATransaction.begin()
        let timing = CAMediaTimingFunction(controlPoints: 0.23, 1, 0.32, 1)
        CATransaction.setAnimationTimingFunction(timing)
        UIView.animate(
            withDuration: 0.5,
            animations: { () -> Void in
                self.navigationController!.view.transform = CGAffineTransform(translationX: 0, y: self.navigationController!.view.frame.size.height)
        },completion: completion)
        CATransaction.commit()
    }
}
