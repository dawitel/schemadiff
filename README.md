# Schema Diff Tool

A powerful tool for comparing SQL schemas and generating migration files. This tool helps database administrators and developers manage schema changes between different environments.

## Features

- Compare SQL schemas between development and production environments
- Generate safe migration files
- Support for tables, sequences, functions, and constraints
- Clean architecture design for maintainability
- Comprehensive test coverage

## Installation

```bash
go install github.com/dawitel/schemadiff@latest
```

## Usage

```bash
schemadiff [options]
```

### Options

- `-dev string`: Path to development schema directory (default: "./schemas/dev")
- `-prod string`: Path to production schema directory (default: "./schemas/prod")
- `-output string`: Output directory for migration files (default: "./migrations")
- `-help`: Show help message

### Example

```bash
schemadiff -dev ./schemas/dev -prod ./schemas/prod -output ./migrations
```

## Schema Directory Structure

Place your SQL schema files in the respective directories:

```
schemas/
├── dev/
│   ├── tables.sql
│   ├── functions.sql
│   └── sequences.sql
└── prod/
    ├── tables.sql
    ├── functions.sql
    └── sequences.sql
```

## Generated Migration

The tool generates a single migration file containing all necessary changes:

```sql
-- Generated migration

BEGIN;

-- Table changes
CREATE TABLE new_table (...);
ALTER TABLE existing_table ...;

-- Sequence changes
CREATE SEQUENCE new_sequence;

-- Function changes
CREATE OR REPLACE FUNCTION ...;

COMMIT;
```

## Features

### Supported Schema Objects

- Tables
  - Columns (add, modify, drop)
  - Constraints
  - Indexes
- Sequences
- Functions
- Triggers

### Safety Features

- Transaction wrapping
- IF EXISTS/IF NOT EXISTS checks
- Dependency order handling
- Safe column modifications

## Development

### Prerequisites

- Go 1.21 or higher
- Make (optional, for using Makefile)

### Building

```bash
go build -o schemadiff cmd/schemadiff/main.go
```

### Testing

```bash
go test ./...
```

### Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [sqlparser](https://github.com/xwb1989/sqlparser) for SQL parsing capabilities
