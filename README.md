# Zendesk Coding Challenge

* [Installation Guide](#installation)
* [User Guide](#run-the-program)
* [Design decisions](#design-decisions)
* [Assumptions](#assumption)
* [Limitations](#limitations)

---
## Installation
1. Install [go](https://golang.org/doc/install)

2. Clone the project into your desired directory:

```
$ git clone https://github.com/willshen8/zendesk-coding-challenge.git zendesk
```

3. cd into the cloned directory:

```
$ cd zendesk
```

4. Build the executable

A Makefile is already created for you and you can simple run
```
make build
```

and then an executable called `zendesk` is created at the root level of your cloned directory.

---
## Run the program

1. To start search run the following command:
```
$ ./zendesk query organisation _id 101
```

where the command follow the convention of 
```
$ ./zendesk query <table> <field> <value>

```
2. Configuration files:
The 3 JSON files are stored in the `config` directory by default. To specify the data source you wish to use, use the following command:

```
$ ./zendesk config <organisations.json> <users.json> <tickets.json>
```
The order of the files matters.

3. To list all the available fields in a table:

```
$ ./zendesk list <table>
```

4. To get help, type 
```
$ ./zendesk --help
```

and you'll receive the following help window:

```
Welcome to Zendesk Search!

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  config <organisation json file> <users json file> <tickets json file>
    Config the data source files by specifying the files you want to use.

  list <table>
    List all available data fields in a table.

  query <table> <field> [<value>]
    Search a specific field in a table.
```

<strong>Note: </strong>Optional fields are denoted by `|`. For example `value` is an optional argument, and when left blank all values will be populated back as results window.

5. To run all test files
```
$ make test
```

6. To check for code coverage
```
$ make coverage
```
7. To all everything in one go, e.g. `build` ->  `lint` -> `test` -> `coverage`
```
$ make all
```
## Design Decisions
---

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

4. Use of CLI Building Library - [Kingping 2](https://github.com/alecthomas/kingpin)
This library is simple to use and does lot of error handling when user is specifying command line arguments. It is chosen over other popular library such as `Cobra` and `urfave/cli` for its better architecture overall.

## Assumptions
---
1. `external_id` field appears to be linked to external system and bears no relationship with other tables.

## Limitations
---
1. The size of the database is depend on the hardware or the max size of the heap, as it is implemented using `map`.
2. Use the `config` command option to specify the source data files, the order matters. Due to time constraint I have not implemented checks to ensure the right files are specified.
3. As `map` is used, the order of the fields when printing data is not guaranteed, this can be easily achieved by using a look up table/slice. But due to time constraint, it is treated as a `non-important feature` and not implemented. 
