# Zendesk Coding Challenge

* [Installation Guide](#installation)
* [User Guide](#run-the-program)
* [Design decisions](#design-decisions)
* [Assumptions](#assumption)

---
## Installation
1. Install [go](https://golang.org/doc/install)

2. Clone the project into your desired directory:

```
git clone https://github.com/willshen8/zendesk-coding-challenge.git zendesk
```

3. cd into the cloned directory:

```
cd zendesk
```

4. Build the executable cli program

A Makefile is already created for you and you can simple run
```
make build
```

and then an executable called `zendesk` is created at the root level of your cloned directory.

---
## Run the program
1. Configuration files - The 3 JSON files are stored in the `config` directory by default.

2. To start search run the following command:
```
$ ./zendesk query organisation _id 101
```

where the command follow the convention of 
```
query <table> <field> <value>
```

3. To get help, type 
```
$ ./zendesk --help
```

and you'll receive the following help window:

```
usage: Zendesk-Search [<flags>] <command> [<args> ...]

Welcome to Zendesk Search!

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  config <organisation> <user> <ticket>
    Config the data source files by specifying the files you want to use.

  query <table> <field> [<value>]
    Search a specific field in a table.
```

<strong>Note: </strong>Optional fields are denoted by `|`. For example `value` is an optional argument, and when left blank all values will be populated back as results window.

4. To run all test files
```
$ make test
```

5. To check for code coverage
```
$ make coverage
```
6. To all everything in one go, e.g. `build` ->  `lint` -> `test` -> `coverage`
```
$ make all
```
---

# Design Decisions

1. Data structure:

Database is best for searching data as we all know, and to resemble database characteristics without
building one. I have considered few options and the best one appears to be a `map` of `map` of `map`.

`map` is native data structure in Go, and is a key-value pair structure.
* The first level map contains the name of the table `organisation`, `user` and `ticket`, and each key points to another map that contains the field/row names.
* The second level `map`'s key values are the `ID` field of the data.
* The third level `map` basically contains the underlying data in each data record.

Go's map are implemented using hash-map, which means it offers the benefits of O(1) for lookups, and O(n) in the worst case.


2. Extensibility 

As we are using map, adding additional tables is simply a matter of insert into the root map, and then define the relationships in `entity.go` file.

3. Database entity relationships
* `Organisation` table has a foreign key of `organization_id` which is part of `user` and `ticket` table.
* `User`table contains two foreign keys: `submitter_id` and `assignee_id`.
* `Ticket` table contains no foreign keys. 

---
# Assumptions

1. `external_id`field appears to be linked to external system and bears no relationship with other tables.

# Limitations

1. The size of the database is depend on the hardware or the max size of the heap, as it is implemented using `map`.