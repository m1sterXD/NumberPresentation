# Number Presentation

This is a dynamic programming algorithm to find the minimal symbolic representation of a number using addition and multiplication. It optimizes the representation to minimize the number of symbols.

## Features
- Calculates the shortest symbolic representation of a given number.
- Supports custom limits for starting numbers.
- Efficiently handles large numbers with optimizations.

## Usage
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/number-presentation.git
   cd number-presentation

### Parameters
goal: The target number to represent.
limit: The maximum starting number for representation.
show: If true, prints the internal dynamic programming table.


Конечно! Документация проекта включает несколько ключевых элементов. Вот основные рекомендации для минимально необходимой документации:

1. README.md
README — это лицо вашего проекта на GitHub. Он должен содержать всю важную информацию о проекте.

Пример содержимого README.md:
markdown
Copy code
# Number Presentation

This is a dynamic programming algorithm to find the minimal symbolic representation of a number using addition and multiplication. It optimizes the representation to minimize the number of symbols.

## Features
- Calculates the shortest symbolic representation of a given number.
- Supports custom limits for starting numbers.
- Efficiently handles large numbers with optimizations.

## Usage
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/number-presentation.git
   cd number-presentation
Run the program:

bash
Copy code
go run main.go
Customize the goal, limit, and show parameters in main.go:

go
Copy code
n, p := NumberPresentation(343, 1, true)
fmt.Println("Representation:", n, "Symbols:", p)
Parameters
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

