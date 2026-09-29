#include "tree_sitter/parser.h"

#include <string.h>

enum TokenType { COMMENT, MISSING_VALUE, MISSING_CLOSE, ERROR_SENTINEL };

static bool word_start(int32_t c) {
  return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || c == '_';
}

static bool word_part(int32_t c) {
  return word_start(c) || (c >= '0' && c <= '9') || c == '.';
}

static void advance(TSLexer *lexer) { lexer->advance(lexer, false); }

void *tree_sitter_spectra_external_scanner_create(void) { return NULL; }
void tree_sitter_spectra_external_scanner_destroy(void *payload) { (void)payload; }
unsigned tree_sitter_spectra_external_scanner_serialize(void *payload, char *buffer) {
  (void)payload;
  (void)buffer;
  return 0;
}
void tree_sitter_spectra_external_scanner_deserialize(void *payload, const char *buffer, unsigned length) {
  (void)payload;
  (void)buffer;
  (void)length;
}

bool tree_sitter_spectra_external_scanner_scan(void *payload, TSLexer *lexer, const bool *valid) {
  (void)payload;
  // Do not manufacture zero-width tokens during Tree-sitter's recovery search.
  if (valid[ERROR_SENTINEL]) return false;
  if (valid[COMMENT] && lexer->get_column(lexer) == 0 &&
      (lexer->lookahead == '#' || lexer->lookahead == '$' || lexer->lookahead == '*')) {
    do { advance(lexer); } while (!lexer->eof(lexer) && lexer->lookahead != '\r' && lexer->lookahead != '\n');
    lexer->mark_end(lexer);
    lexer->result_symbol = COMMENT;
    return true;
  }
  if (lexer->lookahead != '\r' && lexer->lookahead != '\n' && !lexer->eof(lexer)) return false;
  lexer->mark_end(lexer);
  if (valid[MISSING_VALUE]) {
    lexer->result_symbol = MISSING_VALUE;
    return true;
  }
  if (!valid[MISSING_CLOSE]) return false;
  if (lexer->eof(lexer)) {
    lexer->result_symbol = MISSING_CLOSE;
    return true;
  }
  // A newline is legal inside series and mask groups. Only close an incomplete
  // value before an unmistakable next statement, never before another value.
  while (lexer->lookahead == '\r' || lexer->lookahead == '\n' || lexer->lookahead == ' ' || lexer->lookahead == '\t') advance(lexer);
  if (!word_start(lexer->lookahead)) return false;
  char name[6] = {0};
  unsigned length = 0;
  while (word_part(lexer->lookahead)) {
    if (length < sizeof(name) - 1) {
      int32_t c = lexer->lookahead;
      name[length] = (char)(c >= 'a' && c <= 'z' ? c - 'a' + 'A' : c);
    }
    ++length;
    advance(lexer);
  }
  if ((length == 3 && strcmp(name, "END") == 0) ||
      (length == 4 && strcmp(name, "ELSE") == 0) ||
      (length == 5 && (strcmp(name, "ENDIF") == 0 || strcmp(name, "TITLE") == 0))) {
    if (lexer->lookahead == '\r' || lexer->lookahead == '\n' || lexer->lookahead == ' ' || lexer->lookahead == '\t' || lexer->eof(lexer)) {
      lexer->result_symbol = MISSING_CLOSE;
      return true;
    }
  }
  if (lexer->lookahead != ' ' && lexer->lookahead != '\t') return false;
  while (lexer->lookahead == ' ' || lexer->lookahead == '\t') advance(lexer);
  if (!word_start(lexer->lookahead)) return false;
  do { advance(lexer); } while (word_part(lexer->lookahead));
  while (lexer->lookahead == ' ' || lexer->lookahead == '\t') advance(lexer);
  if (lexer->lookahead != '=') return false;
  lexer->result_symbol = MISSING_CLOSE;
  return true;
}
