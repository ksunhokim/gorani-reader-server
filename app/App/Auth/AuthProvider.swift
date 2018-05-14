import Foundation
import OAuthSwift

fileprivate let StateLength = 20

struct AuthService: Decodable {
    var key: String
    var secret: String
    var authorizeurl: String
    var tokenurl: String
    var callback: String
}

class AuthProvider {
    var oauthswift: OAuth2Swift!
    var services: [String:AuthService] = [:]
    
    init() {}
    
    class func fromPlist(path: String) -> AuthProvider? {
        guard let data = try? Data(contentsOf: URL(fileURLWithPath: path)) else {
            assertionFailure()
            return nil
        }
        
        let provider = AuthProvider()
        guard let services = try? PropertyListDecoder().decode([String:AuthService].self, from: data) else {
            assertionFailure()
            return nil
        }
        
        provider.services = services

        return provider
    }
    
    func beginAuth(name: String, _ success: ((String) -> Void)? = nil, _ failure: ((OAuthSwiftError) -> Void)? = nil) {
        
        guard let service = self.services[name] else {
            assertionFailure()
            return
        }
        
        self.oauthswift = OAuth2Swift(
            consumerKey: service.key,
            consumerSecret: service.secret,
            authorizeUrl: service.authorizeurl,
            responseType: "code"
        )
    
        self.oauthswift.authorize(
            withCallbackURL: URL(string: service.callback)!,
            scope: "",
            state: generateState(withLength: StateLength),
            success: { credential, response, parameters in
                if let success = success {
                    success(credential.oauthToken)
                }
        },
            failure: { error in
                if let failure = failure {
                    failure(error)
                }
        }
        )
    }
}
