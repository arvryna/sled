# Sled - Time management for Geeks
- A CLI alternative to Harvest(getharvest.com)
- Provides a way to create and manage timesheet to share with your employers
- Interactive CLI for time tracking
- Code is written with Clean Architecture
- Ability to analyse how you spend your time, to make a summary of how your day was, and
  to reflect how the time was spent to a finer detail

## Features:
- Can be configured to generate weekly/monthly timesheet and auto-email as report
- All data is stored as SQLite file in `build/sled_db` folder
- Ability to track time with a timer
- Export data as CSV/JSON
- Fine control over different config

## Todo:
- Make the CLI more interactive
- Ability to generate analytics using differnt charts 
- Add a watch command to see a short summary of week progress
- Introduce plugin based system to push task to Asana/Jira/elsewhere
- Introduce a central server for data dump for teams
- To generate invoice based on the time spent
- Ability to add task without Timer
- Use Cobra lib to add multiple commands instead of using prompt