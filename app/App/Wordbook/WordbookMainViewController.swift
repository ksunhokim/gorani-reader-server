
import UIKit

class WordbookMainViewController: UIViewController, UITableViewDataSource, UITableViewDelegate {
    @IBOutlet var tableView: UITableView!
    
    var wordbooks: [Wordbook] = []
    
    override func viewDidLoad() {
        super.viewDidLoad()

        self.tableView.tableFooterView = UIView()
        self.tableView.delegate = self
        self.tableView.dataSource = self
        self.reloadWordbooks()
    }
    
    func reloadWordbooks() {
        self.wordbooks = Wordbook.get()
    }

    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return self.wordbooks.count
    }
    
    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        tableView.deselectRow(at: indexPath, animated: true)
    }
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: kWordbooksTableCell) as! WordbooksTableCell
        
        let item = self.wordbooks[indexPath.row]
        cell.titleLabel.text = item.name
        cell.countLabel.text = "\(item.count)"
        cell.quizIcons[0].isSelected = true
        UIUtill.dropShadow(cell.back, offset: CGSize(width: 0, height: 3), radius: 4)

        return cell
    }
    
    override func prepare(for segue: UIStoryboardSegue, sender: Any?)
    {
        if segue.destination is WordbookDetailViewController
        {
            let vc = segue.destination as? WordbookDetailViewController
            
            let row = self.tableView.indexPathForSelectedRow!.row
            let item = self.wordbooks[row]
            vc?.wordbook = item
        }
    }
}
