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
import FolioReaderKit

class ShareViewController: UIViewController {
    @IBOutlet weak var okayButton: UIButton!
    @IBOutlet weak var noButton: UIButton!
    @IBOutlet weak var titleLabel: UILabel!
    @IBOutlet weak var dialogLabel: UILabel!
    @IBOutlet weak var bookView: UIView!
    @IBOutlet weak var coverView: UIImageView!

    var spinner: Spinner!
    var bookURL: ManagedEpubURL!

    override func viewDidLoad() {
        super.viewDidLoad()

        //loading
        self.spinner = Spinner(target: self.bookView)
        self.view.addSubview(self.spinner)
        self.spinner.startAnimating()
        self.okayButton.isHidden = true
        
        self.layout()
        self.handleAttachment()
    }

    func layout() {
        self.okayButton.layer.cornerRadius = 10
        self.okayButton.clipsToBounds = true
    }
    
    func handleAttachment() {
        let content = self.extensionContext!.inputItems[0] as! NSExtensionItem
        let attachment = content.attachments!.first as! NSItemProvider
        
        guard attachment.hasItemConformingToTypeIdentifier("public.url") else {
            self.handleError(.notURL)
            return
        }
        attachment.loadItem(forTypeIdentifier: "public.url", options: nil, completionHandler: self.loadCompletionHandler(coding:error:))
    }

    func loadCompletionHandler(coding:NSSecureCoding?, error:Error!) {
        guard let url = coding as? URL else {
            DispatchQueue.main.async {
                self.spinner.stopAnimating()
                self.handleError(ShareError.system)
            }
            return
        }
        do {
            let epub = try NewEpub(epub: url)
            DispatchQueue.main.async {
                self.spinner.stopAnimating()
                self.bookURL = epub.tempURL
                self.coverView.image = epub.cover
                self.titleLabel.text = epub.name
                self.dialogLabel.text = "이 책을 고라니 리더로 가져오겠습니까?"
                self.okayButton.isHidden = false
            }
        } catch let err as ShareError {
            DispatchQueue.main.async {
                self.spinner.stopAnimating()
                self.handleError(err)
            }
        } catch {
            DispatchQueue.main.async {
                self.spinner.stopAnimating()
                self.handleError(ShareError.system)
            }
        }
        
    }

   
    @IBAction func okButtonTouch(_ sender: Any) {
        self.bookURL.keep = true
        self.dismiss(success: true)
    }
    
    @IBAction func noButtonTouch(_ sender: Any) {
        self.dismiss(success: false)
    }
    
    func handleError(_ e: ShareError) {
        dismiss(success: false)
    }
    
    func dismiss(success: Bool) {
        self.hideExtensionWithCompletionHandler(completion: { (Bool) in
            if success {
                self.extensionContext!.completeRequest(returningItems: nil, completionHandler: nil)
            } else {
                self.extensionContext!.cancelRequest(withError: NSError(domain: "copyEpub", code: 0, userInfo: nil))
            }
        })
    }
    
    func hideExtensionWithCompletionHandler(completion: @escaping (Bool) -> Void) {
        Ease.begin(.quintOut)
        let transform = CGAffineTransform(translationX:0, y: self.view.frame.size.height)
        UIView.animate(
            withDuration: 0.5,
            animations: { self.navigationController!.view.transform = transform },
            completion: completion)
        Ease.end()
    }
}
