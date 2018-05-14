import UIKit
import OAuthSwift

class AuthViewController: UIViewController {
    var provider: AuthProvider!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        let plist = Bundle.main.path(forResource: kServicesPlistName, ofType: "plist")!
        self.provider = AuthProvider.fromPlist(path: plist)!
    }
    
    @IBAction func beginAuthorization(_ sender: Any) {
        self.provider.beginAuth(name: "naver")
    }
}
