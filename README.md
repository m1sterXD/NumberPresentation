# Number Presentation

This is a dynamic programming algorithm to find the minimal symbolic representation of a number using addition and multiplication. It optimizes the representation to minimize the number of symbols.

## Features
- Calculates the shortest symbolic representation of a given number.
- Supports custom limits for starting numbers.
- Efficiently handles large numbers with optimizations.

## Usage
1. Clone the repository:
   git clone https://github.com/m1sterXD/NumberPresentation
   cd number-presentation

### Parameters
goal: The target number to represent.
limit: The maximum starting number for representation.
show: If true, prints the internal dynamic programming table.

### How It Works
The algorithm uses dynamic programming to:
Minimize the number of symbols in the representation.
Combine addition and multiplication to find the optimal path.

### Adding Section
The algorithm checks combinations of smaller numbers and adds their lengths.

### Multiplication Section
Efficiently handles multiplication, considering only integer divisors and optimizing for fewer symbols.

### Contributing
Feel free to fork this repository and submit pull requests for improvements.

### License
This project is licensed under the MIT License. See LICENSE for details.


### number-presentation/
├── main.go       
├── README.md       
├── LICENSE         

