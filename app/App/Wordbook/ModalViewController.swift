import UIKit

class ModalViewController: UIViewController {
    @IBOutlet weak var dialogLabel: UILabel!
    @IBOutlet weak var subDialogLabel: UILabel!
    @IBOutlet weak var modalConstraint: NSLayoutConstraint!
    @IBOutlet weak var nameInput: UITextField!
    @IBOutlet weak var modalView: UIView!
    
    var completion: ((String) -> Void)!
    var dialog: String!
    var subDialog: String!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        let tap: UITapGestureRecognizer = UITapGestureRecognizer(target: self, action: #selector(dismissKeyboard))
        tap.cancelsTouchesInView = false
        self.view.addGestureRecognizer(tap)
        
        NotificationCenter.default.addObserver(self, selector: #selector(keyboardWillShow), name: NSNotification.Name.UIKeyboardWillShow, object: nil)
        NotificationCenter.default.addObserver(self, selector: #selector(keyboardWillHide), name: NSNotification.Name.UIKeyboardWillHide, object: nil)
        
        self.dialogLabel.text = self.dialog
        self.subDialogLabel.text = self.subDialog
        
        self.layout()
    }
    
    fileprivate func layout() {
        UIUtill.dropShadow(self.modalView, offset: CGSize(width: 0, height: 3), radius: 4, alpha: 0.3)
    }
    
    @IBAction func didTabCancel(_ sender: Any) {
        self.dismiss(animated: true)
    }
    
    @IBAction func didTabConfirm(_ sender: Any) {
        self.completion(self.nameInput.text ?? "")
        self.dismiss(animated: true)
    }
    
    @objc func dismissKeyboard() {
        view.endEditing(true)
    }
    
    @objc func keyboardWillShow(notification: NSNotification) {
        guard let keyboardSize = (notification.userInfo?[UIKeyboardFrameBeginUserInfoKey] as? NSValue)?.cgRectValue else {
            return
        }
        
        self.modalConstraint.constant = -keyboardSize.height + self.view.frame.height / 2 - self.modalView.frame.height
        UIView.animate(withDuration: 0.3, animations: {
            self.view.layoutIfNeeded()
        })
    }
    
    @objc func keyboardWillHide(notification: NSNotification) {
        guard (notification.userInfo?[UIKeyboardFrameBeginUserInfoKey] as? NSValue)?.cgRectValue != nil else {
            return
        }
        
        self.modalConstraint.constant = 0
        UIView.animate(withDuration: 0.3, animations: {
            self.view.layoutIfNeeded()
        })
    }
}
