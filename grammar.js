function caseInsensitive(word) {
  return new RegExp(
    word
      .split("")
      .map((character) => `[${character.toLowerCase()}${character.toUpperCase()}]`)
      .join(""),
  );
}

const keyword = (word) => token(prec(4, caseInsensitive(word)));

export default grammar({
  name: "spectra",

  extras: () => [],

  rules: {
    source_file: ($) =>
      seq(repeat(seq(optional($._line), $._newline)), optional($._line)),

    _line: ($) =>
      choice(
        $.comment,
        $.title_statement,
        $.if_statement,
        $.else_statement,
        $.endif_statement,
        $.end_statement,
        $.continuation_line,
        $.statement,
        $.unparsed_line,
      ),

    comment: () => token(prec(5, /[#$*][^\r\n]*/)),

    title_statement: ($) =>
      prec(
        4,
        seq(
          optional($._hspace),
          field("name", alias(keyword("TITLE"), $.statement_name)),
          optional(seq($._hspace, field("title", $.title_text))),
          optional($._hspace),
        ),
      ),

    title_text: () => token(/[^\r\n]*\S/),

    if_statement: ($) =>
      prec(
        4,
        seq(
          optional($._hspace),
          field("name", alias(keyword("IF"), $.statement_name)),
          optional($._hspace),
          field("condition", $.condition),
          optional($._hspace),
        ),
      ),

    else_statement: ($) =>
      prec(
        4,
        seq(
          optional($._hspace),
          field("name", alias(keyword("ELSE"), $.statement_name)),
          optional($._hspace),
        ),
      ),

    endif_statement: ($) =>
      prec(
        4,
        seq(
          optional($._hspace),
          field("name", alias(keyword("ENDIF"), $.statement_name)),
          optional($._hspace),
        ),
      ),

    end_statement: ($) =>
      prec(
        4,
        seq(
          optional($._hspace),
          field("name", alias(keyword("END"), $.statement_name)),
          optional($._hspace),
        ),
      ),

    statement: ($) =>
      seq(
        optional($._hspace),
        field("name", alias($._word, $.statement_name)),
        choice(
          seq(
            $._parameter_separator,
            $.parameter,
            repeat(seq($._parameter_separator, $.parameter)),
            optional($._hspace),
          ),
          prec(-2, optional($._hspace)),
        ),
      ),

    continuation_line: ($) =>
      prec(
        3,
        seq(
          optional($._hspace),
          $.parameter,
          repeat(seq($._parameter_separator, $.parameter)),
          optional($._hspace),
        ),
      ),

    parameter: ($) =>
      seq(
        field("name", alias($._word, $.parameter_name)),
        optional($._hspace),
        field("operator", $.assignment_operator),
        optional($._hspace),
        field("value", $._value),
      ),

    _value: ($) =>
      choice(
        $.quoted_string,
        $.numeric_value,
        $.function_call,
        $.parenthesized_series,
        $.braced_expression,
        $.composite_value,
        alias($._word, $.character_value),
        alias($._bare, $.character_value),
      ),

    quoted_string: () =>
      token(choice(/"([^"\\]|\\.)*"/, /'([^'\\]|\\.)*'/)),

    numeric_value: () =>
      token(
        prec(
          4,
          /[+-]?(?:[0-9]+(?:\.[0-9]*)?|\.[0-9]+)(?:[eEdDbB][+-]?[0-9]+)?/,
        ),
      ),

    composite_value: ($) =>
      prec(
        3,
        seq(
          choice(
            seq($._character_fragment, $.braced_expression),
            seq($.braced_expression, $._character_fragment),
          ),
          repeat(choice($._character_fragment, $.braced_expression)),
        ),
      ),

    _character_fragment: ($) =>
      choice(
        alias($._word, $.character_fragment),
        alias($._bare, $.character_fragment),
      ),

    function_call: ($) =>
      choice(
        prec(
          4,
          seq(
            field("name", alias($._word, $.function_name)),
            "(",
            optional(field("arguments", $.argument_list)),
            ")",
          ),
        ),
        prec(
          -1,
          seq(
            field("name", alias($._word, $.function_name)),
            $._hspace,
            "(",
            optional(field("arguments", $.argument_list)),
            ")",
          ),
        ),
      ),

    argument_list: ($) =>
      seq(
        optional($._value_space),
        $._argument,
        repeat(seq($._value_separator, $._argument)),
        optional($._value_space),
      ),

    _argument: ($) =>
      choice(
        $.quoted_string,
        $.numeric_value,
        $.function_call,
        $.braced_expression,
        $.parenthesized_series,
        $.variable_reference,
        alias($._word, $.character_value),
        alias($._bare, $.character_value),
      ),

    parenthesized_series: ($) =>
      prec(
        2,
        seq(
          "(",
          optional(
            seq(
              optional($._value_space),
              $._series_value,
              repeat(seq($._value_separator, $._series_value)),
              optional($._value_space),
            ),
          ),
          ")",
        ),
      ),

    _series_value: ($) =>
      choice(
        $.quoted_string,
        $.numeric_value,
        $.function_call,
        $.braced_expression,
        $.variable_reference,
        alias($._word, $.character_value),
        alias($._bare, $.character_value),
      ),

    braced_expression: ($) =>
      seq(
        "{",
        repeat(choice($._expression_item, $._expression_space)),
        "}",
      ),

    condition: ($) =>
      seq(
        "[",
        repeat(choice($._expression_item, $._expression_space)),
        "]",
      ),

    _expression_item: ($) =>
      choice(
        $.braced_expression,
        $.function_call,
        $.numeric_value,
        $.format_specifier,
        $.variable_reference,
        $.expression_operator,
        alias($._word, $.identifier),
        alias($._expression_bare, $.expression_fragment),
      ),

    format_specifier: () => token(prec(5, /%[-+0-9.#]*[A-Za-z]/)),

    variable_reference: () =>
      token(prec(5, /[@#$]+[A-Za-z_][A-Za-z0-9_.]*/)),

    expression_operator: () =>
      token(prec(5, /<=|>=|==|!=|\|\||&&|[+\-*\/\\^<>=:,%&!]/)),

    assignment_operator: () => "=",

    unparsed_line: ($) =>
      choice(
        token(prec(-10, /[^\r\n]+/)),
        seq(
          $._hspace,
          token.immediate(prec(-10, /[^\r\n]+/)),
        ),
      ),

    _word: () => token(/[A-Za-z_][A-Za-z0-9_.-]*/),

    _bare: () => token(/[^\s,={}()\[\]'"<>]+/),

    _expression_bare: () =>
      token(/[^\s{}()\[\],+\-*\/\\^<>=:%&!]+/),

    _parameter_separator: ($) =>
      choice(
        $._hspace,
        seq(optional($._hspace), ",", optional($._hspace)),
      ),

    _value_separator: ($) =>
      choice(
        $._value_space,
        seq(optional($._value_space), ",", optional($._value_space)),
      ),

    _value_space: ($) => repeat1(choice($._hspace, $._newline)),

    _expression_space: ($) => choice($._hspace, $._newline, ","),

    _hspace: () => /[ \t]+/,

    _newline: () => /\r?\n/,
  },
});
