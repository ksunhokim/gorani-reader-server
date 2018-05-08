//
//  IPA.swift
//  app
//
//  Created by Sunho Kim on 08/05/2018.
//  Copyright © 2018 sunho. All rights reserved.
//

// https://en.wikipedia.org/wiki/ARPABET

import Foundation
import Regex

fileprivate let stressPattern = try! Regex(pattern:"(.+)[0123]+", groupNames: "arp")
fileprivate let arpToIpa = ["AA":"ɑ","AE":"æ","AH":"ʌ","AO":"ɔ","AW":"aʊ","AY":"aɪ","B":"b","CH":"tʃ","D":"d","DH":"ð","EH":"ɛ","ER":"ɝ","EY":"eɪ","F":"f","G":"ɡ","HH":"h","IH":"ɪ","IY":"i","JH":"dʒ","K":"k","L":"l","M":"m","N":"n","NG":"ŋ","OW":"oʊ","OY":"ɔɪ","P":"p","R":"ɹ","S":"s","SH":"ʃ","T":"t","TH":"θ","UH":"ʊ","UW":"u","V":"v","W":"w","Y":"j","Z":"z","ZH":"ʒ"]

extension String {
    var unstressed: String {
        let arr = self.components(separatedBy: " ")
        var new = ""
        for str in arr {
            let match = stressPattern.findFirst(in: str)
            if let match = match {
                new += match.group(named: "arp")! + " "
            } else {
                new += str + " "
            }
        }
        return String(new.prefix(new.count - 1))
    }

    var ipa: String {
        let str = self.unstressed
        let arr = str.components(separatedBy: " ")
    
        var ipa = ""
        for syl in arr {
            if let ipaSyl = arpToIpa[syl] {
                ipa += ipaSyl
            } else {
                print("undefined arp: \(syl)")
            }
        }
        return ipa
    }
}
