# Text Editing/Auto-Correction Tool

## Description

The Text Editing/Auto-Correction Tool is a Go program that reads an input text file, applies specific modifications to the text according to a set of defined rules, and writes the modified text to an output file. The modifications include converting hexadecimal and binary numbers to their decimal equivalents, changing words to uppercase, lowercase, or capitalized format, adjusting punctuation spacing, and altering the usage of the indefinite articles "a" and "an" based on the following word.

## Project Objectives

- Learn how to use Go's file system (fs) API.
- Practice string and numbers manipulation in Go.
- Implement text editing and auto-correction features based on a set of rules.

## Modifications

The tool supports the following modifications:

- **Hexadecimal to Decimal:** Convert words marked with `(hex)` to their decimal representation.
- **Binary to Decimal:** Convert words marked with `(bin)` to their decimal representation.
- **Case Conversions:** Convert words marked with `(up)`, `(low)`, or `(cap)` to uppercase, lowercase, or capitalized format, respectively. This can also apply to a specified number of words following the format `(up, <number>)`, `(low, <number>)`, or `(cap, <number>)`.
- **Punctuation Adjustment:** Adjust spacing around punctuation marks such as `.`, `,`, `!`, `?`, `:` and `;` to be close to the previous word and spaced from the next word, except for groups of punctuation like `...` or `!?`.
- **A/An Usage:** Correct the usage of "a" and "an" based on whether the following word begins with a vowel (a, e, i, o, u) or an "h".
- **Quotation Marks:** Adjust the placement of quotation marks (`'`) to be adjacent to the word or phrase enclosed within them.

## Usage

To run the program, use the following command:

```sh
go run . <input_file> <output_file>
```

- `<input_file>`: The path to the input file containing the text to be modified.
- `<output_file>`: The path to the output file where the modified text will be saved.

For example:

```sh
go run . sample.txt result.txt
```

After running the program, the modified text will be saved to the specified output file.

## Project Structure

- `main.go`: The main program file that reads the input file, applies the modifications, and writes the modified text to the output file.
- `test files`: Test files for unit testing to ensure the program functions correctly.

## Good Practices

- Follow Go's standard coding style and best practices.
- Write clear and concise code comments.
- Implement comprehensive unit tests to verify the program's correctness.

## Getting Started

1. Clone the repository to your local machine.
2. Ensure you have Go installed.
3. Navigate to the project directory.
4. Run the program using the command `go run . <input_file> <output_file>`.
5. Review the output file to verify the modifications.
