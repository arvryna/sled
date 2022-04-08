# Sled - Time tracker for Geeks
- A CLI alternative to [Harvest](https://www.getharvest.com/), provides interactive time tracking
- Provides a way to create and manage timesheet to share with your employers or to track for personal todos
- Ability to analyse how you spend your time, to make a summary of how your day was, and to reflect how the time was spent to a finer detail
- Code is written with [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) principles

## Features:
- All data is stored as SQLite file in `~/.sled/sled_db`, can be quickly explored using [SQlite](https://sqlitebrowser.org/) browser
- Ability to track time with a timer
- Export data as CSV/JSON
- Fine control over different config

## Usage:
![sled-prompt](https://github.com/arvryna/sled/blob/main/res/sled-prompt.png)

## FAQ:

### How can i make DB migrations if i make changes to Entity ?
- If you wanted to make changes to DB, sled comes with gorm and automigration enabled, simply change the required fields in Entity/add new ones, restart the application, db will be auto-migrated. NOTE: Gorm can add new columns, entities, change existing types, but as of version `Gorm v1.23.2`, existing columns can't be deleted. So if a perfect migration is important to you, then you need to write manual migrations. In future i would add scripts and flows for it.
