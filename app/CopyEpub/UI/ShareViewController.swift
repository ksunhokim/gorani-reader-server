import UIKit
import Social
import MobileCoreServices
import FolioReaderKit

class ShareViewController: UIViewController {
    @IBOutlet weak var dialogLabel: UILabel!
    
    @IBOutlet weak var bookView: UIView!
    @IBOutlet weak var coverView: UIImageView!
    @IBOutlet weak var titleLabel: UILabel!
    @IBOutlet weak var difficultyLabel: UILabel!
    
    @IBOutlet weak var okayButton: UIButton!
    @IBOutlet weak var noButton: UIButton!

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

    fileprivate func layout() {
        UIUtill.roundView(self.okayButton)
        UIUtill.roundView(self.difficultyLabel)
        UIUtill.dropShadow(self.coverView, offset: CGSize(width: 0, height: 4), radius: 5, alpha: 0.25)
    }
    
    fileprivate func handleAttachment() {
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
            let difficulty = 1 - epub.calculateKnownWordRate()
            DispatchQueue.main.async {
                self.spinner.stopAnimating()
                self.bookURL = epub.tempURL
                self.coverView.image = epub.cover
                self.titleLabel.text = epub.title
                self.dialogLabel.text =  NSLocalizedString("CopyEpubConfirmDialog", comment: "")
                self.okayButton.isHidden = false
                self.difficultyLabel.text = self.getDifficultyString(difficulty)
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

    fileprivate func getDifficultyString(_ difficulty: Double) -> String {
        if difficulty > 0.5 {
            return NSLocalizedString("HardBook", comment: "")
        } else {
            return NSLocalizedString("EasyBook", comment: "")
        }
    }
   
    @IBAction func okButtonTouch(_ sender: Any) {
        self.bookURL.keep = true
        self.dismiss(success: true)
    }
    
    @IBAction func noButtonTouch(_ sender: Any) {
        self.dismiss(success: false)
    }

    fileprivate func handleError(_ e: ShareError) {
        switch e {
        case .notNew:
            self.dialogLabel.text = NSLocalizedString("ExistingBook", comment: "")
            self.bookView.isHidden = true
        default:
            self.dialogLabel.text = "AsdASd"
        }
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
    
    fileprivate func hideExtensionWithCompletionHandler(completion: @escaping (Bool) -> Void) {
        Ease.begin(.quintOut)
        let transform = CGAffineTransform(translationX:0, y: self.view.frame.size.height)
        UIView.animate(
            withDuration: 0.5,
            animations: { self.navigationController!.view.transform = transform },
            completion: completion)
        Ease.end()
    }
}
