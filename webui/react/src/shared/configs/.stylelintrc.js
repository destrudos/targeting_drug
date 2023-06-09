module.exports = {
  extends: [
    'stylelint-config-standard',
    'stylelint-config-prettier', // stylelint-config-prettier should be the last
  ],
  plugins: ['stylelint-order', 'stylelint-scss'],
  rules: {
    'at-rule-no-unknown': null,
    'at-rule-semicolon-space-before': 'never',
    'custom-property-empty-line-before': 'never',
    'declaration-block-semicolon-newline-after': 'always-multi-line',
    'declaration-block-semicolon-newline-before': 'never-multi-line',
    'declaration-block-semicolon-space-before': 'never',
    'declaration-block-trailing-semicolon': null,
    'declaration-empty-line-before': 'never',
    'function-name-case': 'lower',
    'no-eol-whitespace': [true, { ignore: ['empty-lines'] }],
    'no-extra-semicolons': true,
    'order/order': [
      'custom-properties',
      'dollar-variables',
      'at-variables',
      'declarations',
      'rules',
      'at-rules',
      'less-mixins',
    ],
    'order/properties-alphabetical-order': true,
    'rule-empty-line-before': [
      'always',
      {
        except: ['after-rule', 'first-nested', 'inside-block-and-after-rule'],
      },
    ],
    'scss/at-rule-no-unknown': true,
    'selector-pseudo-class-no-unknown': [true, { ignorePseudoClasses: ['global'] }],
    'string-quotes': 'single',
    'value-keyword-case': null,
  },
};
