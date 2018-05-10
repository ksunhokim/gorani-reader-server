
import UIKit

fileprivate let MaxChar = 120

class DictViewController: UIViewController, UITableViewDelegate, UITableViewDataSource {
    var tableView: UITableView!
    var sentenceLabel: UILabel!
    var cancelButton: UIButton!
    
    var word: String
    var sentence: String
    var index: Int
    
    var entries: [DictEntry]
    
    override var prefersStatusBarHidden: Bool {
        return true
    }
    
    init(word: String, sentence: String, index: Int) {
        self.word = word
        self.sentence = sentence
        self.index = index
        self.entries = Dict.shared.search(word: word)
        super.init(nibName: nil, bundle: Bundle.main)
    }
    
    required init?(coder aDecoder: NSCoder) {
        fatalError("storyboard is not good" )
    }

    @objc func onCacnelButton(_ sender: Any? = nil) {
        dismiss(animated: true)
    }
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        
        self.view.backgroundColor = UIColor.white
        let attrs = [NSAttributedStringKey.font : UIFont.boldSystemFont(ofSize: 20)]

        let (front, middle, end) = self.getFrontMiddleEnd()
        let frontString = NSMutableAttributedString(string: front)
        let middleString = NSMutableAttributedString(string: middle, attributes:attrs)
        let endString = NSMutableAttributedString(string: end)
        frontString.append(middleString)
        frontString.append(endString)
        
        self.sentenceLabel = UILabel(frame: CGRect(x: 14, y: 20, width: view.frame.width - 28, height: 70))
        self.sentenceLabel.attributedText = frontString
        self.sentenceLabel.textAlignment = .center
        self.sentenceLabel.numberOfLines = 0
        self.view.addSubview(self.sentenceLabel)
        
        let line = UIView(frame: CGRect(x: 0, y: self.sentenceLabel.frame.height + self.sentenceLabel.frame.origin.y + 20, width: view.frame.width, height: 0.7))
        line.backgroundColor = UIUtill.gray1

        self.tableView = UITableView(frame: CGRect(x: 0, y: line.frame.origin.y + 0.7 , width: view.frame.width, height: view.frame.height - 200))
        self.tableView.delegate = self
        self.tableView.dataSource = self
        self.tableView.backgroundColor = UIUtill.lightGray1
        self.tableView.separatorStyle = .none
        self.tableView.rowHeight = UITableViewAutomaticDimension;
        self.tableView.estimatedRowHeight = 100;
        self.tableView.showsVerticalScrollIndicator = false
        self.tableView.register(UINib(nibName: kDictViewTableCell, bundle: nil), forCellReuseIdentifier: kDictViewTableCell)
        self.tableView.contentInset = UIEdgeInsets(top: 0, left: 0, bottom: 10, right: 0)
        self.view.addSubview(self.tableView)

        self.cancelButton = UIButton(frame: CGRect(x: 14, y: view.frame.height - 70, width: view.frame.width - 28, height: 50))
        self.cancelButton.backgroundColor = UIUtill.blue
        self.cancelButton.setTitleColor(UIUtill.white, for: .normal)
        self.cancelButton.setTitle("Cancel", for: .normal)
        self.cancelButton.addTarget(self, action: #selector(onCacnelButton(_:)), for: .touchUpInside)
        UIUtill.roundView(self.cancelButton)
        self.view.addSubview(self.cancelButton)
    }
    
    func numberOfSections(in tableView: UITableView) -> Int {
        return self.entries.count
    }
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return self.entries[section].defs.count
    }
    
    func tableView(_ tableView: UITableView, heightForHeaderInSection section: Int) -> CGFloat {
        return 58
    }
    
    func tableView(_ tableView: UITableView, viewForHeaderInSection section: Int) -> UIView? {
        let entry = self.entries[section]
        
        let whole = UIView()
        
        let back = UIView(frame: CGRect(x: 0, y: 0, width: self.tableView.bounds.width, height: 40))
        back.backgroundColor = UIUtill.lightGray1
        whole.addSubview(back)
        
        let view = UIView(frame: CGRect(x: 4, y: 8, width: self.tableView.bounds.width - 8, height: 50))
        view.backgroundColor = UIUtill.lightGray0
        UIUtill.roundView(view)
        back.addSubview(view)
        
        let label = UILabel()
        label.font = UIFont.boldSystemFont(ofSize: 20)
        label.frame.origin.x = 14
        label.textColor = UIUtill.gray2
        label.text = entry.word
        label.sizeToFit()
        label.frame = CGRect(origin: label.frame.origin, size: CGSize(width: label.frame.width, height: 50))
        view.addSubview(label)
        
        let typeButton = UIButton()
        typeButton.setTitle(entry.pron.ipa, for: .normal)
        typeButton.contentEdgeInsets = UIEdgeInsets(top: 0, left: 20, bottom: 0, right: 20)
        typeButton.backgroundColor = UIUtill.green
        typeButton.setTitleColor(UIUtill.white, for: .normal)
        typeButton.titleLabel?.font = UIFont.systemFont(ofSize: 14)
        typeButton.titleLabel?.baselineAdjustment = .alignCenters
        typeButton.contentVerticalAlignment = .center
        typeButton.titleLabel?.sizeToFit()
        typeButton.sizeToFit()
        typeButton.frame = CGRect(x: view.frame.width -  typeButton.frame.width - 5, y: 5, width: typeButton.frame.width, height: 40)
        UIUtill.roundView(typeButton)
        view.addSubview(typeButton)
        
        return whole
    }
    
    fileprivate func getDictEntryColor(entry: DictEntry) -> UIColor {
        if entry is DictEntryRedirect {
            return UIUtill.green
        }
        return UIUtill.lightGray0
    }
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let entry = self.entries[indexPath.section].defs[indexPath.row]
        
        let cell = self.tableView.dequeueReusableCell(withIdentifier: kDictViewTableCell, for: indexPath) as! DictViewTableCell
        cell.backgroundColor = UIColor.clear
        cell.label.text = entry.def
        
        return cell
    }

    fileprivate func getFrontMiddleEnd() -> (String, String, String) {
        let arr = self.sentence.components(separatedBy: " ")
        
        var front = ""
        var end = ""
        
        if arr.count > index {
            let frontarr = arr[...(index - 1)]
            front = frontarr.joined(separator: " ")
            
            let endarr = arr[(index + 1)...]
            end = endarr.joined(separator: " ")
        }
        
        // trim
        var frontLength = 0
        let candidate1 = min(front.count, MaxChar / 2)
        let candidate2 = min(end.count, MaxChar / 2)
        if candidate1 < candidate2 {
            frontLength = candidate1
        } else {
            frontLength = MaxChar - candidate2
        }
        
        front = String(front.suffix(frontLength))
        end = String(end.prefix(MaxChar - frontLength))
        return (front, " \(arr[index]) ", end)
    }
}
