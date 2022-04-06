# Sled - Time management for Geeks
- A CLI alternative to [Harvest](https://www.getharvest.com/), provides interactive time tracking
- Provides a way to create and manage timesheet to share with your employers or to track for personal tracking
- Ability to analyse how you spend your time, to make a summary of how your day was, and to reflect how the time was spent to a finer detail
- Code is written with [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) principles

## Features:
- All data is stored as SQLite file in `build/sled_db` folder, can be quickly explored using [SQlite](https://sqlitebrowser.org/) browser
- Can be configured to generate weekly/monthly timesheet and auto-email as report
- Ability to track time with a timer
- Export data as CSV/JSON
- Fine control over different config

## Usage:
![sled-prompt](https://github.com/arvryna/sled/blob/main/res/sled-prompt.png)