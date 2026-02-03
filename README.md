# expense-tracker
Simple CLI expense-tracker to manage finances.
It helps you manage your finances by recording expenses, updating or deleting them, and viewing summaries.

Project Purpose is to practice Go fundamentals, cobra-cli, file handling, JSON serialization.

The project idea is from the: https://roadmap.sh/projects/expense-tracker
### Features
The application supports the following commands:
* Add an expense with a description and amount
* Update an existing expense by ID
* Delete an expense by ID
* View all expenses
* View a summary of all expenses

### Requirements
* Go 1.20+
* Command-line environment (Windows, macOS, or Linux)
## Installation
Clone the repository:
```
git clone https://github.com/DinaraGil/expense-tracker.git
cd expense-tracker
```
Build the application:
```
go build -o expense-tracker
```
### Example
```
$ expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)

$ expense-tracker add --description "Dinner" --amount 10
# Expense added successfully (ID: 2)

$ expense-tracker list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10

$ expense-tracker summary
# Total expenses: $30

$ expense-tracker delete --id 2
# Expense deleted successfully

$ expense-tracker summary
# Total expenses: $20
```

### Data Storage

Expenses are stored locally in a JSON file.

Each expense includes:
* ID
* Date
* Description
* Amount
