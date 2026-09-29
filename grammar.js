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
  word: ($) => $._word,
  externals: ($) => [$.comment, $._missing_value, $._missing_close, $._error_sentinel],

  conflicts: ($) => [
    [$._value, $.function_call],
    [$._argument, $.function_call],
    [$._series_value, $.function_call],
    [$.positional_argument, $.parameter],
    [$._atom, $.function_call],
    [$._expression, $._or_binary],
    [$._or, $._and_binary],
    [$._and, $._comparison_binary],
    [$._comparison, $._sum_binary],
    [$._sum, $._product_binary],
    [$._power, $._power_binary],
    [$._or_binary, $._and_binary],
    [$._and_binary, $._comparison_binary],
    [$._comparison_binary, $._sum_binary],
    [$._sum_binary, $._product_binary],
  ],

  rules: {
    source_file: ($) =>
      seq(repeat(seq(optional($._line), $._newline)), optional($._line)),

    _line: ($) =>
      choice(
        $._hspace,
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
            choice($.parameter, $.mask_group, $.positional_argument),
            repeat(seq($._parameter_separator, choice($.parameter, $.mask_group, $.positional_argument))),
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
          choice($.parameter, $.mask_group),
          repeat(seq($._parameter_separator, choice($.parameter, $.mask_group))),
          optional($._hspace),
        ),
      ),

    positional_argument: ($) => alias($._word, $.character_value),

    mask_group: ($) =>
      seq(
        optional(seq(field("operator", "&"), optional($._hspace))),
        "[",
        optional($._value_space),
        choice($.parameter, $.mask_group),
        repeat(seq($._value_separator, choice($.parameter, $.mask_group))),
        optional($._value_space),
        choice("]", alias($._missing_close, $.unparsed_line)),
      ),

    linked_series: ($) =>
      seq(field("reference", $.variable_reference), field("values", $.parenthesized_series)),

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
        alias($._missing_value, $.unparsed_line),
        alias($._unfinished_string, $.unparsed_line),
        $.quoted_string,
        $.numeric_value,
        $.function_call,
        $.variable_reference,
        alias($._invalid_reference, $.unparsed_line),
        $.linked_series,
        $.parenthesized_series,
        $.braced_expression,
        $.composite_value,
        alias($._word, $.character_value),
        alias($._bare, $.character_value),
      ),

    quoted_string: () =>
      token(choice(/"([^"\\\r\n]|\\[^\r\n])*"/, /'([^'\\\r\n]|\\[^\r\n])*'/)),

    _unfinished_string: () =>
      token(prec(-1, choice(/"([^"\\\r\n]|\\[^\r\n])*/, /'([^'\\\r\n]|\\[^\r\n])*/))),

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
      prec.dynamic(1, seq(
        field("name", alias($._word, $.function_name)),
        optional($._hspace),
        "(",
        optional(field("arguments", $.argument_list)),
        choice(")", alias($._missing_close, $.unparsed_line)),
      )),

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
        alias($._invalid_reference, $.unparsed_line),
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
          choice(")", alias($._missing_close, $.unparsed_line)),
        ),
      ),

    _series_value: ($) =>
      choice(
        $.quoted_string,
        $.numeric_value,
        $.function_call,
        $.braced_expression,
        $.variable_reference,
        alias($._invalid_reference, $.unparsed_line),
        alias($._word, $.character_value),
        alias($._bare, $.character_value),
      ),

    braced_expression: ($) =>
      seq("{", optional($._hspace), choice($._expression, $.formatted_expression), optional($._hspace),
        choice("}", alias($._missing_close, $.unparsed_line))),

    condition: ($) =>
      seq("[", optional($._hspace), $._expression, optional($._hspace),
        choice("]", alias($._missing_close, $.unparsed_line))),

    formatted_expression: ($) =>
      seq(field("format", $.format_specifier), alias(":", $.expression_operator),
        optional($._expression_space), field("value", $._expression)),

    _expression: ($) => $._or,
    _or: ($) => choice($._and, alias($._or_binary, $.binary_expression)),
    _or_binary: ($) => seq(
      field("left", $._or), optional($._expression_space), field("operator", alias("||", $.expression_operator)),
      optional($._expression_space), field("right", $._and)),
    _and: ($) => choice($._comparison, alias($._and_binary, $.binary_expression)),
    _and_binary: ($) => seq(
      field("left", $._and), optional($._expression_space), field("operator", alias("&&", $.expression_operator)),
      optional($._expression_space), field("right", $._comparison)),
    _comparison: ($) => choice($._sum, alias($._comparison_binary, $.binary_expression)),
    _comparison_binary: ($) => seq(
      field("left", $._comparison), optional($._expression_space), field("operator", alias(choice("=", "==", "!=", "<", ">", "<=", ">="), $.expression_operator)),
      optional($._expression_space), field("right", $._sum)),
    _sum: ($) => choice($._product, alias($._sum_binary, $.binary_expression)),
    _sum_binary: ($) => seq(
      field("left", $._sum), optional($._expression_space), field("operator", alias(choice("+", "-"), $.expression_operator)),
      optional($._expression_space), field("right", $._product)),
    _product: ($) => choice($._unary, alias($._product_binary, $.binary_expression)),
    _product_binary: ($) => seq(
      field("left", $._product), optional($._expression_space), field("operator", alias(choice("*", "/", "\\"), $.expression_operator)),
      optional($._expression_space), field("right", $._unary)),
    _unary: ($) => choice($._power, $.unary_expression),
    _power: ($) => choice($._atom, alias($._power_binary, $.binary_expression)),
    _power_binary: ($) => seq(
      field("left", $._atom), optional($._expression_space), field("operator", alias("^", $.expression_operator)),
      optional($._expression_space), field("right", $._unary)),

    _atom: ($) => choice(
      $.braced_expression, $.parenthesized_expression, $.function_call,
      alias($._expression_number, $.numeric_value), $.variable_reference,
      alias($._invalid_reference, $.unparsed_line),
      alias($._word, $.identifier), alias($._invalid_expression, $.unparsed_line)),

    _invalid_expression: () => token(prec(-20, /[^}\])\r\n \t]+/)),

    _expression_number: () =>
      token(prec(4, /(?:[0-9]+(?:\.[0-9]*)?|\.[0-9]+)(?:[eEdDbB][+-]?[0-9]+)?/)),

    parenthesized_expression: ($) =>
      seq("(", optional($._hspace), $._expression, optional($._hspace),
        choice(")", alias($._missing_close, $.unparsed_line))),

    unary_expression: ($) =>
      seq(field("operator", alias(choice("+", "-", "!"), $.expression_operator)),
        optional($._expression_space), field("operand", $._unary)),

    format_specifier: () => token(prec(5, /%[-+0-9.#]*[A-Za-z]/)),

    variable_reference: () =>
      token(prec(5, /(?:@|\$\$?|#)[A-Za-z_][A-Za-z0-9_.]*/)),

    _invalid_reference: () =>
      token(prec(4, /[@$#]+[A-Za-z_][A-Za-z0-9_.]*/)),


    assignment_operator: () => "=",

    unparsed_line: ($) =>
      choice(
        token(prec(-10, /[^\r\n]+/)),
        seq(
          $._hspace,
          token.immediate(prec(-10, /[^\r\n]+/)),
        ),
      ),

    _word: () => token(/[A-Za-z_][A-Za-z0-9_.]*/),

    _bare: () => token(/[^\s,={}()\[\]'"<>]+/),


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
    _expression_space: ($) => repeat1(choice($._hspace, $._newline)),


    _hspace: () => /[ \t]+/,

    _newline: () => /\r?\n/,
  },
});
