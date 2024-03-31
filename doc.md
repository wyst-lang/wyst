# Lexer
The Lexer struct is responsible for tokenizing input strings. It reads characters from the input string and converts them into tokens, which represent the smallest units of meaning in the programming language.

# Constructor: new
The new method is a constructor for the Lexer struct. It takes an input string as its argument and returns a new Lexer instance initialized with the input string and an initial position of 0.

# Method: advance
The advance method moves the lexer's position to the next character in the input string. It increments the position field by 1, allowing the lexer to progress through the input string.

# Method: current_char
The current_char method retrieves the current character from the input string without advancing the lexer's position. It returns the character at the current position as an Option<char>. If the end of the input string is reached, it returns None.

# Method: next_token
The next_token method tokenizes the input string and returns the next token. It uses the advance and current_char methods to iterate through the input string, skipping whitespace characters and returning tokens based on the current character. It matches characters to known token types such as numbers, operators, and parentheses, and constructs tokens accordingly. If an invalid character is encountered, it panics with an error message.