# Contributing to Kick&Roar Backend

## Development Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and configure
3. Install dependencies: `make deps`
4. Run the server: `make run`

## Coding Standards

### Go Style
- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Write meaningful commit messages

### Commit Message Format
```
<type>(<scope>): <subject>

<body>
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `chore`: Maintenance tasks
- `refactor`: Code refactoring
- `test`: Adding tests
- `perf`: Performance improvements

**Examples:**
```
feat(auth): add JWT refresh token functionality
fix(user): resolve email validation bug
docs(readme): update installation instructions
```

## Branch Strategy

- `main`: Production-ready code
- `develop`: Development branch (future)
- `feature/*`: New features
- `fix/*`: Bug fixes

## Pull Request Process

1. Create a feature branch
2. Make your changes
3. Write/update tests
4. Update documentation
5. Submit PR with clear description

## Code Review

- All code must be reviewed before merging
- Address all comments
- Ensure tests pass
- Keep PRs focused and small

## Questions?

Contact: [Your Email]
