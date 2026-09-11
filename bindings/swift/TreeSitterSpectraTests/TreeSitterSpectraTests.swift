import XCTest
import SwiftTreeSitter
import TreeSitterSpectra

final class TreeSitterSpectraTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_spectra())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Spectra grammar")
    }
}
