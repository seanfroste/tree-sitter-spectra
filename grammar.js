function caseInsensitive(word) {
  return new RegExp(
    word
      .split("")
      .map((character) => `[${character.toLowerCase()}${character.toUpperCase()}]`)
      .join(""),
  );
}

export default grammar({
  name: "spectra",

  extras: ($) => [/[\t \r]/],

  rules: {
    source_file: ($) => repeat(choice($.title_statement, $.end_statement, $._newline)),

    title_statement: ($) =>
      seq(
        field("name", alias(caseInsensitive("TITLE"), $.statement_name)),
        optional(field("title", alias(token(/[^\r\n]+/), $.title_text))),
      ),

    end_statement: ($) =>
      field("name", alias(caseInsensitive("END"), $.statement_name)),

    _newline: () => /\n/,
  },
});
