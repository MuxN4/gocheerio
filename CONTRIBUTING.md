# Contributing to GoCheerio
Thank you for considering contributing to GoCheerio! While I'm currently working on this solo, any help to improve the project is welcome and appreciated.

## Development Roadmap
The project follows a comprehensive roadmap with six key phases:

1. Project Setup and Core DOM Implementation ✓
2. Core Cheerio Features
3. Advanced Features Implementation
4. Performance Optimization
5. Documentation 

## How to Contribute
Currently looking for help in several areas. Below are the priority areas where contributions are most needed:


### Priority Areas
1. CSS Selector Enhancements
   - Implement support for complex selectors.
   - Add pseudo-class selector parsing.
   - Improve attribute selector matching.
2. Performance Optimization
   - Create benchmarking tests.
   - Optimize selector matching performance.
   - Implement caching mechanisms for frequently used nodes.
3. Streaming Parser Development
   - Design a memory-efficient streaming parser.
   - Develop callback-based processing.
   - Enable handling of large HTML documents.
4. Test Coverage
   - Increase unit test coverage across the codebase.


## Contribution Process
1. Fork the Repository
2. Create a Feature Branch
   - Use a descriptive name for your branch: feature/short-description.
3. Implementation Guidelines
   - Follow Go formatting standards (`gofmt`).
   - Write clear, concise comments and documentation.
   - Ensure consistent code style and formatting.
   - Add tests to cover any new functionality.
4. Submit a Pull Request
   - Ensure your pull request description explains the problem you’re solving and the solution.
   - Link any relevant issues (e.g., `Closes #12`).
   - Make sure all tests pass and your changes are properly linted.

## Development Setup
To get started with contributing, follow these setup instructions:

```bash
# Clone the repository
git clone https://github.com/MuxN4/gocheerio.git
cd gocheerio

# Install dependencies
go mod download

# Run tests to ensure everything is working
go test ./...

# Run linters to check for code quality
golangci-lint run
```

## Code of Conduct
By participating in this project, you agree to abide by the following principles:

- Respect and Inclusion: Treat all contributors with respect and kindness.
- Constructive Feedback: Provide feedback that is helpful and supportive.
- Transparent Collaboration: Work openly and collaboratively with others.

## Licensing
By contributing to GoCheerio, you agree that your contributions will be licensed under the project's MIT License.