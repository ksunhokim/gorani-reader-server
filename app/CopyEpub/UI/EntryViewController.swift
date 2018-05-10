import UIKit
import FolioReaderKit

@objc(EntryViewController)
class EntryViewController: UINavigationController {
    init() {
        let storyboard = UIStoryboard(name: "Main", bundle: nil)
        let controller = storyboard.instantiateViewController(withIdentifier: "ShareViewController")
        super.init(rootViewController: controller)
    }
    
    required init?(coder aDecoder: NSCoder) {
        super.init(coder: aDecoder)
    }
    
    override init(nibName nibNameOrNil: String?, bundle nibBundleOrNil: Bundle?) {
        super.init(nibName: nibNameOrNil, bundle: nibBundleOrNil)
    }
    
    override func viewDidLoad() {
        super.viewDidLoad()
        self.isNavigationBarHidden = true
    }
    
    override func viewWillAppear(_ animated: Bool) {
        super.viewWillAppear(animated)
        
        self.view.transform = CGAffineTransform(translationX:0, y: self.view.frame.size.height)
        Ease.begin(.quintOut)
        UIView.animate(withDuration: 0.5, animations: { self.view.transform = .identity })
        Ease.end()
    }
}
