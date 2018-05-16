
import UIKit

class WordbookDetailViewController: UIViewController, UITableViewDelegate, UITableViewDataSource {
    @IBOutlet weak var headerView: UIView!
    
    @IBOutlet weak var titleLabel: UILabel!
    
    @IBOutlet weak var memorizeButton: UIButton!
    @IBOutlet weak var flashcardButton: UIButton!
    @IBOutlet weak var sentenceButton: UIButton!
    @IBOutlet weak var speakButton: UIButton!
    
    @IBOutlet weak var tableViewBack: UIView!
    @IBOutlet weak var tableView: UITableView!
    
    @IBOutlet weak var navView: UIView!
    @IBOutlet weak var closeButton: UIButton!
    var wordbook: Wordbook!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.tableView.dataSource = self
        self.tableView.delegate = self
        
        self.titleLabel.text = self.wordbook.name

        self.layout()
    }
    
    fileprivate func layout() {

        UIUtill.roundView(self.memorizeButton)
        UIUtill.roundView(self.flashcardButton)
        UIUtill.roundView(self.speakButton)
        UIUtill.roundView(self.sentenceButton)
        UIUtill.dropShadow(self.tableViewBack, offset: CGSize(width: 0, height: 3), radius: 4)
    }
    
    override func viewDidAppear(_ animated: Bool) {
        super.viewDidAppear(animated)
        
        let view = UIView()
        let label = UILabel()
        label.text = self.wordbook.name
        label.sizeToFit()
        label.alpha = 0
        view.addSubview(label)
        view.frame = label.frame
        
        self.navigationItem.titleView = view
    }
    
    func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
        return UITableViewAutomaticDimension
    }
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return self.wordbook.count
    }
    
    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        tableView.deselectRow(at: indexPath, animated: true)
    }
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: kWordsTableCell)!
        
        let item = self.wordbook.entries[indexPath.row]
        cell.textLabel!.text = item.word

        return cell
    }
    
    @IBAction func didPressClose(_ sender: Any) {
        self.dismiss(animated: true)
    }
}
